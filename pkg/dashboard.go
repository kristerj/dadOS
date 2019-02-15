package dadlib

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)
      const (
	      subtitle   = `CO2 ppm`
	      navigation = `Ctrl-N: Next service    Ctrl-P: Previous service    Ctrl-C: Exit`
      )

      // Cover returns the cover page.
      func Dashboard(nextSlide func()) (title string, content tview.Primitive) {
	      // What's the size of the logo?
	      lines := strings.Split(bignine, "\n")
	      logoWidth := 0
	      logoHeight := len(lines)
	      for _, line := range lines {
		      if len(line) > logoWidth {
			      logoWidth = len(line)
		      }
	      }
	      logoBox := tview.NewTextView().
	      SetTextColor(tcell.ColorGreen).
	      SetDoneFunc(func(key tcell.Key) {
		      nextSlide()
	      })
	      fmt.Fprint(logoBox, Bignum(8))

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

	      return "Dashboard", flex
      }
