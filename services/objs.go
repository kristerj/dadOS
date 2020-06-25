package services

import (
	"github.com/rivo/tview"
)

//Service interface defines services(screens) provided
type Service interface {
	GetSharedVals(int)int
	SetSharedVals(int, int)
	GetName() string
	GetContent() *tview.Flex
}
