package toml

import (
	"bytes"
	"fmt"
	"testing"
)

type Config struct {
	Tools map[string]Entry             `toml:"tools"`
	Env   map[string]map[string]string `toml:"env"`
}

type Entry struct {
	Version string            `toml:"version"`
	IsLink  bool              `toml:"is_link"`
	Env     map[string]string `toml:"env"`
}

func TestNewVfoxToml(t *testing.T) {
	//cfg := Config{
	//	Tools: map[string]Entry{
	//		"python": {
	//			Version: "3.8.0",
	//			IsLink:  true,
	//			Env: map[string]string{
	//				"PATH":            "/usr/local/bin",
	//				"LD_LIBRARY_PATH": "/usr/local/lib",
	//			},
	//		},
	//	},
	//	//Env: map[string]map[string]string{
	//	//	"python": {
	//	//		"PATH":            "/usr/local/bin",
	//	//		"LD_LIBRARY_PATH": "/usr/local/lib",
	//	//	},
	//	//},
	//}
	cfg := map[string]map[string]Entry{
		"python": {
			"path": {
				Version: "3.8.0",
				IsLink:  true,
				Env: map[string]string{
					"PATH": "/usr/local/bin",
				},
			},
		},
		"nodejs": {
			"path": {
				Version: "3.8.0",
				IsLink:  true,
				Env: map[string]string{
					"PATH": "/usr/local/bin",
				},
			},
		},
	}

	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	enc.SetRootInline(true)
	//enc.SetTablesInline(true)

	err := enc.Encode(cfg)
	if err != nil {
		panic(err)
	}

	//data, err := Marshal(cfg)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(string(buf.String()))

}
