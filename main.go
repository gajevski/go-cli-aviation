package main

import (
	"github.com/gajevski/go-cli-aviation/cmd"
	"github.com/gajevski/go-cli-aviation/utils"
	"github.com/liamg/gobless"
)

func main() {
	cmd.Execute()
	utils.RenderGui([]gobless.Component{
		utils.RenderQuitTextbox(),
	})
}
