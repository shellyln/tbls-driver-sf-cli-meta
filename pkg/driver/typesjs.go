package driver

// tbls JSON Schema
// https://github.com/k1LoW/tbls/blob/main/spec/tbls.schema.json_schema.json

// Column ...
type Column struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Nullable bool        `json:"nullable"`
	Default  interface{} `json:"default,omitempty"`
	ExtraDef string      `json:"extra_def,omitempty"`
	Labels   []Label     `json:"labels,omitempty"`
	Comment  string      `json:"comment,omitempty"`
}

// Constraint ...
type Constraint struct {
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Def               string   `json:"def"`
	Table             string   `json:"table"`
	ReferencedTable   string   `json:"referenced_table,omitempty"`
	Columns           []string `json:"columns,omitempty"`
	ReferencedColumns []string `json:"referenced_columns,omitempty"`
	Comment           string   `json:"comment,omitempty"`
}

// Driver ...
type Driver struct {
	Name            string      `json:"name"`
	DatabaseVersion string      `json:"database_version,omitempty"`
	Meta            *DriverMeta `json:"meta,omitempty"`
}

// DriverMeta ...
type DriverMeta struct {
	CurrentSchema string            `json:"current_schema,omitempty"`
	SearchPaths   []string          `json:"search_paths,omitempty"`
	Dict          map[string]string `json:"dict,omitempty"`
}

// Enum ...
type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// Function ...
type Function struct {
	Name       string `json:"name"`
	ReturnType string `json:"return_type"`
	Arguments  string `json:"arguments"`
	Type       string `json:"type"`
}

// Index ...
type Index struct {
	Name    string   `json:"name"`
	Def     string   `json:"def"`
	Table   string   `json:"table"`
	Columns []string `json:"columns"`
	Comment string   `json:"comment,omitempty"`
}

// Label ...
type Label struct {
	Name    string `json:"name"`
	Virtual bool   `json:"virtual,omitempty"`
}

// Labels ...
//type Labels []Label

// Relation ...
type Relation struct {
	Table             string   `json:"table"`
	Columns           []string `json:"columns"`
	Cardinality       string   `json:"cardinality,omitempty"`
	ParentTable       string   `json:"parent_table"`
	ParentColumns     []string `json:"parent_columns"`
	ParentCardinality string   `json:"parent_cardinality,omitempty"`
	Def               string   `json:"def"`
	Virtual           bool     `json:"virtual,omitempty"`
}

// Schema ...
type Schema struct {
	Name       string      `json:"name,omitempty"`
	Desc       string      `json:"desc,omitempty"`
	Tables     []Table     `json:"tables"`
	Relations  []Relation  `json:"relations,omitempty"`
	Functions  []Function  `json:"functions,omitempty"`
	Enums      []Enum      `json:"enums,omitempty"`
	Driver     *Driver     `json:"driver,omitempty"`
	Labels     []Label     `json:"labels,omitempty"`
	Viewpoints []Viewpoint `json:"viewpoints,omitempty"`
}

// Table ...
type Table struct {
	Name             string       `json:"name"`
	Type             string       `json:"type,omitempty"`
	Comment          string       `json:"comment,omitempty"`
	Columns          []Column     `json:"columns"`
	Indexes          []Index      `json:"indexes,omitempty"`
	Constraints      []Constraint `json:"constraints,omitempty"`
	Triggers         []Trigger    `json:"triggers,omitempty"`
	Def              string       `json:"def,omitempty"`
	Labels           []Label      `json:"labels,omitempty"`
	ReferencedTables []string     `json:"referenced_tables,omitempty"`
}

// Trigger ...
type Trigger struct {
	Name    string `json:"name"`
	Def     string `json:"def"`
	Comment string `json:"comment,omitempty"`
}

// Viewpoint ...
type Viewpoint struct {
	Name     string           `json:"name"`
	Desc     string           `json:"desc"`
	Labels   []string         `json:"labels,omitempty"`
	Tables   []string         `json:"tables,omitempty"`
	Distance int              `json:"distance,omitempty"`
	Groups   []ViewpointGroup `json:"groups,omitempty"`
}

// ViewpointGroup ...
type ViewpointGroup struct {
	Name   string   `json:"name"`
	Desc   string   `json:"desc"`
	Labels []string `json:"labels,omitempty"`
	Tables []string `json:"tables,omitempty"`
	Color  string   `json:"color,omitempty"`
}

// Viewpoints ...
//type Viewpoints []Viewpoint
