package ctlprint

import (
	"fmt"

	"github.com/fatih/color"
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
		fmt.Print(color.CyanString("Debug:"))
		fmt.Println(str...)
		return
	}
}

func (cp *CtlPrint) InfoPrint(str ...interface{}) {
	if cp.debug || cp.verbose {
		fmt.Print(color.BlueString("Info:"))
		fmt.Println(str...)
		return
	}
}

func (cp *CtlPrint) DonePrint(str ...interface{}) {
	if cp.debug || cp.verbose {
		fmt.Print(color.GreenString("Done:"))
		fmt.Println(str...)
	}
}

func (cp *CtlPrint) WarningPrint(str ...interface{}) {
	fmt.Print(color.YellowString("Warning:"))
	fmt.Println(str...)
}

func (cp *CtlPrint) ErrorPrint(str ...interface{}) {
	fmt.Print(color.RedString("Error:"))
	fmt.Println(str...)
}
