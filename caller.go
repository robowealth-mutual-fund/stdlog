package stdlog

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var (
	stdlogPkg string
)

const (
	MAXIMUM_CALLER_DEPTH int = 3
)

func getCaller() string {
	var minimumCallerDepth int
	var callerInitOnce sync.Once
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(3, pcs)
		stdlogPkg = getPackageName(runtime.FuncForPC(pcs[0]).Name())
		minimumCallerDepth = MAXIMUM_CALLER_DEPTH
	})

	pcs := make([]uintptr, MAXIMUM_CALLER_DEPTH)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		if pkg != stdlogPkg {
			return fmt.Sprintf("%s: %s:%d", f.Func.Name(), f.File, f.Line)
		}
	}

	return ""
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
