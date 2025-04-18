package driver

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func readGlobalValueSetsMeta(baseDir string) (map[string]*SfGlobalValueSet, error) {
	vsMap := make(map[string]*SfGlobalValueSet)

	valueSetDir, err := filepath.Abs(filepath.Join(baseDir, "force-app", "main", "default", "globalValueSets"))
	if err != nil {
		return nil, err
	}
	valueSets, err := os.ReadDir(valueSetDir)
	if err != nil {
		return nil, err
	}

	for _, vs := range valueSets {
		if vs.IsDir() {
			continue
		}
		if !strings.HasSuffix(vs.Name(), ".globalValueSet-meta.xml") {
			continue
		}

		fvs, err := os.Open(filepath.Join(valueSetDir, vs.Name()))
		if err != nil {
			return nil, err
		}
		defer fvs.Close()

		var vsMeta SfGlobalValueSet
		vsDec := xml.NewDecoder(fvs)
		err = vsDec.Decode(&vsMeta)
		if err != nil {
			return nil, err
		}

		vsMeta.Name = strings.TrimSuffix(vs.Name(), ".globalValueSet-meta.xml")
		vsMap[vsMeta.Name] = &vsMeta
	}

	return vsMap, nil
}

func readRestrictionRulesMeta(baseDir string) (map[string]*SfRestrictionRule, error) {
	ruleMap := make(map[string]*SfRestrictionRule)

	rulesDir, err := filepath.Abs(filepath.Join(baseDir, "force-app", "main", "default", "restrictionRules"))
	if err != nil {
		return nil, err
	}
	rules, err := os.ReadDir(rulesDir)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if rule.IsDir() {
			continue
		}
		if !strings.HasSuffix(rule.Name(), ".rule-meta.xml") {
			continue
		}

		frule, err := os.Open(filepath.Join(rulesDir, rule.Name()))
		if err != nil {
			return nil, err
		}
		defer frule.Close()

		var ruleMeta SfRestrictionRule
		ruleDec := xml.NewDecoder(frule)
		err = ruleDec.Decode(&ruleMeta)
		if err != nil {
			return nil, err
		}

		ruleMap[rule.Name()] = &ruleMeta
	}

	return ruleMap, nil
}

func readFlowsMeta(baseDir string) (map[string]*SfFlow, error) {
	flowMap := make(map[string]*SfFlow)

	flowsDir, err := filepath.Abs(filepath.Join(baseDir, "force-app", "main", "default", "flows"))
	if err != nil {
		return nil, err
	}
	flows, err := os.ReadDir(flowsDir)
	if err != nil {
		return nil, err
	}

	for _, flow := range flows {
		if flow.IsDir() {
			continue
		}
		if !strings.HasSuffix(flow.Name(), ".flow-meta.xml") {
			continue
		}

		fflow, err := os.Open(filepath.Join(flowsDir, flow.Name()))
		if err != nil {
			return nil, err
		}
		defer fflow.Close()

		var flowMeta SfFlow
		flowDec := xml.NewDecoder(fflow)
		err = flowDec.Decode(&flowMeta)
		if err != nil {
			return nil, err
		}

		flowMeta.Name = strings.TrimSuffix(flow.Name(), ".flow-meta.xml")
		flowMap[flowMeta.Name] = &flowMeta
	}

	return flowMap, nil
}

func readValidationRulesMeta(entitiesDir string, entityName string) (map[string]*SfValidationRule, error) {
	ruleMap := make(map[string]*SfValidationRule)

	rulesDir := filepath.Join(entitiesDir, entityName, "validationRules")
	ruless, err := os.ReadDir(rulesDir)
	if err != nil {
		return nil, nil
	}

	for _, rule := range ruless {
		if rule.IsDir() {
			continue
		}
		if !strings.HasSuffix(rule.Name(), ".validationRule-meta.xml") {
			continue
		}

		frule, err := os.Open(filepath.Join(rulesDir, rule.Name()))
		if err != nil {
			return nil, err
		}
		defer frule.Close()

		var ruleMeta SfValidationRule
		ruleDec := xml.NewDecoder(frule)
		err = ruleDec.Decode(&ruleMeta)
		if err != nil {
			return nil, err
		}

		ruleMap[rule.Name()] = &ruleMeta
	}

	return ruleMap, nil
}

