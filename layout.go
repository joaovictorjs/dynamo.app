package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dynamoLayout struct {
	top, topSeparator, content fyne.CanvasObject
}

func newDynamoLayout(top, topSeparator, content fyne.CanvasObject) *dynamoLayout {
	return &dynamoLayout{top, topSeparator, content}
}

func (l *dynamoLayout) Layout(_ []fyne.CanvasObject, currentSize fyne.Size) {
	topSize := l.top.MinSize()
	separatorTickness := theme.Size(theme.SizeNameSeparatorThickness)

	l.top.Move(fyne.NewPos(0, 0))
	l.top.Resize(fyne.NewSize(currentSize.Width, topSize.Height))

	l.topSeparator.Move(fyne.NewPos(0, topSize.Height))
	l.topSeparator.Resize(fyne.NewSize(currentSize.Width, separatorTickness))

	l.content.Move(fyne.NewPos(0, topSize.Height))
	l.content.Resize(fyne.NewSize(currentSize.Width, currentSize.Height-topSize.Height))
}

func (r *dynamoLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
