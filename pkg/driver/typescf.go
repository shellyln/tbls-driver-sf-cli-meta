package driver

import "strings"

type CfDriverConfig struct {
	SuppressFieldDescription bool       `json:"suppressFieldDescription"`
	Escape                   []CfEscape `json:"escape"`
	Include                  *[]string  `json:"include"`
	Exclude                  *[]string  `json:"exclude"`
}

type CfEscape struct {
	Def      string `json:"def"`
	Prop     string `json:"prop"`
	Char     string `json:"char"`
	replacer *strings.Replacer
}
