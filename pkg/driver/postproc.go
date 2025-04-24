package driver

import (
	"reflect"
	"regexp"
	"strings"
)

func PostProcess(config *CfDriverConfig, schema *Schema) error {
	escapeSpecialCharacters(config, schema)
	return nil
}

func escapeSpecialCharacters(config *CfDriverConfig, schema *Schema) {
	if len(config.Escape) > 0 {
		for i, r := range config.Escape {
			if len(r.Char) > 0 {
				replPair := make([]string, len(r.Char)*2)
				for i, c := range r.Char {
					replPair[2*i] = string(c)
					replPair[2*i+1] = "\\" + string(c)
				}
				config.Escape[i].replacer = strings.NewReplacer(replPair...)
			}
		}
	}
	escDict := make(map[strPair]CfEscape)

	WalkAndApply("Schema", schema, func(ancestors []string, parent string, ty reflect.Type, value any) any {
		switch ty.Kind() {
		case reflect.String:
			str := reflect.ValueOf(value).String()
			lastAtor := ""
			if len(ancestors) > 0 {
				lastAtor = ancestors[len(ancestors)-1]
			}
			repl := findReplacer(lastAtor, parent, config.Escape, escDict)

			if repl != nil {
				str = repl.Replace(str)
			}

			return str
		}
		return value
	})
}

func findReplacer(def string, prop string, escapes []CfEscape, dict map[strPair]CfEscape) *strings.Replacer {
	key := strPair{A: def, B: prop}
	if repl, ok := dict[key]; ok {
		return repl.replacer
	}

	var defMatched *CfEscape
	var propMatched *CfEscape
	var globalMatched *CfEscape

	for _, esc := range escapes {
		if strings.EqualFold(esc.Def, def) && strings.EqualFold(esc.Prop, prop) {
			dict[key] = esc
			return esc.replacer
		}
		if defMatched == nil && strings.EqualFold(esc.Def, def) && esc.Prop == "" {
			defMatched = &esc
		}
		if propMatched == nil && esc.Def == "" && strings.EqualFold(esc.Prop, prop) {
			propMatched = &esc
		}
		if globalMatched == nil && esc.Def == "" && esc.Prop == "" {
			globalMatched = &esc
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

	dict[key] = CfEscape{}
	return nil
}

func matchWildcard(pattern, text string) (bool, error) {
	escapedPattern := regexp.QuoteMeta(pattern)
	regexPattern := "^" + strings.ReplaceAll(escapedPattern, `\*`, ".*") + "$"
	return regexp.MatchString(regexPattern, text)
}
