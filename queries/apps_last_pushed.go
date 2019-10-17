package queries

import (
	"fmt"

	"github.com/micahyoung/cloudfoundry-footprint-plugin/parsers"

	"code.cloudfoundry.org/cli/plugin"
)

type AppData struct {
	AppName          string
	AppState         string
	SpaceName        string
	AppLastUpdatedBy string
	AppLastUpdatedAt string
}

//GetAppsLastPushed returns app data
func GetAppsLastPushed(cliConnection plugin.CliConnection) (appsData []*AppData, err error) {
	appsURL := "v2/apps"

	var appsResponse []string
	if appsResponse, err = cliConnection.CliCommandWithoutTerminalOutput("curl", appsURL); err != nil {
		return nil, err
	}

	var apiApps []*parsers.APIApp
	if apiApps, err = parsers.ParseAppsResponse(appsResponse); err != nil {
		return nil, err
	}

	for _, apiApp := range apiApps {
		appData := &AppData{}
		appData.AppName = apiApp.Name
		appData.AppState = apiApp.State

		spaceURL := fmt.Sprintf("v2/spaces/%s", apiApp.SpaceGUID)

		var spaceResponse []string
		if spaceResponse, err = cliConnection.CliCommandWithoutTerminalOutput("curl", spaceURL); err != nil {
			return nil, err
		}

		var apiSpace *parsers.APISpace
		if apiSpace, err = parsers.ParseSpaceResponse(spaceResponse); err != nil {
			return nil, err
		}

		appData.SpaceName = apiSpace.Name

		appUpdateEventURL := fmt.Sprintf("/v2/events?q=actee:%s&type:audit.app.update", apiApp.GUID)
		var appUpdateEventResponse []string
		if appUpdateEventResponse, err = cliConnection.CliCommandWithoutTerminalOutput("curl", appUpdateEventURL); err != nil {
			return nil, err
		}

		var apiAppUpdateEvent *parsers.APIEvent
		if apiAppUpdateEvent, err = parsers.ParseAppUpdateEventResponse(appUpdateEventResponse); err != nil {
			return nil, err
		}
		appData.AppLastUpdatedBy = apiAppUpdateEvent.LastUpdatedBy
		appData.AppLastUpdatedAt = apiAppUpdateEvent.LastUpdatedAt

		appsData = append(appsData, appData)
	}

	return appsData, nil
}
