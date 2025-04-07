package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dynamoTheme struct {
	fyne.Theme
}

func newDynamoTheme() *dynamoTheme {
	return &dynamoTheme{Theme: theme.DefaultTheme()}
}

func (t *dynamoTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameShadow {
		return color.Transparent
	}

	if name == theme.ColorNameSeparator {
		return color.RGBA{54, 54, 54, 255}
	}

	return t.Theme.Color(name, theme.VariantDark)
}

func (t *dynamoTheme) Font(name fyne.TextStyle) fyne.Resource {
	return t.Theme.Font(name)
}

func (t *dynamoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.Theme.Icon(name)
}

func (t *dynamoTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}
