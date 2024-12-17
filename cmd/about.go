package cmd

import (
	"github.com/gajevski/go-cli-aviation/utils"
	"github.com/liamg/gobless"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "A brief description of go-cli-aviation",
	Long:  "go-cli-aviation is a CLI tool for retrieving flight information from opensky-network",
	Run: func(cmd *cobra.Command, args []string) {
		render()
	},
}

func render() {
	helloTextbox := gobless.NewTextBox()
	helloTextbox.SetX(10)
	helloTextbox.SetY(3)
	helloTextbox.SetWidth(64)
	helloTextbox.SetHeight(10)
	helloTextbox.SetTextWrap(true)
	helloTextbox.SetText(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc gravida vulputate augue, ut lobortis neque semper at. In hac habitasse platea dictumst. In tincidunt diam vitae bibendum posuere. Maecenas imperdiet nunc non ex dignissim rhoncus. Integer consectetur turpis sit amet fermentum eleifend. Fusce placerat urna in purus efficitur dictum ut at nibh. Cras bibendum arcu eget ligula posuere, a tristique dolor mattis. Vivamus quis lectus cursus, eleifend leo sit amet, pellentesque dolor.`)
	utils.RenderGui(
		[]gobless.Component{
			helloTextbox,
		})
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
