package config

import (
	"os"

	baleen "github.com/francois2metz/steampipe-plugin-baleen/baleen/client"
)

func GetClient() *baleen.Client {
	token := os.Getenv("BALEEN_TOKEN")
	return baleen.New(
		baleen.WithToken(token),
	)
}
