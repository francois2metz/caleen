package cmd

import (
	"os"

	"github.com/francois2metz/caleen/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List namespaces",
	Long:  `List Baleen namespaces`,
	Run: func(cmd *cobra.Command, args []string) {
		client := config.GetClient()
		account, err := client.GetAccount()
		if err != nil {
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Namespace", "Name"})

		for _, namespace := range account.Namespaces {
			table.Append([]string{namespace.ID, namespace.Name})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
}
