package services

import (
	"github.com/rivo/tview"
)

//Service interface defines services(screens) provided
type Service interface {
	GetName() string
	GetContent() *tview.Flex
}
