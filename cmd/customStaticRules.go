package cmd

import (
	"os"
	"strconv"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var namespace string

var customStaticRulesCmd = &cobra.Command{
	Use:   "custom-static-rules",
	Short: "List custom static rules",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("BALEEN_TOKEN")
		c := baleen.New(
			baleen.WithToken(token),
		)
		rules, err := c.GetCustomStaticRules(namespace)
		if err != nil {
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Category", "Enabled"})

		for _, rule := range rules {
			table.Append([]string{rule.ID, rule.Category, strconv.FormatBool(rule.Enabled)})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(customStaticRulesCmd)

	customStaticRulesCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The namespace")
	customStaticRulesCmd.MarkPersistentFlagRequired("namespace")
}
