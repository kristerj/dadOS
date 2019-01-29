package main


import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"./lib"
)

// Service is a function which returns the service's main primitive and its title.
// It receives a "nextService" function which can be called to advance the
// presentation to the next service.
type Service func(nextService func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

// Starting point for the application.
func main() {
	// The services (features this application provides).
	services := []Service{
		dadlib.Cover,
		dadlib.HelloWorld,
	}

	form := tview.NewForm()
	form.AddInputField("Login name", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", loginFn).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetWrap(false)

	// Create the pages for all services.
	currentService := 0
	info.Highlight(strconv.Itoa(currentService))
	pages := tview.NewPages()
	previousService := func() {
		currentService = (currentService - 1 + len(services)) % len(services)
		info.Highlight(strconv.Itoa(currentService)).
		ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentService))
	}
	nextService := func() {
		currentService = (currentService + 1) % len(services)
		info.Highlight(strconv.Itoa(currentService)).
		ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentService))
	}
	for index, service := range services {
		title, primitive := service(nextService)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentService)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}

	// Create the main layout.
	layout := tview.NewFlex().
	SetDirection(tview.FlexRow).
	AddItem(pages, 0, 1, true).
	AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextService()
		} else if event.Key() == tcell.KeyCtrlP {
			previousService()
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func loginFn(){

}
