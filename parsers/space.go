package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type spaceWrapper struct {
	Entity APISpace `json:"entity"`
}

// APISpace data
type APISpace struct {
	Name string `json:"name"`
}

// ParseSpaceResponse parses json API response and returns data
func ParseSpaceResponse(spaceResponse []string) (apiSpace *APISpace, err error) {
	spaceResponseJSON := strings.Join(spaceResponse, "")

	var wrapper spaceWrapper
	err = json.Unmarshal([]byte(spaceResponseJSON), &wrapper)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#+v\n", wrapper)

	if (wrapper.Entity == APISpace{}) {
		return nil, errors.New("No space found")
	}
	apiSpace = &wrapper.Entity

	return apiSpace, nil
}
