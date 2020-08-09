package services

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

//Service interface defines services(screens) provided
type Service interface {
	GetSharedVals(int) int
	SetSharedVals(int, int)
	GetName() string
	GetContent() *tview.Flex
	PressKey(tcell.Key)
	BackgroundTasks(c chan string)
}
