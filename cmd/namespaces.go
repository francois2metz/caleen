package cmd

import (
	"os"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List namespaces",
	Long:  `List Baleen namespaces`,
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("BALEEN_TOKEN")
		c := baleen.New(
			baleen.WithToken(token),
		)
		account, err := c.GetAccount()
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
