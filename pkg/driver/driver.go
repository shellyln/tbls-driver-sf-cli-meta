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

	baseDir, err := MakeBaseDir(curDir)
	if err != nil {
		return err
	}

	sfMeta, err := ReadSalseforceMeta(config, baseDir)
	if err != nil {
		return err
	}

	schema, err := ConvertSchema(sfMeta)
	if err != nil {
		return err
	}

	err = PostProcess(config, schema)
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
