package driver

import "encoding/xml"

type SfCustomObject struct {
	Fields          map[string]*SfCustomField
	ValidationRules map[string]*SfValidationRule
	XMLName         xml.Name `xml:"CustomObject"`

	FullName string
	Label    string `xml:"label"`

	AllowInChatterGroups    bool   `xml:"allowInChatterGroups"`
	CompactLayoutAssignment string `xml:"compactLayoutAssignment"`
	CustomSettingsType      string `xml:"customSettingsType"`
	DeploymentStatus        string `xml:"deploymentStatus"`
	EnableActivities        bool   `xml:"enableActivities"`
	EnableBulkApi           bool   `xml:"enableBulkApi"`
	EnableFeeds             bool   `xml:"enableFeeds"`
	EnableHistory           bool   `xml:"enableHistory"`
	EnableLicensing         bool   `xml:"enableLicensing"`
	EnableReports           bool   `xml:"enableReports"`
	EnableSearch            bool   `xml:"enableSearch"`
	EnableSharing           bool   `xml:"enableSharing"`
	EnableStreamingApi      bool   `xml:"enableStreamingApi"`
	ExternalSharingModel    string `xml:"externalSharingModel"`

	NameField struct {
		DisplayFormat string `xml:"displayFormat"`
		Label         string `xml:"label"`
		TrackHistory  bool   `xml:"trackHistory"`
		Type          string `xml:"type"`
	} `xml:"nameField"`

	ActionOverrides struct {
		ActionName string `xml:"actionName"`
		FormFactor string `xml:"formFactor"`
		Type       string `xml:"type"`
	} `xml:"actionOverrides"`
}

type SfCustomField struct {
	XMLName xml.Name `xml:"CustomField"`

	Type     string `xml:"type"`
	FullName string `xml:"fullName"`
	Label    string `xml:"label"`

	Required      bool `xml:"required"`
	Unique        bool `xml:"unique"`
	CaseSensitive bool `xml:"caseSensitive"`
	ExternalId    bool `xml:"externalId"`

	Length    int `xml:"length"`
	Precision int `xml:"precision"`
	Scale     int `xml:"scale"`

	DefaultValue         string `xml:"defaultValue"`
	Formula              string `xml:"formula"`
	FormulaTreatBlanksAs string `xml:"formulaTreatBlanksAs"`

	ReferenceTo       string `xml:"referenceTo"`
	RelationshipName  string `xml:"relationshipName"`
	RelationshipLabel string `xml:"relationshipLabel"`

	DeleteConstraint         string `xml:"deleteConstraint"`
	Description              string `xml:"description"`
	DisplayFormat            string `xml:"displayFormat"`
	DisplayLocationInDecimal bool   `xml:"displayLocationInDecimal"`
	FieldManageability       string `xml:"fieldManageability"`
	InlineHelpText           string `xml:"inlineHelpText"`
	MaskChar                 string `xml:"maskChar"`
	MaskType                 string `xml:"maskType"`
	RelationshipOrder        int    `xml:"relationshipOrder"`
	ReparentableMasterDetail bool   `xml:"reparentableMasterDetail"`
	TrackHistory             bool   `xml:"trackHistory"`
	TrackTrending            bool   `xml:"trackTrending"`
	VisibleLines             int    `xml:"visibleLines"`
	WriteRequiresMasterRead  bool   `xml:"writeRequiresMasterRead"`

	LookupFilter struct {
		Active      bool   `xml:"active"`
		InfoMessage string `xml:"infoMessage"`
		IsOptional  bool   `xml:"isOptional"`

		FilterItems []struct {
			Field     string `xml:"field"`
			Operation string `xml:"operation"`
			Value     string `xml:"value"`
		} `xml:"filterItems"`
	} `xml:"lookupFilter"`

	ValueSet struct {
		ControllingField   string               `xml:"controllingField"`
		Restricted         bool                 `xml:"restricted"`
		ValueSetDefinition SfValueSetDefinition `xml:"valueSetDefinition"`
		ValueSetName       string               `xml:"valueSetName"`

		ValueSettings []struct {
			ControllingFieldValue []string `xml:"controllingFieldValue"`
			ValueName             string   `xml:"valueName"`
		} `xml:"valueSettings"`
	} `xml:"valueSet"`
}

type SfCustomValue struct {
	FullName string `xml:"fullName"`
	Default  bool   `xml:"default"`
	Label    string `xml:"label"`
}

type SfValueSetDefinition struct {
	Sorted bool            `xml:"sorted"`
	Value  []SfCustomValue `xml:"value"`
}

type SfGlobalValueSet struct {
	XMLName     xml.Name `xml:"GlobalValueSet"`
	Name        string
	CustomValue []SfCustomValue `xml:"customValue"`
	Sorted      bool            `xml:"sorted"`
}

type SfValidationRule struct {
	XMLName               xml.Name `xml:"ValidationRule"`
	FullName              string   `xml:"fullName"`
	Active                bool     `xml:"active"`
	Description           string   `xml:"description"`
	ErrorConditionFormula string   `xml:"errorConditionFormula"`
	ErrorDisplayField     string   `xml:"errorDisplayField"`
	ErrorMessage          string   `xml:"errorMessage"`
}

type SfRestrictionRule struct {
	XMLName         xml.Name `xml:"RestrictionRule"`
	Active          bool     `xml:"active"`
	Description     string   `xml:"description"`
	EnforcementType string   `xml:"enforcementType"`
	MasterLabel     string   `xml:"masterLabel"`
	RecordFilter    string   `xml:"recordFilter"`
	TargetEntity    string   `xml:"targetEntity"`
	UserCriteria    string   `xml:"userCriteria"`
	Version         int      `xml:"version"`
}

type SfFlow struct {
	XMLName xml.Name `xml:"Flow"`
	Name    string
	Label   string `xml:"label"`
	Status  string `xml:"status"`
	Start   struct {
		Object            string `xml:"object"`
		RecordTriggerType string `xml:"recordTriggerType"`
		TriggerType       string `xml:"triggerType"`
	} `xml:"start"`
}

type SfApexTrigger struct {
	Name         string
	TargetEntity string
	Events       string
}

type SalesforceMeta struct {
	GlobalValueSets  map[string]*SfGlobalValueSet
	RestrictionRules map[string]*SfRestrictionRule
	Flows            map[string]*SfFlow
	ApexTriggers     map[string]*SfApexTrigger
	SObjects         map[string]*SfCustomObject
}
