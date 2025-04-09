package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type responsiveLayout struct {
	top, topSeparator, content fyne.CanvasObject
}

func newResponsiveLayout(top, topSeparator, content fyne.CanvasObject) *responsiveLayout {
	return &responsiveLayout{top, topSeparator, content}
}

func (r *responsiveLayout) Layout(_ []fyne.CanvasObject, currentSize fyne.Size) {
	topSize := r.top.MinSize()
	separatorTickness := theme.Size(theme.SizeNameSeparatorThickness)

	r.top.Move(fyne.NewPos(0, 0))
	r.top.Resize(fyne.NewSize(currentSize.Width, topSize.Height))

	r.topSeparator.Move(fyne.NewPos(0, topSize.Height))
	r.topSeparator.Resize(fyne.NewSize(currentSize.Width, separatorTickness))

	r.content.Move(fyne.NewPos(0, topSize.Height))
	r.content.Resize(fyne.NewSize(currentSize.Width, currentSize.Height-topSize.Height))
}

func (r *responsiveLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
