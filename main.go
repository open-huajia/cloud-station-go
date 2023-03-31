package main

import (
	"fmt"
	"github.com/open-huajia/cloud-station-go/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
