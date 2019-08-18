package printer

import (
	"github.com/gofish-bot/gofish-bot/models"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Table(applications []*models.Application) {

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Org", "Repo", "CurrentVersion", "NewRelease", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, app := range applications {
		status := ""
		if app.CurrentVersion != "" && app.CurrentVersion != app.Version {
			status = "Needs update"
		} else if app.CurrentVersion == "" {
			status = "Missing"
		}
		tbl.AddRow(app.Organization, app.Name, app.CurrentVersion, app.Version, status)
	}

	tbl.Print()
}
