package driver

import (
	"os"
	"path"

	"github.com/shellyln/go-loose-json-parser/jsonlp"
	"github.com/shellyln/go-loose-json-parser/marshal"
)

func ReadConfig(curDir string) (*CfDriverConfig, error) {
	var config CfDriverConfig

	bytes, err := os.ReadFile(path.Join(curDir, ".tbls-sf-cli-meta.toml"))
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		//json.Unmarshal(bytes, &config)
		parsed, err := jsonlp.ParseTOML(string(bytes), jsonlp.Linebreak_Lf, jsonlp.Interop_None)
		if err != nil {
			return nil, err
		}
		if err := marshal.Unmarshal(parsed, &config, nil); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
