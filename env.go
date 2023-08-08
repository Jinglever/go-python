package gopy

import (
	"context"
	"errors"

	python3 "github.com/go-python/cpy3"
)

var state *python3.PyThreadState

// InitPythonEnv : 初始化python3环境
// Attention: 由于python3的GIL机制，
// python3环境只能在主线程中初始化，否则会导致python3.Py_Initialize()失败；
// 该函数返回一个取消函数，用于释放python3环境；
// 注意：如果没有在其它地方初始化python3环境，那么应该调用该函数完成初始化，
// 并且本函数只能调用一次。
func InitPythonEnv() (context.CancelFunc, error) {
	// The following will also create the GIL explicitly
	// by calling PyEval_InitThreads(), without waiting
	// for the interpreter to do that
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		return nil, errors.New("python3.Py_Initialize failed")
	}
	// Initialize() has locked the the GIL but at this point we don't need it
	// anymore. We save the current state and release the lock
	// so that goroutines can acquire it
	state = python3.PyEval_SaveThread()
	return func() {
		// At this point we know we won't need Python anymore in this
		// program, we can restore the state and lock the GIL to perform
		// the final operations before exiting.
		python3.PyEval_RestoreThread(state)

		python3.Py_Finalize()
	}, nil
}
