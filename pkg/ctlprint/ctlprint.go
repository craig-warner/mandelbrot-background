package ctlprint

import (
	"fmt"
)

type CtlPrint struct {
	verbose bool
	debug   bool
}

func NewCtlPrint(verbose, debug bool) CtlPrint {
	cp := CtlPrint{
		verbose: verbose,
		debug:   debug,
	}
	return cp
}

func (cp *CtlPrint) SetCltPrint(verbose, debug bool) {
	cp.verbose = verbose
	cp.debug = debug
}

func (cp *CtlPrint) DbgPrint(str ...interface{}) {
	if cp.debug {
		fmt.Println(str...)
		return
	}
}

func (cp *CtlPrint) VerbosePrint(str ...interface{}) {
	if cp.debug || cp.verbose {
		fmt.Println(str...)
		return
	}
}
