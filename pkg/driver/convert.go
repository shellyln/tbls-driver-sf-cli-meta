package driver

import (
	"sort"
	"strconv"
	"strings"
)

func ConvertSchema(config *CfDriverConfig, sfMeta SalesforceMeta) (*Schema, error) {

	schema := Schema{
		Name:      "",
		Desc:      "",
		Tables:    make([]Table, 0),
		Relations: make([]Relation, 0),
		Functions: make([]Function, 0),
		Enums:     make([]Enum, 0),
		Driver: &Driver{
			Name:            "Salesforce CLI Metadata driver",
			DatabaseVersion: "0",
		},
		Labels:     make([]Label, 0),
		Viewpoints: make([]Viewpoint, 0),
	}

	type permMetaAndObjPerm struct {
		name string
		objp SfObjectPermission
	}
	objPermsMap := make(map[string][]permMetaAndObjPerm)

	type permMetaAndFldPerm struct {
		name     string
		permMeta *SfPermissionSet
		fldp     SfFieldPermission
	}
	fldPermsMap := make(map[string][]permMetaAndFldPerm)

	for permName, permMeta := range sfMeta.PermissionSets {
		for _, objp := range permMeta.ObjectPermissions {
			objPermsMap[objp.Object] = append(objPermsMap[objp.Object], permMetaAndObjPerm{permName, objp})
		}
		for _, fldp := range permMeta.FieldPermissions {
			fldPermsMap[fldp.Field] = append(fldPermsMap[fldp.Field], permMetaAndFldPerm{permName, permMeta, fldp})
		}
	}

	for _, objMeta := range sfMeta.SObjects {
		table := Table{
			Name:             objMeta.FullName,
			Type:             "",
			Comment:          objMeta.Label,
			Columns:          make([]Column, 0),
			Indexes:          make([]Index, 0),
			Constraints:      make([]Constraint, 0),
			Triggers:         make([]Trigger, 0),
			Def:              "",
			Labels:           make([]Label, 0),
			ReferencedTables: make([]string, 0),
		}

		if len(objMeta.CustomSettingsType) > 0 {
			table.Type = "Custom setting"
		} else if strings.HasSuffix(objMeta.FullName, "__c") {
			table.Type = "Custom object"
		} else if strings.HasSuffix(objMeta.FullName, "__mdt") {
			table.Type = "Custom metadata"
		} else {
			table.Type = "Standard object"
		}

		for _, x := range objPermsMap[objMeta.FullName] {
			label := Label{
				Name: x.name + ":",
			}
			if x.objp.AllowCreate {
				label.Name += "C"
			} else {
				label.Name += "-"
			}
			if x.objp.AllowRead {
				label.Name += "R"
			} else {
				label.Name += "-"
			}
			if x.objp.AllowEdit {
				label.Name += "U"
			} else {
				label.Name += "-"
			}
			if x.objp.AllowDelete {
				label.Name += "D"
			} else {
				label.Name += "-"
			}
			label.Name += "/"
			if x.objp.ViewAllRecords {
				label.Name += "V"
			} else {
				label.Name += "-"
			}
			if x.objp.ModifyAllRecords {
				label.Name += "M"
			} else {
				label.Name += "-"
			}
			table.Labels = append(table.Labels, label)
		}

		for _, fldMeta := range objMeta.Fields {
			column := Column{
				Name:     fldMeta.FullName,
				Type:     fldMeta.Type,
				Nullable: !fldMeta.Required,
				Default:  nil,
				ExtraDef: "",
				Labels:   make([]Label, 0),
				Comment:  fldMeta.Label,
			}

			for _, x := range fldPermsMap[objMeta.FullName+"."+fldMeta.FullName] {
				label := Label{
					Name: x.name + ":",
				}
				if x.fldp.Readable {
					label.Name += "R"
				} else {
					label.Name += "-"
				}
				if x.fldp.Editable {
					label.Name += "U"
				} else {
					label.Name += "-"
				}
				column.Labels = append(column.Labels, label)
			}

			if len(fldMeta.ValueSet.ValueSetDefinition.Value) > 0 {
				if len(fldMeta.ValueSet.ValueSetName) > 0 {
					column.Type += "(" + fldMeta.ValueSet.ValueSetName + ")"
				}
				for _, vsMeta := range fldMeta.ValueSet.ValueSetDefinition.Value {
					if len(column.ExtraDef) > 0 {
						column.ExtraDef += "; "
					}
					if vsMeta.Default {
						column.ExtraDef += "[Default] "
					}
					if vsMeta.FullName != vsMeta.Label {
						column.ExtraDef += "{" + vsMeta.Label + ", " + vsMeta.FullName + "}"
					} else {
						column.ExtraDef += vsMeta.FullName
					}
				}
			}

			if !config.SuppressFieldDescription && len(fldMeta.Description) > 0 {
				column.Comment += "; " + fldMeta.Description
			}

			if fldMeta.Length != 0 {
				column.Type = column.Type + "(" + strconv.Itoa(fldMeta.Length) + ")"
			} else if fldMeta.Precision != 0 {
				column.Type = column.Type + "(" + strconv.Itoa(fldMeta.Precision) + ", " + strconv.Itoa(fldMeta.Scale) + ")"
			}

			if len(fldMeta.ReferenceTo) > 0 {
				column.Type = fldMeta.Type + "(" + fldMeta.ReferenceTo + ")"
				column.ExtraDef = "Relation=" + fldMeta.RelationshipName + "; List=" + fldMeta.RelationshipLabel
				if len(fldMeta.LookupFilter.FilterItems) > 0 {
					column.ExtraDef += "; Filter="
					for i, filterItem := range fldMeta.LookupFilter.FilterItems {
						if i != 0 {
							column.ExtraDef += " And "
						}
						column.ExtraDef += filterItem.Field + " " + filterItem.Operation + " " + filterItem.Value
					}
				}
			}
			if len(fldMeta.MaskType) > 0 {
				column.ExtraDef = fldMeta.MaskType + ", " + fldMeta.MaskChar
			}
			if fldMeta.DisplayLocationInDecimal {
				column.ExtraDef = "DisplayLocationInDecimal"
			}
			if len(fldMeta.DisplayFormat) > 0 {
				column.ExtraDef = fldMeta.DisplayFormat
			}
			if len(fldMeta.Formula) > 0 {
				column.Type = "Formula(" + fldMeta.Type + ", " + fldMeta.FormulaTreatBlanksAs + ")"
				column.ExtraDef = fldMeta.Formula
			}

			if len(fldMeta.DefaultValue) > 0 {
				column.Default = fldMeta.DefaultValue
			}

			if fldMeta.FullName == "RecordTypeId" {
				recTypes := make([]*SfRecordType, 0)
				for _, recTypeMeta := range objMeta.RecordTypes {
					recTypes = append(recTypes, recTypeMeta)
				}
				sort.Slice(recTypes, func(i, j int) bool {
					return strings.Compare(recTypes[i].FullName, recTypes[j].FullName) < 0
				})
				for _, recTypeMeta := range recTypes {
					if len(column.ExtraDef) > 0 {
						column.ExtraDef += "; "
					}
					if recTypeMeta.Label != recTypeMeta.FullName {
						column.ExtraDef += "{" + recTypeMeta.Label + ", " + recTypeMeta.FullName + "}"
					} else {
						column.ExtraDef += recTypeMeta.FullName
					}
				}
			}

			table.Columns = append(table.Columns, column)

			if fldMeta.ExternalId {
				index := Index{
					Name:    fldMeta.FullName,
					Def:     "",
					Table:   objMeta.FullName,
					Columns: []string{fldMeta.FullName},
					Comment: "",
				}
				if fldMeta.FullName == "Id" {
					index.Def = "Primary Key"
				} else if fldMeta.FullName == "Name" {
					index.Def = fldMeta.Type
				} else if fldMeta.Unique {
					index.Def = "Unique External Id"
				} else {
					index.Def = "Nonunique External Id"
				}
				table.Indexes = append(table.Indexes, index)
			}

			if fldMeta.FullName == "Id" {
				constraint := Constraint{
					Name:              fldMeta.FullName,
					Type:              "Primary Key",
					Def:               "Primary Key",
					Table:             objMeta.FullName,
					ReferencedTable:   "",
					Columns:           []string{fldMeta.FullName},
					ReferencedColumns: nil,
					Comment:           "",
				}
				table.Constraints = append(table.Constraints, constraint)
			}

			if fldMeta.Unique {
				constraint := Constraint{
					Name:              fldMeta.FullName,
					Type:              "Unique",
					Def:               "",
					Table:             objMeta.FullName,
					ReferencedTable:   "",
					Columns:           []string{fldMeta.FullName},
					ReferencedColumns: nil,
					Comment:           "",
				}
				if fldMeta.CaseSensitive {
					constraint.Def = "Unique Case Sensitive"
				} else {
					constraint.Def = "Unique Case Insensitive"
				}
				table.Constraints = append(table.Constraints, constraint)
			}

			if fldMeta.Type == "MasterDetail" || fldMeta.Type == "Lookup" {
				if len(fldMeta.ReferenceTo) > 0 {
					relation := Relation{
						Table:             objMeta.FullName,
						Columns:           []string{fldMeta.FullName},
						Cardinality:       "zero or more",
						ParentTable:       fldMeta.ReferenceTo,
						ParentColumns:     []string{"Id"},
						ParentCardinality: "exactly one",
						Def:               fldMeta.Type + "\n(" + objMeta.FullName + "." + fldMeta.FullName + ")\n(" + fldMeta.ReferenceTo + "." + fldMeta.RelationshipName + ")",
					}
					schema.Relations = append(schema.Relations, relation)
				}
			}
		}

		for _, flowMeta := range sfMeta.Flows {
			if flowMeta.Start.Object == objMeta.FullName && len(flowMeta.Start.RecordTriggerType) > 0 {
				trigger := Trigger{
					Name:    "flow." + flowMeta.Name,
					Def:     "",
					Comment: flowMeta.Label,
				}
				if flowMeta.Status != "Active" {
					trigger.Def = "[Inactive] "
				}
				trigger.Def += flowMeta.Start.RecordTriggerType + ", " + flowMeta.Start.TriggerType
				table.Triggers = append(table.Triggers, trigger)
			}
		}

		for _, trigMeta := range sfMeta.ApexTriggers {
			if trigMeta.TargetEntity == objMeta.FullName {
				trigger := Trigger{
					Name:    "trigger." + trigMeta.Name,
					Def:     "",
					Comment: "",
				}
				if trigMeta.Status != "Active" {
					trigger.Def = "[Inactive] "
				}
				trigger.Def += trigMeta.Events
				table.Triggers = append(table.Triggers, trigger)
			}
		}

		for _, ruleMeta := range objMeta.ValidationRules {
			constraint := Constraint{
				Name:              ruleMeta.FullName,
				Type:              "ValidationRule",
				Def:               "",
				Table:             objMeta.FullName,
				ReferencedTable:   "",
				Columns:           nil,
				ReferencedColumns: nil,
				Comment:           ruleMeta.Description,
			}
			if !ruleMeta.Active {
				constraint.Def = "[Inactive] "
			}
			if len(ruleMeta.ErrorDisplayField) > 0 {
				constraint.Def += "[" + ruleMeta.ErrorDisplayField + "] "
			}
			constraint.Def += ruleMeta.ErrorConditionFormula
			table.Constraints = append(table.Constraints, constraint)
		}

		for _, ruleMeta := range sfMeta.RestrictionRules {
			if ruleMeta.TargetEntity == objMeta.FullName {
				constraint := Constraint{
					Name:              ruleMeta.MasterLabel,
					Type:              ruleMeta.EnforcementType,
					Def:               "",
					Table:             objMeta.FullName,
					ReferencedTable:   "",
					Columns:           nil,
					ReferencedColumns: nil,
					Comment:           ruleMeta.Description,
				}
				if !ruleMeta.Active {
					constraint.Def = "[Inactive] "
				}
				constraint.Def += ruleMeta.UserCriteria + "; " + ruleMeta.RecordFilter
				table.Constraints = append(table.Constraints, constraint)
			}
		}

		for objName, rules := range sfMeta.SharingRules {
			if objName == objMeta.FullName {
				for _, ruleMeta := range rules.SharingCriteriaRules {
					constraint := Constraint{
						Name:              ruleMeta.FullName,
						Type:              "SharingCriteriaRule",
						Def:               ruleMeta.ToDescription(),
						Table:             objMeta.FullName,
						ReferencedTable:   "",
						Columns:           nil,
						ReferencedColumns: nil,
						Comment:           ruleMeta.Label,
					}
					if len(ruleMeta.Description) > 0 {
						constraint.Comment += "; " + ruleMeta.Description
					}
					table.Constraints = append(table.Constraints, constraint)
				}
				for _, ruleMeta := range rules.SharingGuestRules {
					constraint := Constraint{
						Name:              ruleMeta.FullName,
						Type:              "SharingGuestRule",
						Def:               ruleMeta.ToDescription(),
						Table:             objMeta.FullName,
						ReferencedTable:   "",
						Columns:           nil,
						ReferencedColumns: nil,
						Comment:           ruleMeta.Label,
					}
					if len(ruleMeta.Description) > 0 {
						constraint.Comment += "; " + ruleMeta.Description
					}
					table.Constraints = append(table.Constraints, constraint)
				}
				for _, ruleMeta := range rules.SharingOwnerRules {
					constraint := Constraint{
						Name:              ruleMeta.FullName,
						Type:              "SharingOwnerRule",
						Def:               ruleMeta.ToDescription(),
						Table:             objMeta.FullName,
						ReferencedTable:   "",
						Columns:           nil,
						ReferencedColumns: nil,
						Comment:           ruleMeta.Label,
					}
					if len(ruleMeta.Description) > 0 {
						constraint.Comment += "; " + ruleMeta.Description
					}
					table.Constraints = append(table.Constraints, constraint)
				}
				for _, ruleMeta := range rules.SharingTerritoryRules {
					constraint := Constraint{
						Name:              ruleMeta.FullName,
						Type:              "SharingTerritoryRule",
						Def:               ruleMeta.ToDescription(),
						Table:             objMeta.FullName,
						ReferencedTable:   "",
						Columns:           nil,
						ReferencedColumns: nil,
						Comment:           ruleMeta.Label,
					}
					if len(ruleMeta.Description) > 0 {
						constraint.Comment += "; " + ruleMeta.Description
					}
					table.Constraints = append(table.Constraints, constraint)
				}
			}
		}

		for key, ruleMeta := range sfMeta.DuplicateRules {
			if strings.HasPrefix(key, objMeta.FullName+".") {
				constraint := Constraint{
					Name:              ruleMeta.MasterLabel,
					Type:              "DuplicateRule",
					Def:               "",
					Table:             objMeta.FullName,
					ReferencedTable:   "",
					Columns:           nil,
					ReferencedColumns: nil,
					Comment:           ruleMeta.Description,
				}
				if !ruleMeta.IsActive {
					constraint.Def = "[Inactive] "
				}
				for i, m := range ruleMeta.DuplicateRuleMatchRules {
					if i > 0 {
						constraint.Def += "; "
					}
					if len(m.ObjectMapping.MappingFields) > 0 {
						constraint.Def += " {"
						for j, mf := range m.ObjectMapping.MappingFields {
							if j > 0 {
								constraint.Def += ", "
							}
							constraint.Def += mf.InputField
						}
						constraint.Def += "} "
					}
					constraint.Def += "(" + m.MatchRuleSObjectType + ") " + m.MatchingRule
					if mm, ok := sfMeta.MatchingRules[objMeta.FullName]; ok {
						for _, matchMeta := range mm.MatchingRules {
							if matchMeta.FullName != m.MatchingRule {
								continue
							}
							if matchMeta.RuleStatus != "Active" {
								constraint.Def = "[" + matchMeta.RuleStatus + "] "
							}
							constraint.Def += " {"
							for k, item := range matchMeta.MatchingRuleItems {
								if k > 0 {
									constraint.Def += ", "
								}
								constraint.Def += item.FieldName
							}
							constraint.Def += "}"
							break
						}
					}
				}
				table.Constraints = append(table.Constraints, constraint)
			}
		}

		schema.Tables = append(schema.Tables, table)
	}

	for _, gvs := range sfMeta.GlobalValueSets {
		enum := Enum{
			Name:   gvs.Name,
			Values: make([]string, 0),
		}
		for _, vsMeta := range gvs.CustomValue {
			value := ""
			if vsMeta.Default {
				value += "[Default] "
			}
			if vsMeta.FullName != vsMeta.Label {
				value += "{" + vsMeta.Label + ", " + vsMeta.FullName + "}"
			} else {
				value += vsMeta.FullName
			}
			enum.Values = append(enum.Values, value)
		}
		schema.Enums = append(schema.Enums, enum)
	}

	return &schema, nil
}
