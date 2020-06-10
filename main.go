package main

import (
	"dados/services"
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()

var numpages int
var pages *tview.Pages
var layout *tview.Flex
var serviceArr [10]services.Service
var titles [10]string
var info *tview.TextView
var currentService int
var indexupdate int

// Starting point for the application.
func main() {
	numpages = 2
	dashboard := services.Dashboard{}
	// The services (features this application provides).
	serviceArr[0] = dashboard
	serviceArr[1] = dashboard

	titles[0] = "Dashboard 1"
	titles[1] = "Dashboard 2"
	// The bottom row has some info on where we are.
	info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	// Create the pages for all services.
	currentService = 0
	info.Highlight(strconv.Itoa(currentService))
	pages = tview.NewPages()

	previousService := func() {
		currentService = (currentService - 1 + numpages%numpages)
		info.Highlight(strconv.Itoa(currentService)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentService))
	}
	nextService := func() {
		currentService = (currentService + 1) % numpages
		info.Highlight(strconv.Itoa(currentService)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentService))
	}
	for index := 0; index < 2; index++ {
		title := titles[index]
		primitive := serviceArr[index].GetContent()
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentService)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}

	// Create the main layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextService()
		} else if event.Key() == tcell.KeyCtrlP {
			previousService()
		} else if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})
	go updateTime()
	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}

}

//end of main function, start of any helper functions
func showLogin() {
	form := tview.NewForm()
	form.AddInputField("Login name", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", loginFn).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)
}

func loginFn() {

}
func updateTime() {
	for {
		time.Sleep(2 * time.Second)
		app.QueueUpdateDraw(func() {
			primitive := serviceArr[currentService].GetContent()
			pages.RemovePage(strconv.Itoa(currentService))
			pages.AddPage(strconv.Itoa(currentService), primitive, true, true)
			layout = tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(pages, 0, 1, true).
				AddItem(info, 1, 1, false)
		})
	}
}
