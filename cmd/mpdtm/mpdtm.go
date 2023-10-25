/*
シ MXIUM SYSTEMS シ
main.go
Main Entry Point of the application
*/

package main

import (
	"os"
	B "github.com/mxiumz/mpdtm/files"
)

func main() {

	// When help flags executed

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		B.FuncHelp()
		return
	}

	// Main Funtion entry Point
	B.Banmain()
	B.PdtsMain()
}
