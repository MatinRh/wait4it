package redis

import (
	"wait4it/pkg/model"
)

type checker struct {
	host          string
	port          int
	password      string
	database      int
	operationMode string
}

// NewChecker creates a new checker
func NewChecker(c *model.CheckContext) (model.CheckInterface, error) {
	checker := &checker{}
	checker.buildContext(*c)
	if err := checker.validate(); err != nil {
		return nil, err
	}

	return checker, nil
}
