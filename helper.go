package gopy

import (
	"context"
	"runtime"

	python3 "github.com/go-python/cpy3"
)

// LockPythonRuntime : 锁定python3运行时
// 包括:
// 1. 锁定当前线程
// 2. 申请GIL, 确保当前线程成为python主线程
// 返回一个取消函数, 用于释放GIL和解锁当前线程
func LockPythonRuntime() context.CancelFunc {
	// 对GIL的操作需要在同一个线程中进行
	runtime.LockOSThread()
	// 申请GIL, 确保当前线程成为python主线程
	_gstate := python3.PyGILState_Ensure()
	return func() {
		python3.PyGILState_Release(_gstate)
		runtime.UnlockOSThread()
	}
}

func InsertPythonPaths(paths []string) {
	if len(paths) == 0 {
		return
	}
	// insert sys.path for import module
	sysModule := python3.PyImport_ImportModule("sys")
	defer sysModule.DecRef()
	pythonPath := sysModule.GetAttrString("path")
	defer pythonPath.DecRef()

	for _, p := range paths {
		python3.PyList_Insert(pythonPath, 0, python3.PyUnicode_FromString(p))
	}
}
