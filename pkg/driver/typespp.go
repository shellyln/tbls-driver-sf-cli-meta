package driver

import "strings"

type PpDriverConfig struct {
	PpReplacements []PpReplacement `json:"replacements"`
}

type PpReplacement struct {
	Def      string `json:"def"`
	Prop     string `json:"prop"`
	Char     string `json:"char"`
	replacer *strings.Replacer
}
