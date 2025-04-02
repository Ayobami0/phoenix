package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseAsh(filename string) (*Ash, error) {
	data, err := os.ReadFile(filename)
	var ash Ash

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &ash); err != nil { return nil, err
	}

	return &ash, nil
}
