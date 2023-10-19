package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		for i, arg := range args{
			fmt.Printf("%d : %s\n", i, arg)
		}
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}

