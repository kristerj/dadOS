package services

import (
	"dados/dadlib"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const rightText = `[green]CO2 Setup Page`

//Co2 exports nextslide
func Co2(nextSlide func()) (title string, content tview.Primitive) {
	// We use a text view because we want to capture keyboard input.
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	textView.SetBorder(true).SetTitle("CO2")
	return "co2", dadlib.Code(textView, 30, 10, rightText)
}
