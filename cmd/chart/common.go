package chart

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func GetString(command *cobra.Command, Name string) string {
	ret, err := command.Flags().GetString(Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return ret
}

func GetInt(cmd *cobra.Command, Name string) int {
	ret, err := cmd.Flags().GetInt(Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ret
}
