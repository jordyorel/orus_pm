package cmd

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Manifest struct {
	Package struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	} `toml:"package"`
	Dependencies map[string]string `toml:"dependencies"`
}

func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m Manifest
	if err := toml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func SaveManifest(path string, m *Manifest) error {
	out, err := toml.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}
