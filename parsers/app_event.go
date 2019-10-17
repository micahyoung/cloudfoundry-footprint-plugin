package parsers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type eventsWrapper struct {
	Resources []struct {
		Entity struct {
			ActorUsername string `json:"actor_username"`
		} `json:"entity"`

		Metadata struct {
			CreatedAt string `json:"created_at"`
		} `json:"metadata"`
	} `json:"resources"`
}

// APIEvent data
type APIEvent struct {
	LastUpdatedBy string
	LastUpdatedAt string
}

// ParseAppUpdateEventResponse parses json API response and returns data
func ParseAppUpdateEventResponse(appUpdateEventResponse []string) (apiEvent *APIEvent, err error) {
	appUpdateEventResponseJSON := strings.Join(appUpdateEventResponse, "")

	var wrapper eventsWrapper
	err = json.Unmarshal([]byte(appUpdateEventResponseJSON), &wrapper)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#+v\n", wrapper)

	if len(wrapper.Resources) == 0 {
		return &APIEvent{}, nil
	}

	apiEvent = &APIEvent{}
	for _, eventResource := range wrapper.Resources {
		fmt.Printf("event: %#+v\n", eventResource)

		apiEvent.LastUpdatedBy = eventResource.Entity.ActorUsername
		apiEvent.LastUpdatedAt = eventResource.Metadata.CreatedAt
	}

	return apiEvent, nil
}
