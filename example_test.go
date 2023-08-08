package gopy

import (
	"path"
	"runtime"
	"testing"
)

func TestModuleAdapterCallFunc(t *testing.T) {
	cancelEnv, err := InitPythonEnv()
	if err != nil {
		t.Fatal(err)
	}
	defer cancelEnv()

	_, filename, _, _ := runtime.Caller(0)
	thisPath := path.Dir(filename)
	adapter, cancelAdapter, err := NewModuleAdapter(
		"example",
		[]string{path.Join(thisPath, "example")},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer cancelAdapter()

	outputStr, err := adapter.CallFunc("echo", `{"name": "gopy"}`)
	if err != nil {
		t.Fatal(err)
	}
	if outputStr != `{"name": "gopy"}` {
		t.Fatalf("outputStr: %s", outputStr)
	}
}
