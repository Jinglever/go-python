# go-python
### Call Python Functions from Go


## Runtime Environment
- Python 3.7.x
- pkg-config
- set PKG_CONFIG_PATH and LD_LIBRARY_PATH, e.g.
	```bash
	export PKG_CONFIG_PATH=/home/xxx/anaconda3/envs/py37/lib/pkgconfig:$PKG_CONFIG_PATH
	export LD_LIBRARY_PATH=/home/xxx/anaconda3/envs/py37/lib:$LD_LIBRARY_PATH
	```

## Usage
- import package
	```go
	import gopy "github.com/Jinglever/go-python"
	```
- call python function, refer to [example_test.go](example_test.go)
	```go
	cancelEnv, err := gopy.InitPythonEnv()
	if err != nil {
		panic(err)
	}
	defer cancelEnv()

	adapter, cancelAdapter, err := gopy.NewModuleAdapter(
		"my_module",
		[]string{"example"},
	)
	if err != nil {
		panic(err)
	}
	defer cancelAdapter()

	outputStr, err := adapter.CallFunc("my_func", `{"name": "gopy"}`)
	if err != nil {
		panic(err)
	}
	```


## Dev Environment
- set vscode setting, e.g.
	```json
	{
		"settings": {
		"go.toolsEnvVars": {
			"PATH": "/home/xxx/.gvm/gos/go1.19/bin:/home/xxx/anaconda3/envs/py37/bin:/usr/bin:${env:PATH}",
			"PKG_CONFIG_PATH": "/home/xxx/anaconda3/envs/py37/lib/pkgconfig:${env:PKG_CONFIG_PATH}",
			"LD_LIBRARY_PATH": "/home/xxx/anaconda3/envs/py37/lib:${env:LD_LIBRARY_PATH}"
		}
	}
	```