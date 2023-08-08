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

func (adapter *ModuleAdapter) CallFunc(
	funcName string,
	inputStr string, // json string
) (string, error) {
	// lock python runtime
	unlock := LockPythonRuntime()
	defer unlock() // unlock python runtime

	// get func
	funcObj := adapter.module.GetAttrString(funcName)
	if funcObj == nil {
		python3.PyErr_Print() // print python error
		return "", errors.New("failed to get func: " + funcName)
	}
	defer funcObj.DecRef()
	// check callable
	if !python3.PyCallable_Check(funcObj) {
		return "", errors.New("func is not callable: " + funcName)
	}

	// call func
	args := python3.PyTuple_New(1)
	defer args.DecRef()
	python3.PyTuple_SetItem(args, 0, python3.PyUnicode_FromString(inputStr))
	result := funcObj.Call(args, python3.Py_None)
	if result == nil {
		python3.PyErr_Print() // print python error
		return "", errors.New("failed to call func: " + funcName)
	}
	defer result.DecRef()

	// result should be a string
	if !python3.PyUnicode_Check(result) {
		return "", errors.New("result is not a string: " + funcName)
	}

	// get result
	outputStr := python3.PyUnicode_AsUTF8(result)
	return outputStr, nil
}
