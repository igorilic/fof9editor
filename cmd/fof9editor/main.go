package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/igorilic/fof9editor/internal/version"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var showVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version.GetVersionInfo())
		os.Exit(0)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow(fmt.Sprintf("FOF9 Editor v%s", version.GetShortVersion()))

	myWindow.SetContent(widget.NewLabel("Hello from FOF9 Editor"))
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}
