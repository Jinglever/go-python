# go-python
### Call Python Functions from Go


## Environment
- Python 3.7.x
- pkg-config
- set PKG_CONFIG_PATH and LD_LIBRARY_PATH, e.g.
	```bash
	export PKG_CONFIG_PATH=/home/howard/anaconda3/envs/py37/lib/pkgconfig:$PKG_CONFIG_PATH
	export LD_LIBRARY_PATH=/home/howard/anaconda3/envs/py37/lib:$LD_LIBRARY_PATH
	```

## Dev Environment
- set vscode setting, e.g.
	```json
	{
		"settings": {
		"go.toolsEnvVars": {
			"PATH": "/home/howard/.gvm/gos/go1.19/bin:/home/howard/anaconda3/envs/py37/bin:/usr/bin:${env:PATH}",
			"PKG_CONFIG_PATH": "/home/howardnaconda3/envs/py37/lib/pkgconfig:${env:PKG_CONFIG_PATH}",
			"LD_LIBRARY_PATH": "/home/howard/anaconda3/envs/py37/lib:${env:LD_LIBRARY_PATH}"
		}
	}
	```