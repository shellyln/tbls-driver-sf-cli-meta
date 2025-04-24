package driver

import "strings"

type CfDriverConfig struct {
	Escape  []CfEscape `json:"escape"`
	Include *[]string  `json:"include"`
	Exclude *[]string  `json:"exclude"`
}

type CfEscape struct {
	Def      string `json:"def"`
	Prop     string `json:"prop"`
	Char     string `json:"char"`
	replacer *strings.Replacer
}
