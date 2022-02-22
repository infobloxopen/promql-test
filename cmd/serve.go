/*
Copyright © 2022 Infoblox

*/
package cmd

import (
	"fmt"

	"github.com/infobloxopen/promql-test/pkg/promock"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a prometheus server",
	Long: `Start a prometheus server:

Starts a prometheus server with test data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		server := promock.NewServer(nil)
		server.Serve(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().String("datafile", "", "the file that contains test data")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
