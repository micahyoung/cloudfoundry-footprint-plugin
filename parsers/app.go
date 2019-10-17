package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type appsWrapper struct {
	Resources []struct {
		Entity struct {
			Name      string `json:"name"`
			State     string `json:"state"`
			SpaceGUID string `json:"space_guid"`
		} `json:"entity"`

		Metadata struct {
			GUID string `json:"guid"`
		} `json:"metadata"`
	} `json:"resources"`
}

// APIApp data
type APIApp struct {
	Name      string
	State     string
	SpaceGUID string
	GUID      string
}

// ParseAppsResponse parses json API response and returns data
func ParseAppsResponse(appsResponse []string) (apiApps []*APIApp, err error) {
	appsResponseJSON := strings.Join(appsResponse, "")

	var wrapper appsWrapper
	err = json.Unmarshal([]byte(appsResponseJSON), &wrapper)
	if err != nil {
		return nil, err
	}
	if len(wrapper.Resources) == 0 {
		return nil, errors.New("No apps found")
	}

	for _, appResource := range wrapper.Resources {
		fmt.Printf("resp: %#+v\n", appResource)

		apiApp := &APIApp{
			Name:      appResource.Entity.Name,
			State:     appResource.Entity.State,
			SpaceGUID: appResource.Entity.SpaceGUID,
			GUID:      appResource.Metadata.GUID,
		}
		apiApps = append(apiApps, apiApp)
	}

	return apiApps, nil
}
