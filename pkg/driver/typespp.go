package driver

import "strings"

type PpDriverConfig struct {
	PpReplacements []PpReplacement `json:"replacements"`
	Include        *[]string       `json:"include"`
	Exclude        *[]string       `json:"exclude"`
}

type PpReplacement struct {
	Def      string `json:"def"`
	Prop     string `json:"prop"`
	Char     string `json:"char"`
	replacer *strings.Replacer
}

type strPair struct {
	A string
	B string
}
