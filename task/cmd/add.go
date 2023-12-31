package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" in your task list\n", task )
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}

