package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var aboveMeCmd = &cobra.Command{
	Use:   "aboveMe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: greet,
}

func greet(cmd *cobra.Command, args []string) {
	fmt.Println("above me called")
	fetchFlightsAboveMeData()
}

func fetchFlightsAboveMeData() {
	url := "https://opensky-network.org/api/states/all?lamin=50.401&lomin=19.382&lamax=50.459&lomax=19.515"

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching opensky-network data %v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error fetching opensky-network data %v", err)
	}

	fmt.Printf("data:\n%s\n", string(body))
}

func init() {
	rootCmd.AddCommand(aboveMeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboveMeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboveMeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
