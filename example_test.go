package gopy

import (
	"encoding/json"
	"path"
	"runtime"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestCallEcho(t *testing.T) {
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

func TestCallAdd(t *testing.T) {
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

	input := struct {
		A int `json:"a"`
		B int `json:"b"`
	}{
		A: 1,
		B: 2,
	}
	inputStr, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	outputStr, err := adapter.CallFunc("add", string(inputStr))
	if err != nil {
		t.Fatal(err)
	}

	var output struct {
		Sum int `json:"sum"`
	}
	err = json.Unmarshal([]byte(outputStr), &output)
	if err != nil {
		t.Fatal(err)
	}
	if output.Sum != 3 {
		t.Fatalf("output.Sum: %d", output.Sum)
	}
}

func TestConcurrenceCall(t *testing.T) {
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

	eg := new(errgroup.Group)
	for i := 0; i < 3; i++ {
		eg.Go(func() error {
			_, err = adapter.CallFunc("print_numbers", "")
			if err != nil {
				return err
			}
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		t.Fatal(err)
	}
}
