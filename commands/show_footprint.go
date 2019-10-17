package commands

import (
	"fmt"
	"io"

	"github.com/micahyoung/cloudfoundry-footprint-plugin/queries"

	"code.cloudfoundry.org/cli/plugin"
)

//ShowFootprint writes footprint info
func ShowFootprint(cliConnection plugin.CliConnection, writer io.Writer) (err error) {
	var appsData []*queries.AppData
	if appsData, err = queries.GetAppsLastPushed(cliConnection); err != nil {
		return err
	}

	for _, appData := range appsData {
		if appData.AppLastUpdatedAt != "" {
			fmt.Fprintf(writer, "%s / %s (%s) last pushed on %s by %s\n", appData.AppName, appData.SpaceName, appData.AppState, appData.AppLastUpdatedAt, appData.AppLastUpdatedBy)

		} else {
			fmt.Fprintf(writer, "%s / %s (%s)\n", appData.AppName, appData.SpaceName, appData.AppState)
		}

	}
	return nil
}
