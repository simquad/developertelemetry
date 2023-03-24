package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Definition []*cli.Command = []*cli.Command{
	{
		Name:  "get",
		Usage: "Get available workitems",
		Action: func(*cli.Context) error {
			fmt.Println("--workitems--")
			return nil
		},
	},
	{
		Name:  "add",
		Usage: "Add a workitem",
		Action: func(*cli.Context) error {
			fmt.Println("--addworkitem--")
			return nil
		},
	},
}
