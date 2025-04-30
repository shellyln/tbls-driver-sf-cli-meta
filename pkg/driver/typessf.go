package driver

import "encoding/xml"

type SfCustomObject struct {
	Fields          map[string]*SfCustomField
	RecordTypes     map[string]*SfRecordType
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
	RecordTypeTrackHistory  bool   `xml:"recordTypeTrackHistory"`
	SharingModel            string `xml:"sharingModel"`
	Visibility              string `xml:"visibility"`

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

type SfRecordType struct {
	XMLName     xml.Name `xml:"RecordType"`
	FullName    string   `xml:"fullName"`
	Description string   `xml:"description"`
	Label       string   `xml:"label"`
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

type SfPermissionSet struct {
	XMLName           xml.Name `xml:"PermissionSet"`
	ObjectPermissions []struct {
		Object           string `xml:"object"`
		AllowCreate      bool   `xml:"allowCreate"`
		AllowRead        bool   `xml:"allowRead"`
		AllowEdit        bool   `xml:"allowEdit"`
		AllowDelete      bool   `xml:"allowDelete"`
		ViewAllRecords   bool   `xml:"viewAllRecords"`
		ModifyAllRecords bool   `xml:"modifyAllRecords"`
	} `xml:"objectPermissions"`

	FieldPermissions []struct {
		Field    string `xml:"field"`
		Readable bool   `xml:"readable"`
		Editable bool   `xml:"editable"`
	} `xml:"fieldPermissions"`
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

type SfSharingRules struct {
	XMLName               xml.Name                  `xml:"SharingRules"`
	SharingCriteriaRules  []SfSharingCriteriaRule   `xml:"sharingCriteriaRules"`
	SharingGuestRules     []SfSharingGuestRules     `xml:"sharingGuestRules"`
	SharingOwnerRules     []SfSharingOwnerRules     `xml:"sharingOwnerRules"`
	SharingTerritoryRules []SfSharingTerritoryRules `xml:"sharingTerritoryRules"`
}

type SfSharingBaseRule struct {
	FullName        string                         `xml:"fullName"`
	Label           string                         `xml:"label"`
	AccessLevel     string                         `xml:"accessLevel"`
	Description     string                         `xml:"description"`
	SharedTo        SfSharedTo                     `xml:"sharedTo"`
	AccountSettings []SfAccountSharingRuleSettings `xml:"accountSettings"`
}

type SfSharingBaseCriteriaRule struct {
	SfSharingBaseRule
	BooleanFilter string         `xml:"booleanFilter"`
	CriteriaItems []SfFilterItem `xml:"criteriaItems"`
}

type SfSharingCriteriaRule struct {
	SfSharingBaseCriteriaRule
	IncludeRecordsOwnedByAll bool `xml:"includeRecordsOwnedByAll"`
}

type SfSharingGuestRules struct {
	SfSharingBaseCriteriaRule
	IncludeHVUOwnedRecords bool `xml:"includeHVUOwnedRecords"`
}

type SfSharingOwnerRules struct {
	SfSharingBaseRule
	SharedFrom SfSharedTo `xml:"sharedFrom"`
}

type SfSharingTerritoryRules struct {
	SfSharingOwnerRules
}

type SfSharedTo struct {
	AllCustomerPortalUsers      *struct{} `xml:"allCustomerPortalUsers"`
	AllInternalUsers            *struct{} `xml:"allInternalUsers"`
	AllPartnerUsers             *struct{} `xml:"allPartnerUsers"`
	ChannelProgramGroup         *struct{} `xml:"channelProgramGroup"`
	Group                       []string  `xml:"group"`
	GuestUser                   []string  `xml:"guestUser"`
	ManagerSubordinates         []string  `xml:"managerSubordinates"`
	Managers                    []string  `xml:"managers"`
	PortalRole                  []string  `xml:"portalRole"`
	PortalRoleAndSubordinates   []string  `xml:"portalRoleAndSubordinates"`
	Role                        []string  `xml:"role"`
	RoleAndSubordinates         []string  `xml:"roleAndSubordinates"`
	RoleAndSubordinatesInternal []string  `xml:"roleAndSubordinatesInternal"`
	Territory                   []string  `xml:"territory"`
	TerritoryAndSubordinates    []string  `xml:"territoryAndSubordinates"`
	Queue                       []string  `xml:"queue"`
}

type SfAccountSharingRuleSettings struct {
	CaseAccessLevel        string `xml:"caseAccessLevel"`
	ContactAccessLevel     string `xml:"contactAccessLevel"`
	OpportunityAccessLevel string `xml:"opportunityAccessLevel"`
}

type SfFilterItem struct {
	Field      string `xml:"field"`
	Operation  string `xml:"operation"`
	Value      string `xml:"value"`
	ValueField string `xml:"valueField"`
}

type SfDuplicateRule struct {
	XMLName                 xml.Name                   `xml:"DuplicateRule"`
	IsActive                bool                       `xml:"isActive"`
	MasterLabel             string                     `xml:"masterLabel"`
	Description             string                     `xml:"description"`
	ActionOnInsert          string                     `xml:"actionOnInsert"`
	ActionOnUpdate          string                     `xml:"actionOnUpdate"`
	AlertText               string                     `xml:"alertText"`
	SecurityOption          string                     `xml:"securityOption"`
	SortOrder               int                        `xml:"sortOrder"`
	DuplicateRuleFilter     DuplicateRuleFilter        `xml:"duplicateRuleFilter"`
	OperationsOnInsert      []string                   `xml:"operationsOnInsert"`
	OperationsOnUpdate      []string                   `xml:"operationsOnUpdate"`
	DuplicateRuleMatchRules []SfDuplicateRuleMatchRule `xml:"duplicateRuleMatchRules"`
}

type DuplicateRuleFilter struct {
	BooleanFilter            string                    `xml:"booleanFilter"`
	DuplicateRuleFilterItems []DuplicateRuleFilterItem `xml:"duplicateRuleFilterItems"`
}

type DuplicateRuleFilterItem struct {
	SfFilterItem
	SortOrder int    `xml:"sortOrder"`
	Table     string `xml:"table"`
}

type SfDuplicateRuleMatchRule struct {
	MatchRuleSObjectType string          `xml:"matchRuleSObjectType"`
	MatchingRule         string          `xml:"matchingRule"`
	ObjectMapping        SfObjectMapping `xml:"objectMapping"`
}

type SfObjectMapping struct {
	InputObject   string `xml:"inputObject"`
	MappingFields []struct {
		InputField  string `xml:"inputField"`
		OutputField string `xml:"outputField"`
	} `xml:"mappingFields"`
}

type SfMatchingRules struct {
	XMLName       xml.Name         `xml:"MatchingRules"`
	MatchingRules []SfMatchingRule `xml:"matchingRules"`
}

type SfMatchingRule struct {
	FullName          string               `xml:"fullName"`
	Description       string               `xml:"description"`
	Label             string               `xml:"label"`
	MatchingRuleItems []SfMatchingRuleItem `xml:"matchingRuleItems"`
	RuleStatus        string               `xml:"ruleStatus"`
}

type SfMatchingRuleItem struct {
	BlankValueBehavior string `xml:"blankValueBehavior"`
	FieldName          string `xml:"fieldName"`
	MatchingMethod     string `xml:"matchingMethod"`
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

type SfApexTriggerCode struct {
	Name         string
	TargetEntity string
	Events       string
	Status       string
}

type SfApexTriggerMeta struct {
	XMLName    xml.Name `xml:"ApexTrigger"`
	ApiVersion string   `xml:"apiVersion"`
	Status     string   `xml:"status"`
}

type SalesforceMeta struct {
	PermissionSets   map[string]*SfPermissionSet
	GlobalValueSets  map[string]*SfGlobalValueSet
	RestrictionRules map[string]*SfRestrictionRule
	SharingRules     map[string]*SfSharingRules
	DuplicateRules   map[string]*SfDuplicateRule
	MatchingRules    map[string]*SfMatchingRules
	Flows            map[string]*SfFlow
	ApexTriggers     map[string]*SfApexTriggerCode
	SObjects         map[string]*SfCustomObject
}
