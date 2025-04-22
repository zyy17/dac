package dashboard

import (
	"bytes"
	"fmt"
	"strings"
)

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

const (
	PlaceholderForEmpty = "--"
	BRTag               = "<br/>"
)

func (d *Dashboard) ToMarkdown() string {
	var buf bytes.Buffer

	for _, group := range d.Groups {
		buf.WriteString(fmt.Sprintf("# %s\n", group.Title))
		buf.WriteString("| Title | Query | Type | Description | Datasource | Unit | Legend Format |\n")
		buf.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
		for _, panel := range group.Panels {
			var query string
			if len(panel.Queries) > 0 {
				var queries []string
				for _, query := range panel.Queries {
					queries = append(queries, codeRef(query.Expr))
				}
				query = strings.Join(queries, BRTag)
			} else {
				query = codeRef(panel.Queries[0].Expr)
			}

			// Replace `\n` with `<br/>` for description.
			description := strings.ReplaceAll(panel.Description, "\n", BRTag)
			buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n", panel.Title, query, codeRef(panel.Type), description, codeRef(panel.Unit), codeRef(panel.Queries[0].Datasource.Type), codeRef(panel.Queries[0].LegendFormat)))
		}
	}

	return buf.String()
}

func codeRef(content string) string {
	if len(content) == 0 {
		return PlaceholderForEmpty
	}

	// Replace `|` with `\\|` for markdown.
	escaped := strings.ReplaceAll(content, "|", "\\|")

	return fmt.Sprintf("`%s`", escaped)
}
