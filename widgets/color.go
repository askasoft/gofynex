package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ImportanceToColorName(imp widget.Importance) (color fyne.ThemeColorName) {
	switch imp {
	case widget.LowImportance:
		color = theme.ColorNameDisabled
	case widget.MediumImportance:
		color = theme.ColorNameForeground
	case widget.HighImportance:
		color = theme.ColorNamePrimary
	case widget.DangerImportance:
		color = theme.ColorNameError
	case widget.WarningImportance:
		color = theme.ColorNameWarning
	case widget.SuccessImportance:
		color = theme.ColorNameSuccess
	default:
		color = theme.ColorNameForeground
	}
	return
}
