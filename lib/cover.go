package dadlib

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const bigzero = `
    000000000     
  00:::::::::00   
00:::::::::::::00 
0:::::::000:::::::0
0::::::0   0::::::0
0:::::0     0:::::0
0:::::0     0:::::0
0:::::0 000 0:::::0
0:::::0 000 0:::::0
0:::::0     0:::::0
0:::::0     0:::::0
0::::::0   0::::::0
0:::::::000:::::::0
 00:::::::::::::00 
   00:::::::::00   
     000000000     
`

const bigone = `
  1111111   
 1::::::1   
1:::::::1   
111:::::1   
   1::::1   
   1::::1   
   1::::1   
   1::::l   
   1::::l   
   1::::l   
   1::::l   
   1::::l   
111::::::111
1::::::::::1
1::::::::::1
111111111111
`

const bigtwo = `
222222222222222    
2:::::::::::::::22  
2::::::222222:::::2 
2222222     2:::::2 
            2:::::2 
            2:::::2 
         2222::::2  
    22222::::::22   
   22::::::::222     
2:::::22222        
2:::::2             
2:::::2             
2:::::2       222222
2::::::2222222:::::2
2::::::::::::::::::2
22222222222222222222
`
const bigthree = `
 333333333333333   
3:::::::::::::::33 
3::::::33333::::::3
3333333     3:::::3
            3:::::3
            3:::::3
    33333333:::::3 
    3:::::::::::3  
    33333333:::::3 
            3:::::3
            3:::::3
            3:::::3
3333333     3:::::3
3::::::33333::::::3
3:::::::::::::::33 
 333333333333333   
`
const bigfour = `
       444444444  
      4::::::::4  
     4:::::::::4  
    4::::44::::4  
   4::::4 4::::4  
  4::::4  4::::4  
 4::::4   4::::4  
4::::444444::::444
4::::::::::::::::4
4444444444:::::444
          4::::4  
          4::::4  
          4::::4  
        44::::::44
        4::::::::4
        4444444444
`
const bigfive = `
555555555555555555 
5::::::::::::::::5 
5::::::::::::::::5 
5:::::555555555555 
5:::::5            
5:::::5            
5:::::5555555555   
5:::::::::::::::5  
555555555555:::::5 
            5:::::5
            5:::::5
5555555     5:::::5
5::::::55555::::::5
 55:::::::::::::55 
    55:::::::::55   
      555555555     
`
const bigsix = `
        66666666   
       6::::::6    
      6::::::6     
     6::::::6      
    6::::::6       
   6::::::6        
  6::::::6         
 6::::::::66666    
6::::::::::::::66  
6::::::66666:::::6 
6:::::6     6:::::6
6:::::6     6:::::6
6::::::66666::::::6
 66:::::::::::::66 
   66:::::::::66   
     666666666     
`
const bigseven = `
77777777777777777777
7::::::::::::::::::7
7::::::::::::::::::7
777777777777:::::::7
           7::::::7 
          7::::::7  
         7::::::7   
        7::::::7    
       7::::::7     
      7::::::7      
     7::::::7       
    7::::::7        
   7::::::7         
  7::::::7          
 7::::::7           
77777777          
`
const bigeight = `
     888888888     
   88:::::::::88   
 88:::::::::::::88 
8::::::88888::::::8
8:::::8     8:::::8
8:::::8     8:::::8
 8:::::88888:::::8 
  8:::::::::::::8  
 8:::::88888:::::8 
8:::::8     8:::::8
8:::::8     8:::::8
8:::::8     8:::::8
8::::::88888::::::8
 88:::::::::::::88 
   88:::::::::88   
     888888888
`
const bignine = `
     999999999     
   99:::::::::99   
 99:::::::::::::99 
9::::::99999::::::9
9:::::9     9:::::9
9:::::9     9:::::9
 9:::::99999::::::9
  99::::::::::::::9
    99999::::::::9 
         9::::::9  
        9::::::9   
       9::::::9    
      9::::::9     
     9::::::9      
    9::::::9       
   99999999    
`
      const (
	      subtitle   = `CO2 ppm`
	      navigation = `Ctrl-N: Next service    Ctrl-P: Previous service    Ctrl-C: Exit`
      )

      // Cover returns the cover page.
      func Cover(nextSlide func()) (title string, content tview.Primitive) {
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
	      fmt.Fprint(logoBox, bignine)

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

	      return "Start", flex
      }
