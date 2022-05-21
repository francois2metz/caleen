package cmd

import (
	"fmt"
	"os"
	"strconv"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
	"github.com/spf13/cobra"
)

var (
	start string
	end string
	logsCmd = &cobra.Command{
		Use:   "logs",
		Short: "Get access logs",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			token := os.Getenv("BALEEN_TOKEN")
			c := baleen.New(
				baleen.WithToken(token),
			)
			startParam, err := strconv.Atoi(start)
			if err != nil {
				return;
			}
			endParam, err := strconv.Atoi(end)
			if err != nil {
				return;
			}
			params := baleen.AccessLogParams{
				Start: startParam,
				End:   endParam,
				Size:  100,
				Page:  0,
			}
			for {
				accessLogs, pagination, err := c.GetAccessLogs(namespace, params)
				if err != nil {
					fmt.Println(err)
					return
				}

				for _, accessLog := range accessLogs {
					fmt.Println(accessLog.Timestamp, accessLog.Status, accessLog.RequestMethod, accessLog.Scheme, accessLog.RequestUri)
				}
				params.Page += 1
				if params.Page > pagination.TotalCount / 100 {
					break;
				}

			}
		},
	}
)

func init() {
	rootCmd.AddCommand(logsCmd)

	logsCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The namespace")
	logsCmd.MarkPersistentFlagRequired("namespace")
	logsCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "Start timestamp")
	logsCmd.MarkPersistentFlagRequired("start")
	logsCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "End timestamp")
	logsCmd.MarkPersistentFlagRequired("end")
}
