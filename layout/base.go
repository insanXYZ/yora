package layout

import (
	"github.com/rivo/tview"
)

type ComponentLayout struct {
	TextView  *tview.TextView
	FormInput *tview.TextArea
	Hint      *tview.TextView
}

func BaseLayout(b *ComponentLayout) *tview.Flex {
	flex := tview.NewFlex()
	flex.AddItem(b.TextView, 0, 1, false)
	flex.AddItem(b.FormInput, 6, 1, true)
	flex.AddItem(b.Hint, 1, 1, false)
	flex.SetDirection(tview.FlexRow)
	return flex
}
