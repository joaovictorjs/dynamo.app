package main

import "fyne.io/fyne/v2"

type responsiveLayout struct {
	top, center, bottom fyne.CanvasObject
}

func newResponsiveLayout(top, center, bottom fyne.CanvasObject) *responsiveLayout {
	return &responsiveLayout{top, center, bottom}
}

func (r *responsiveLayout) Layout(_ []fyne.CanvasObject, currentSize fyne.Size) {

}

func (r *responsiveLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
