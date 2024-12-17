package utils

import (
	"github.com/liamg/gobless"
)

var Gui = gobless.NewGUI()

func RenderGui(components []gobless.Component) {
	if err := Gui.Init(); err != nil {
		panic(err)
	}
	defer Gui.Close()
	for _, component := range components {
		Gui.Render(component)
	}

	Gui.HandleKeyPress(gobless.KeyCtrlQ, func(event gobless.KeyPressEvent) {
		Gui.Close()
	})

	Gui.Loop()
}

func RenderQuitTextbox() *gobless.TextBox {
	quitTextbox := gobless.NewTextBox()
	quitTextbox.SetWidth(22)
	quitTextbox.SetHeight(3)
	quitTextbox.SetText(`Press Ctrl-q to exit.`)
	return quitTextbox
}
