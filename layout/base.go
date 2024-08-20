package layout

import "github.com/rivo/tview"

type Base struct {
	TextView  *tview.TextView
	FormInput *tview.TextArea
}

func BaseLayout(b *Base) *tview.Flex {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(b.TextView, 0, 1, false)
	flex.AddItem(b.FormInput, 6, 1, true)
	return flex
}
