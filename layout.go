package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dynamoLayout struct {
	left, top fyne.CanvasObject
	dividers  [3]fyne.Widget
}

func newDynamoLayout(
	left, top fyne.CanvasObject,
	dividers [3]fyne.Widget,
) *dynamoLayout {
	return &dynamoLayout{
		left:     left,
		top:      top,
		dividers: dividers,
	}
}

func (l *dynamoLayout) Layout(_ []fyne.CanvasObject, currSize fyne.Size) {
	topH := l.top.MinSize().Height
	leftH := (currSize.Height * 0.65) - topH
	leftW := (currSize.Width * 0.25)
	sepTickness := theme.Size(theme.SizeNameSeparatorThickness)

	l.top.Move(fyne.NewPos(0, 0))
	l.top.Resize(fyne.NewSize(currSize.Width, topH))

	l.left.Move(fyne.NewPos(0, topH))
	l.left.Resize(fyne.NewSize(leftW, leftH))

	l.dividers[0].Move(fyne.NewPos(0, topH))
	l.dividers[0].Resize(fyne.NewSize(currSize.Width, sepTickness))

	l.dividers[1].Move(fyne.NewPos(leftW, topH))
	l.dividers[1].Resize(fyne.NewSize(sepTickness, leftH))
}

func (l *dynamoLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
