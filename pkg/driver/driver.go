package driver

import (
	"encoding/json"
	"os"
)

func Run() error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}

	config, err := ReadConfig(curDir)
	if err != nil {
		return err
	}

	baseDir, err := makeBaseDir(curDir)
	if err != nil {
		return err
	}

	sfMeta, err := ReadSalseforceMeta(baseDir)
	if err != nil {
		return err
	}

	schema, err := ConvertSchema(sfMeta)
	if err != nil {
		return err
	}

	err = PostProcess(config, baseDir, schema)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(bytes)
	if err != nil {
		return err
	}

	// err = os.Stdout.Sync()
	// if err != nil {
	// 	panic(err) // TODO:
	// }

	return nil
}
