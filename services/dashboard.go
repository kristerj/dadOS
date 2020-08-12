package services

import (
	"dados/dadlib"
	"fmt"
	//	"math/rand"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	subtitle   = `dadOS Dashboard`
	navigation = `Ctrl-N: Next service    Ctrl-P: Previous service    Ctrl-C: Exit`
)

//Dashboard is a thing
type Dashboard struct {
	Name       string
	Content    *tview.Flex
	SharedVals []int
}

var showButtons bool

func (d Dashboard) GetSharedVals(i int) int {
	return d.SharedVals[i]
}
func (d Dashboard) SetSharedVals(i int, j int) {
	d.SharedVals[i] = j
}

func (d Dashboard) BackgroundTasks(c chan string) {
	for {
		c <- ("Service " + d.Name + " check in\n")
		time.Sleep(time.Duration(d.SharedVals[1]) * time.Second)
	}
}

//GetContent spits out the content
func (d Dashboard) GetContent() *tview.Flex {
	d.SetSharedVals(0, d.GetSharedVals(0)+1)
	//fmt.Printf(".%d.\n", d.SharedVals[0])
	// What's the size of the logo?
	lines := strings.Split(dadlib.Bignum(d.GetSharedVals(0)%9), "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen)

	fmt.Fprint(logoBox, dadlib.Bignum(d.GetSharedVals(0)%9))

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow)
	if showButtons {
		flex.AddItem(tview.NewForm().AddButton("FIRE LASERS", func() {}), 0, 3, false)
	}
	flex.AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 2, false)

	return flex
}

//GetName spits out a name
func (d Dashboard) GetName() (name string) {
	return "Dashboard"
}

func (d Dashboard) PressKey(k tcell.Key) {
	if k == tcell.KeyF1 {
		showButtons = !showButtons
	} else if k == tcell.KeyF2 {

	} else if k == tcell.KeyF3 {

	} else if k == tcell.KeyF4 {

	} else if k == tcell.KeyF5 {

	} else if k == tcell.KeyF6 {

	} else if k == tcell.KeyF7 {

	} else if k == tcell.KeyF8 {

	} else if k == tcell.KeyF9 {

	} else if k == tcell.KeyF10 {

	}
}
