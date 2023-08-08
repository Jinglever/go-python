package gopy

import (
	"path"
	"runtime"
	"testing"
)

func TestNewModuleAdapter(t *testing.T) {
	cancelEnv, err := InitPythonEnv()
	if err != nil {
		t.Fatal(err)
	}
	defer cancelEnv()

	_, filename, _, _ := runtime.Caller(0)
	thisPath := path.Dir(filename)
	_, cancelAdapter, err := NewModuleAdapter(
		"example",
		[]string{path.Join(thisPath, "example")},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer cancelAdapter()
}
