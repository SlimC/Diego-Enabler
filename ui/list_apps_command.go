package ui

import (
	"fmt"

	"github.com/cloudfoundry/cli/cf/terminal"
)

type ListAppsCommand struct {
	Username     string
	Runtime      Runtime
	Organization string
	Space        string
	UI           terminal.UI
}

func (c *ListAppsCommand) BeforeAll() {
	switch {
	case c.Space != "" && c.Organization != "":
		fmt.Printf(
			"Getting apps on the %s runtime in org %s / %s as %s...\n",
			terminal.EntityNameColor(c.Runtime.String()),
			terminal.EntityNameColor(c.Organization),
			terminal.EntityNameColor(c.Space),
			terminal.EntityNameColor(c.Username),
		)
	case c.Organization != "":
		fmt.Printf(
			"Getting apps on the %s runtime in org %s as %s...\n",
			terminal.EntityNameColor(c.Runtime.String()),
			terminal.EntityNameColor(c.Organization),
			terminal.EntityNameColor(c.Username),
		)
	default:
		fmt.Printf(
			"Getting apps on the %s runtime as %s...\n",
			terminal.EntityNameColor(c.Runtime.String()),
			terminal.EntityNameColor(c.Username),
		)
	}
}

func (c *ListAppsCommand) AfterAll(apps []ApplicationPrinter) {
	SayOK()

	headers := []string{
		"name",
		"space",
		"org",
	}
	t := terminal.NewTable(c.UI, headers)

	for _, app := range apps {
		t.Add(app.Name(), app.Space(), app.Organization())
	}

	t.Print()
}
