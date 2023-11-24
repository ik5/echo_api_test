package runtimeutils

import (
	"fmt"
	"path/filepath"
	"strings"
)

// CallerInfo holds information regarding a given caller frame
type CallerInfo struct {
	Entry   uintptr
	LineNo  int
	Package string
	Func    string
	File    string
}

// PackageName returns only the Package name and ignoring the rest of the path.
func (ci CallerInfo) PackageName() string {
	return filepath.Base(ci.Package)
}

// FileName returns only the file name without the rest of the path.
func (ci CallerInfo) FileName() string {
	return filepath.Base(ci.File)
}

func (ci CallerInfo) String() string {
	return strings.Join([]string{ci.PackageName(), ci.Func}, ".")
}

// DebugInfo provides string with verbose information of each field
func (ci CallerInfo) DebugInfo() string {
	return fmt.Sprintf(
		"%s:%d->%s.%s", ci.File, ci.LineNo, ci.PackageName(), ci.Func,
	)
}
