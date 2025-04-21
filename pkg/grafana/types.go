package grafana

import (
	"github.com/zyy17/dac/pkg/dashboard"
)

// Dashboard is a JSON model of Grafana dashboard.
// Since the Grafana dashboard JSON model is not a open standard, we don't need to define every field here.
type Dashboard struct {
	Panels               []*Panel       `json:"panels,omitempty"`
	Annotations          map[string]any `json:"annotations,omitempty"`
	Description          string         `json:"description,omitempty"`
	Editable             bool           `json:"editable,omitempty"`
	FiscalYearStartMonth int            `json:"fiscalYearStartMonth,omitempty"`
	GraphTooltip         int            `json:"graphTooltip,omitempty"`
	ID                   int            `json:"id,omitempty"`
	Links                []any          `json:"links,omitempty"`
	LiveNow              bool           `json:"liveNow,omitempty"`
	Refresh              string         `json:"refresh,omitempty"`
	SchemaVersion        int            `json:"schemaVersion,omitempty"`
	Tags                 []string       `json:"tags,omitempty"`
	Templating           map[string]any `json:"templating,omitempty"`
	Time                 map[string]any `json:"time,omitempty"`
	Timepicker           map[string]any `json:"timepicker,omitempty"`
	Timezone             string         `json:"timezone,omitempty"`
	Title                string         `json:"title,omitempty"`
	UID                  string         `json:"uid,omitempty"`
	Version              int            `json:"version,omitempty"`
	WeekStart            string         `json:"weekStart,omitempty"`
}

type Panel struct {
	ID            int        `json:"id,omitempty"`
	Title         string     `json:"title,omitempty"`
	GridPos       GridPos    `json:"gridPos,omitempty"`
	Panels        []*Panel   `json:"panels,omitempty"`
	Type          string     `json:"type,omitempty"`
	Description   string     `json:"description,omitempty"`
	Datasource    Datasource `json:"datasource,omitempty"`
	PluginVersion string     `json:"pluginVersion,omitempty"`
	Targets       []*Target  `json:"targets,omitempty"`

	// Treat the following as opaque structs.
	FieldConfig *FieldConfig   `json:"fieldConfig,omitempty"`
	Options     map[string]any `json:"options,omitempty"`
}

const (
	PanelTypeRow = "row"
)

type FieldConfig struct {
	Defaults  *FieldConfigDefaults `json:"defaults,omitempty"`
	Overrides []any                `json:"overrides,omitempty"`
}

type FieldConfigDefaults struct {
	Unit string `json:"unit,omitempty"`

	// Ignore the following fields that are not needed.
}

type GridPos struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	W int `json:"w,omitempty"`
	H int `json:"h,omitempty"`
}

type Datasource struct {
	Type string `json:"type,omitempty"`
	UID  string `json:"uid,omitempty"`
}

type Target struct {
	Datasource   *Datasource    `json:"datasource,omitempty"`
	Expr         string         `json:"expr,omitempty"`
	Instant      bool           `json:"instant,omitempty"`
	LegendFormat string         `json:"legendFormat,omitempty"`
	Range        bool           `json:"range,omitempty"`
	RefID        string         `json:"refId,omitempty"`
	UseBackend   bool           `json:"useBackend,omitempty"`
	EditorMode   string         `json:"editorMode,omitempty"`
	Format       string         `json:"format,omitempty"`
	RawQuery     bool           `json:"rawQuery,omitempty"`
	RawSql       string         `json:"rawSql,omitempty"`
	Sql          map[string]any `json:"sql,omitempty"`
}

func (p *Panel) GetUnit() string {
	if p.FieldConfig != nil &&
		p.FieldConfig.Defaults != nil &&
		len(p.FieldConfig.Defaults.Unit) > 0 {
		return p.FieldConfig.Defaults.Unit
	}
	return ""
}

func (p *Panel) CovertToDashboardPanel() *dashboard.Panel {
	if p.Type == PanelTypeRow || len(p.Targets) == 0 {
		return nil
	}

	panel := &dashboard.Panel{
		Title:       p.Title,
		Unit:        p.GetUnit(),
		Type:        p.Type,
		Description: p.Description,
	}

	for _, target := range p.Targets {
		var expr string
		if target.Expr != "" {
			expr = target.Expr
		} else {
			expr = target.RawSql
		}
		panel.Queries = append(panel.Queries, &dashboard.Query{
			Expr:         expr,
			LegendFormat: target.LegendFormat,
			Datasource: &dashboard.Datasource{
				Type: target.Datasource.Type,
				UID:  target.Datasource.UID,
			},
		})
	}

	return panel
}
