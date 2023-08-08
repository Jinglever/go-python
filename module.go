package gopy

import (
	"context"
	"errors"

	python3 "github.com/go-python/cpy3"
)

// ModuleAdapter : 模块接合器
type ModuleAdapter struct {
	module *python3.PyObject
}

// NewModuleAdapter : 创建模块接合器
// moduleName: 模块名称
// pythonPaths: 需要插入到python3环境的sys.path中的路径列表
// 返回一个模块接合器, 以及一个取消函数, 用于释放模块接合器
func NewModuleAdapter(
	moduleName string,
	pythonPaths []string,
) (*ModuleAdapter, context.CancelFunc, error) {
	adapter := &ModuleAdapter{}
	cancel := func() {
		if adapter.module != nil {
			adapter.module.DecRef()
			adapter.module = nil
		}
	}

	// lock python runtime
	unlock := LockPythonRuntime()
	defer unlock() // unlock python runtime

	// insert python paths
	InsertPythonPaths(pythonPaths)

	// import module
	adapter.module = python3.PyImport_ImportModule(moduleName)
	if adapter.module == nil {
		python3.PyErr_Print() // print python error
		return nil, nil, errors.New("failed to import module: " + moduleName)
	}
	return adapter, cancel, nil
}
