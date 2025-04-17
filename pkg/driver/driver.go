package driver

import (
	"encoding/json"
	"os"
)

func ReadSalseforceMeta(baseDir string) (SalesforceMeta, error) {
	var retval SalesforceMeta
	var err error

	retval.GlobalValueSets, err = readGlobalValueSetsMeta(baseDir)
	if err != nil {
		return retval, err
	}

	retval.RestrictionRules, err = readRestrictionRulesMeta(baseDir)
	if err != nil {
		return retval, err
	}

	retval.Flows, err = readFlowsMeta(baseDir)
	if err != nil {
		return retval, err
	}

	retval.ApexTriggers, err = readApexTriggers(baseDir)
	if err != nil {
		return retval, err
	}

	retval.SObjects, err = readObjectsMeta(baseDir, retval.GlobalValueSets)
	if err != nil {
		return retval, err
	}

	return retval, nil
}

func Run() error {
	baseDir, err := makeBaseDir()
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
