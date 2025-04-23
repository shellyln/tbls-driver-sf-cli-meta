package driver

import (
	"reflect"
	"strings"
)

func PostProcess(config *PpDriverConfig, baseDir string, schema *Schema) error {
	if len(config.PpReplacements) > 0 {
		for i, r := range config.PpReplacements {
			if len(r.Char) > 0 {
				replPair := make([]string, len(r.Char)*2)
				for i, c := range r.Char {
					replPair[2*i] = string(c)
					replPair[2*i+1] = "\\" + string(c)
				}
				config.PpReplacements[i].replacer = strings.NewReplacer(replPair...)
			}
		}
	}
	replDict := make(map[strPair]PpReplacement)

	WalkAndApply("Schema", schema, func(ancestors []string, parent string, ty reflect.Type, value any) any {
		switch ty.Kind() {
		case reflect.String:
			str := reflect.ValueOf(value).String()
			lastAtor := ""
			if len(ancestors) > 0 {
				lastAtor = ancestors[len(ancestors)-1]
			}
			repl := findReplacer(lastAtor, parent, config.PpReplacements, replDict)

			if repl != nil {
				str = repl.Replace(str)
			}

			return str
		}
		return value
	})

	return nil
}

func findReplacer(def string, prop string, replacements []PpReplacement, dict map[strPair]PpReplacement) *strings.Replacer {
	key := strPair{A: def, B: prop}
	if repl, ok := dict[key]; ok {
		return repl.replacer
	}

	var defMatched *PpReplacement
	var propMatched *PpReplacement
	var globalMatched *PpReplacement

	for _, repl := range replacements {
		if strings.EqualFold(repl.Def, def) && strings.EqualFold(repl.Prop, prop) {
			dict[key] = repl
			return repl.replacer
		}
		if defMatched == nil && strings.EqualFold(repl.Def, def) && repl.Prop == "" {
			defMatched = &repl
		}
		if propMatched == nil && repl.Def == "" && strings.EqualFold(repl.Prop, prop) {
			propMatched = &repl
		}
		if globalMatched == nil && repl.Def == "" && repl.Prop == "" {
			globalMatched = &repl
		}
	}
	if propMatched != nil {
		dict[key] = *propMatched
		return propMatched.replacer
	}
	if defMatched != nil {
		dict[key] = *defMatched
		return defMatched.replacer
	}
	if globalMatched != nil {
		dict[key] = *globalMatched
		return globalMatched.replacer
	}

	dict[key] = PpReplacement{}
	return nil
}
