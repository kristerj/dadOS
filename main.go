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
var savedLayout *tview.Flex
var serviceArr [10]services.Service
var titles [10]string
var info *tview.TextView
var console *tview.TextView
var currentService int
var indexupdate int
var currentTime string
var refreshCount int

// Starting point for the application.
func main() {
	refreshCount = 0
	numpages = 2
	console = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetScrollable(true)
	console.Box.SetBorder(true)
	currentTime = time.Now().Format("2 Jan 15:04:05")
	console.Write([]byte(currentTime + ": Console initialized\n"))

	// The services (features this application provides).
	serviceArr[0] = services.Dashboard{Name: "Dashboard 1", Content: nil, SharedVals: []int{1, 5, 2, 3, 4, 5, 6, 7, 8, 9}}
	serviceArr[1] = services.Dashboard{"Dashboard 2", nil, []int{2, 10, 2, 3, 4, 5, 6, 7, 8, 9}}
	bgChannel := make(chan string)
	titles[0] = "Dashboard 1"
	titles[1] = "Dashboard 2"
	go serviceArr[0].BackgroundTasks(bgChannel)
	go serviceArr[1].BackgroundTasks(bgChannel)
	// The bottom row has some info on where we are.
	info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)
	buildInfo(info)
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
		primitive := serviceArr[index].GetContent()
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentService)
	}
	// Create the main layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(console, 7, 0, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pressedKey := event.Key()
		if pressedKey == tcell.KeyCtrlN {
			nextService()
		} else if pressedKey == tcell.KeyCtrlP {
			previousService()
		} else if pressedKey == tcell.KeyCtrlL {
			showLogin()
		} else if pressedKey == tcell.KeyEsc || event.Rune() == 'q' {
			app.Stop()
		} else {
			serviceArr[currentService].PressKey(pressedKey)
		}
		return event
	})
	go updateConsole(bgChannel)
	go updateTime()
	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}

}

//end of main function, start of any helper functions
func showLogin() {
	savedLayout = layout
	form := tview.NewForm()
	form.AddInputField("Login name", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Login", func() { time.Sleep(2) }).
		AddButton("Quit", func() {
			app.SetRoot(savedLayout, true).Run()
		})
	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)
	app.SetRoot(layout, true)
}

func buildInfo(t *tview.TextView) {
	t.SetText(" ")
	for index := 0; index < numpages; index++ {
		title := titles[index]
		fmt.Fprintf(t, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}
	currentTime = time.Now().Format("Mon Jan 2 15:04:05")
	fmt.Fprintf(t, `/ %s`, currentTime)
	t.Highlight(strconv.Itoa(currentService)).
		ScrollToHighlight()

}

func updateConsole(c chan string) {
	for {
		ct := time.Now().Format("2 Jan 15:04:05")
		msg := <-c
		console.Write([]byte(ct + ": " + msg))
		console.ScrollToEnd()
	}
}

func updateTime() {
	for {
		time.Sleep(10 * time.Millisecond)
		app.QueueUpdateDraw(func() {
			buildInfo(info)
			primitive := serviceArr[currentService].GetContent()
			//there has to be a better way to redraw the current page than this
			pages.RemovePage(strconv.Itoa(currentService))
			pages.AddPage(strconv.Itoa(currentService), primitive, true, true)
		})
	}
}
