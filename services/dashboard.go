package services

import (
	"dados/dadlib"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	subtitle   = `CO2 ppm`
	navigation = `Ctrl-N: Next service    Ctrl-P: Previous service    Ctrl-C: Exit`
)

//Dashboard is a thing
type Dashboard struct {
	Name    string
	content *tview.Flex
}

//GetContent spits out the content
func (d Dashboard) GetContent() *tview.Flex {
	// What's the size of the logo?
	somenum := rand.Intn(9)
	lines := strings.Split(dadlib.Bignum(somenum), "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen)

	fmt.Fprint(logoBox, dadlib.Bignum(somenum))

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return flex
}

//GetName spits out a name
func (d Dashboard) GetName() (name string) {
	return "Dashboard"
}
