package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dynamoLayout struct {
	left, top, right, bottom fyne.CanvasObject
	dividers                 [3]fyne.Widget
}

func newDynamoLayout(
	left, top, right, bottom fyne.CanvasObject,
	dividers [3]fyne.Widget,
) *dynamoLayout {
	return &dynamoLayout{
		left:     left,
		top:      top,
		right:    right,
		bottom:   bottom,
		dividers: dividers,
	}
}

func (l *dynamoLayout) Layout(_ []fyne.CanvasObject, currSize fyne.Size) {
	topH := l.top.MinSize().Height
	leftH := (currSize.Height * 0.65) - topH
	leftW := (currSize.Width * 0.25)
	rightH := leftH
	rightW := currSize.Width - leftW
	bottomH := currSize.Height - leftH
	bottomW := currSize.Width
	sepTickness := theme.Size(theme.SizeNameSeparatorThickness)

	l.top.Move(fyne.NewPos(0, 0))
	l.top.Resize(fyne.NewSize(currSize.Width, topH))

	l.left.Move(fyne.NewPos(0, topH))
	l.left.Resize(fyne.NewSize(leftW, leftH))

	l.right.Move(fyne.NewPos(leftW, topH))
	l.right.Resize(fyne.NewSize(rightW, rightH))

	l.bottom.Move(fyne.NewPos(0, topH+leftH))
	l.bottom.Resize(fyne.NewSize(bottomW, bottomH))

	l.dividers[0].Move(fyne.NewPos(0, topH))
	l.dividers[0].Resize(fyne.NewSize(currSize.Width, sepTickness))

	l.dividers[1].Move(fyne.NewPos(leftW, topH))
	l.dividers[1].Resize(fyne.NewSize(sepTickness, leftH))

	l.dividers[2].Move(fyne.NewPos(0, topH+leftH))
	l.dividers[2].Resize(fyne.NewSize(bottomW, sepTickness))
}

func (l *dynamoLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
