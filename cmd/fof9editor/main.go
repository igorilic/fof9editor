package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/igorilic/fof9editor/internal/ui"
	"github.com/igorilic/fof9editor/internal/version"

	"fyne.io/fyne/v2/app"
)

var showVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version.GetVersionInfo())
		os.Exit(0)
	}

	myApp := app.New()
	mainWindow := ui.NewMainWindow(myApp)

	mainWindow.ShowAndRun()
}
