package parsers

import (
	"encoding/json"
	"errors"
	"strings"
)

type spaceWrapper struct {
	Entity struct {
		Name string `json:"name"`
	} `json:"entity"`
}

// APISpace data
type APISpace struct {
	Name string
}

// ParseSpaceResponse parses json API response and returns data
func ParseSpaceResponse(spaceResponse []string) (apiSpace *APISpace, err error) {
	spaceResponseJSON := strings.Join(spaceResponse, "")

	var wrapper spaceWrapper
	err = json.Unmarshal([]byte(spaceResponseJSON), &wrapper)
	if err != nil {
		return nil, err
	}

	if wrapper.Entity.Name == "" {
		return nil, errors.New("No space found")
	}

	apiSpace = &APISpace{
		Name: wrapper.Entity.Name,
	}

	return apiSpace, nil
}
