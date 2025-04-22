package transform

import (
	"github.com/zyy17/dac/pkg/dashboard"
	"github.com/zyy17/dac/pkg/grafana"
)

// Transform transforms a Grafana dashboard to a lightweight dashboard.
func Transform(grafanaDashboard *grafana.Dashboard) (*dashboard.Dashboard, error) {
	var (
		intermediateDashboard dashboard.Dashboard
		currentGroup          *dashboard.Group
		isRow                 = false
	)

	for _, panel := range grafanaDashboard.Panels {
		if panel.Type == grafana.PanelTypeRow {
			if len(panel.Panels) > 0 {
				if currentGroup != nil {
					intermediateDashboard.Groups = append(intermediateDashboard.Groups, currentGroup)
				}

				// Create a new group for the new row.
				currentGroup = &dashboard.Group{
					Title: panel.Title,
				}

				// Convert each sub-panel to a dashboard panel and add it to the group.
				for _, subPanel := range panel.Panels {
					if len(subPanel.Targets) == 0 {
						continue
					}

					if panel := subPanel.CovertToDashboardPanel(); panel != nil {
						currentGroup.Panels = append(currentGroup.Panels, panel)
					}
				}
			} else {
				if currentGroup != nil {
					intermediateDashboard.Groups = append(intermediateDashboard.Groups, currentGroup)
				}
				currentGroup = &dashboard.Group{
					Title: panel.Title,
				}
				isRow = true
				continue
			}
		}

		if isRow {
			if panel := panel.CovertToDashboardPanel(); panel != nil {
				currentGroup.Panels = append(currentGroup.Panels, panel)
			}
		}
	}

	if currentGroup != nil {
		intermediateDashboard.Groups = append(intermediateDashboard.Groups, currentGroup)
	}

	return &intermediateDashboard, nil
}
