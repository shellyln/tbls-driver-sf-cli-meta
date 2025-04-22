package driver

import (
	"encoding/json"
	"os"
	"path"
	"strings"
)

func escapeSchemaValues(config PpDriverConfig, schema *Schema) {
	if config.EscapedCharacters == "" {
		return
	}

	replPair := make([]string, len(config.EscapedCharacters)*2)
	for i, c := range config.EscapedCharacters {
		replPair[2*i] = string(c)
		replPair[2*i+1] = "\\" + string(c)
	}
	replacer := strings.NewReplacer(replPair...)

	schema.Name = replacer.Replace(schema.Name)
	schema.Desc = replacer.Replace(schema.Desc)

	for i, tbl := range schema.Tables {
		tbl.Name = replacer.Replace(tbl.Name)
		tbl.Type = replacer.Replace(tbl.Type)
		tbl.Def = replacer.Replace(tbl.Def)
		tbl.Comment = replacer.Replace(tbl.Comment)

		for j, col := range tbl.Columns {
			col.Name = replacer.Replace(col.Name)
			col.Type = replacer.Replace(col.Type)
			if v, ok := col.Default.(string); ok {
				col.Default = replacer.Replace(v)
			}
			col.ExtraDef = replacer.Replace(col.ExtraDef)
			col.Comment = replacer.Replace(col.Comment)

			tbl.Columns[j] = col
		}

		for j, con := range tbl.Constraints {
			con.Name = replacer.Replace(con.Name)
			con.Type = replacer.Replace(con.Type)
			con.Table = replacer.Replace(con.Table)
			con.ReferencedTable = replacer.Replace(con.ReferencedTable)
			con.Comment = replacer.Replace(con.Comment)

			//con.ReferencedColumns

			tbl.Constraints[j] = con
		}

		for j, idx := range tbl.Indexes {
			idx.Name = replacer.Replace(idx.Name)
			idx.Table = replacer.Replace(idx.Table)
			idx.Def = replacer.Replace(idx.Def)
			idx.Comment = replacer.Replace(idx.Comment)

			//idx.Columns

			tbl.Indexes[j] = idx
		}

		for j, trg := range tbl.Triggers {
			trg.Name = replacer.Replace(trg.Name)
			trg.Def = replacer.Replace(trg.Def)
			trg.Comment = replacer.Replace(trg.Comment)

			tbl.Triggers[j] = trg
		}

		//tbl.Labels
		//tbl.ReferencedTables

		schema.Tables[i] = tbl
	}

	//schema.Enums
	//schema.Functions
	//schema.Labels
	//schema.Relations
	//schema.Viewpoints
}

func PostProcess(baseDir string, schema *Schema) error {
	var config PpDriverConfig

	bytes, err := os.ReadFile(path.Join(baseDir, ".tbls-sf-cli-meta.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		json.Unmarshal(bytes, &config)
	}

	escapeSchemaValues(config, schema)

	return nil
}
