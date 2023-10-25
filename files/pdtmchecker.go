/*
シ MXIUM SYSTEMS シ
pdtmchecker.go
Checks in pdtm in installed
*/

package files

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	C "github.com/fatih/color"
)

func PdtsMain() {
	PdtsChecker()
}

var (
	programName = "pdtm"
	programArgs = " "
	s           = `
[HINT] pdtm required for this tool to work
[HINT] go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
[HINT] If pdtm is already installed - source ~/.bashrc , exit shell and try agaain
[HINT] Check  - https://github.com/projectdiscovery/pdtm for further help

`
)

// Main Function
func PdtsChecker() {

	cr := C.New(C.FgRed).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
	cb := C.New(C.FgBlue).SprintFunc()

	cmd := exec.Command(programName, programArgs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logger := log.New(os.Stderr, "", 0)
	cmd.Stderr = logger.Writer()
	err := cmd.Run()

	if err != nil {
		logger.Printf(cr("\n[FAIL] %s"), err)
		fmt.Println(cr("[FAIL] PDTM - Project Discovery Open Source Tool Manager - not installed!\n" + cb(s)))

	} else {

		fmt.Println(chg("\n[OK] Ran PDTS Successfully!...Choose Install below..."))
		TableFuncMain()
	}

}
