package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dynamoLayout struct {
	top      fyne.CanvasObject
	dividers [3]fyne.Widget
}

func newDynamoLayout(
	top fyne.CanvasObject,
	dividers [3]fyne.Widget,
) *dynamoLayout {
	return &dynamoLayout{
		top:      top,
		dividers: dividers,
	}
}

func (l *dynamoLayout) Layout(_ []fyne.CanvasObject, curSize fyne.Size) {
	topH := l.top.MinSize().Height
	sepTickness := theme.Size(theme.SizeNameSeparatorThickness)

	l.top.Move(fyne.NewPos(0, 0))
	l.top.Resize(fyne.NewSize(curSize.Width, topH))

	l.dividers[0].Move(fyne.NewPos(0, topH))
	l.dividers[0].Resize(fyne.NewSize(curSize.Width, sepTickness))
}

func (l *dynamoLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
