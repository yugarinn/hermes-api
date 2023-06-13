package lib

import (
	"fmt"
)


type Hermes struct {
	dependencies map[string]interface{}
}

func (h *Hermes) Register(name string, dependency interface{}) {
    h.dependencies[name] = dependency
}

func (h *Hermes) Resolve(name string) (interface{}, error) {
    dependency, retrievalError := h.dependencies[name]

    if !retrievalError {
        return nil, fmt.Errorf("dependency_not_found: %s", name)
    }

    return dependency, nil
}