func readFieldsMeta(
	entitiesDir string, entityName string, sobjMap map[string]*SfCustomObject, vsMap map[string]*SfGlobalValueSet) error {

	fieldsDir := filepath.Join(entitiesDir, entityName, "fields")
	if _, err := os.Stat(fieldsDir); err != nil {
		return nil
	}

	fields, err := os.ReadDir(fieldsDir)
	if err != nil {
		return err
	}

	for _, fld := range fields {
		if fld.IsDir() {
			continue
		}
		if !strings.HasSuffix(fld.Name(), ".field-meta.xml") {
			continue
		}

		ffld, err := os.Open(filepath.Join(fieldsDir, fld.Name()))
		if err != nil {
			return err
		}
		defer ffld.Close()

		var fldMeta SfCustomField
		fldDec := xml.NewDecoder(ffld)
		err = fldDec.Decode(&fldMeta)
		if err != nil {
			return err
		}

		if fldMeta.FullName == "Name" {
			// Standard objects
			fldMeta.Type = "Name"
			fldMeta.Required = true
			fldMeta.ExternalId = true
		}

		if len(fldMeta.ValueSet.ValueSetName) > 0 {
			vsMeta := vsMap[fldMeta.ValueSet.ValueSetName]
			if vsMeta != nil {
				fldMeta.ValueSet.ValueSetDefinition = SfValueSetDefinition{
					Sorted: vsMeta.Sorted,
					Value:  vsMeta.CustomValue,
				}
			}
		}

		if fldMeta.Type == "" {
			fldMeta.Type = "_" // reqired
		}

		objMeta := sobjMap[entityName]
		objMeta.Fields[fld.Name()] = &fldMeta
	}

	return nil
}

func readObjectsMeta(baseDir string, vsMap map[string]*SfGlobalValueSet) (map[string]*SfCustomObject, error) {
	sobjMap := make(map[string]*SfCustomObject)

	entitiesDir, err := filepath.Abs(filepath.Join(baseDir, "force-app", "main", "default", "objects"))
	if err != nil {
		return nil, err
	}
	entities, err := os.ReadDir(entitiesDir)
	if err != nil {
		return nil, err
	}

	for _, ent := range entities {
		if !ent.IsDir() {
			continue
		}

		fobj, err := os.Open(filepath.Join(entitiesDir, ent.Name(), ent.Name()+".object-meta.xml"))
		if err != nil {
			return nil, err
		}
		defer fobj.Close()

		var objMeta SfCustomObject
		objDec := xml.NewDecoder(fobj)
		err = objDec.Decode(&objMeta)
		if err != nil {
			return nil, err
		}

		objMeta.FullName = ent.Name()
		objMeta.Fields = make(map[string]*SfCustomField)
		sobjMap[ent.Name()] = &objMeta

		idFldMeta := SfCustomField{
			Type:       "Id",
			FullName:   "Id",
			Label:      "Id",
			Required:   true,
			ExternalId: true,
		}
		objMeta.Fields[idFldMeta.FullName] = &idFldMeta

		if len(objMeta.NameField.Type) > 0 {
			// Custom objects
			fldMeta := SfCustomField{
				Type:          "Name(" + objMeta.NameField.Type + ")",
				FullName:      "Name",
				Label:         objMeta.NameField.Label,
				Required:      true,
				ExternalId:    true,
				DisplayFormat: objMeta.NameField.DisplayFormat,
				TrackHistory:  objMeta.NameField.TrackHistory,
			}
			objMeta.Fields[fldMeta.FullName] = &fldMeta
		}

		ruleMap, err := readValidationRulesMeta(entitiesDir, ent.Name())
		if err != nil {
			return nil, err
		}
		objMeta.ValidationRules = ruleMap

		err = readFieldsMeta(entitiesDir, ent.Name(), sobjMap, vsMap)
		if err != nil {
			return nil, err
		}
	}

	return sobjMap, nil
}

func readApexTriggers(baseDir string) (map[string]*SfApexTrigger, error) {
	re, err := regexp.Compile(`\btrigger\s+(\S+)\s+on\s+(\S+)\s*\(([^)]*)\)`)
	if err != nil {
		return nil, err
	}

	trigMap := make(map[string]*SfApexTrigger)

	trigDir, err := filepath.Abs(filepath.Join(baseDir, "force-app", "main", "default", "triggers"))
	if err != nil {
		return nil, err
	}
	triggers, err := os.ReadDir(trigDir)
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		if trigger.IsDir() {
			continue
		}
		if !strings.HasSuffix(trigger.Name(), ".trigger") {
			continue
		}

		ftrig, err := os.Open(filepath.Join(trigDir, trigger.Name()))
		if err != nil {
			return nil, err
		}
		defer ftrig.Close()

		bytes, err := io.ReadAll(ftrig)
		if err != nil {
			return nil, err
		}
		lines := string(bytes)

		result := re.FindAllStringSubmatch(lines, 1)
		if result == nil {
			continue
		}

		trigMeta := SfApexTrigger{
			Name:         result[0][1],
			TargetEntity: result[0][2],
			Events:       result[0][3],
		}

		trigMap[trigger.Name()] = &trigMeta
	}

	return trigMap, nil
}
