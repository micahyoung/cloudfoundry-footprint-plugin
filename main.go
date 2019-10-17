package main

import (
	"fmt"
	"os"

	"github.com/micahyoung/cloudfoundry-footprint-plugin/commands"

	"code.cloudfoundry.org/cli/plugin"
)

// FootprintPlugin is the Plugin implementation
type FootprintPlugin struct {
	TraceEnabled bool
}

// Run parses command line args and runs respective commands
func (p *FootprintPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	out := os.Stdout
	if args[0] == "footprint" {
		fmt.Println("Running the footprint command")
		if err := commands.ShowFootprint(cliConnection, out); err != nil {
			panic(err)
		}
	}
}

// GetMetadata for plugin
func (p *FootprintPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "CloudFoundryFootprintPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "footprint",
				HelpText: "Manage your org/space/app usage footprint",

				UsageDetails: plugin.Usage{
					Usage: "footprint\n   cf footprint",
				},
			},
		},
	}
}

func main() {

	plugin.Start(new(FootprintPlugin))
}
