package dashboard

// Dashboard is a lightweight dashboard definition.
type Dashboard struct {
	// Groups contains a list of Group. For each group, it contains a list of Panels.
	Groups []*Group `json:"groups,omitempty" yaml:"groups,omitempty"`

	// Panels contains a list of Panel.
	Panels []*Panel `json:"panels,omitempty" yaml:"panels,omitempty"`
}

// Group is a group of panels. It's like a row in Grafana.
type Group struct {
	// Title is the title of the group.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Panels contains a list of Panel.
	Panels []*Panel `json:"panels,omitempty" yaml:"panels,omitempty"`
}

// Panel is a lightweight panel definition that contains a list of queries.
type Panel struct {
	// Title is the title of the panel.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Type is the type of the panel.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description is the description of the panel.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unit is the unit of the panel.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Queries contains a list of Query.
	Queries []*Query `json:"queries,omitempty" yaml:"queries,omitempty"`
}

// Query is a query definition that contains a list of queries.
type Query struct {
	// Expr is the expression of the query.
	Expr string `json:"expr,omitempty" yaml:"expr,omitempty"`

	// Datasource is the datasource of the query.
	Datasource *Datasource `json:"datasource,omitempty" yaml:"datasource,omitempty"`

	// LegendFormat is the legend format of the query.
	LegendFormat string `json:"legendFormat,omitempty" yaml:"legendFormat,omitempty"`
}

// Datasource is the datasource of the query.
type Datasource struct {
	// Type is the type of the datasource.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// UID is the uid of the datasource.
	UID string `json:"uid,omitempty" yaml:"uid,omitempty"`
}
