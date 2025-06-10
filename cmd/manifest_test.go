package cmd

import (
	"os"
	"testing"
)

func TestLoadSaveManifest(t *testing.T) {
	tmp := "test_orus.toml"
	m := &Manifest{Dependencies: map[string]string{"dep": "1.0"}}
	m.Package.Name = "test"
	m.Package.Version = "0.1.0"
	if err := SaveManifest(tmp, m); err != nil {
		t.Fatalf("save: %v", err)
	}
	loaded, err := LoadManifest(tmp)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if loaded.Package.Name != m.Package.Name || loaded.Dependencies["dep"] != "1.0" {
		t.Fatalf("manifest mismatch")
	}
	os.Remove(tmp)
}
