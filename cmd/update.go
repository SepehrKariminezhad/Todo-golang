package cmd

import (
	"Todo_app/dbmanegment"
	"Todo_app/tools"
	"fmt"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
)

func CmdUpdate (mydb *dbmanegment.Mydb)*cobra.Command{
return &cobra.Command{
    Use:   "update [ID to Update] [Updated log]",
    Short: "Update the log for the given ID",
    Long: "",
	Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
    	id, err := strconv.ParseUint(args[0] , 10 , 64) 
			if err != nil {
				fmt.Println("Please provide a valid ID.")
				return
			}
    	tools.RemoveFromStr(args , 0)
		mydb.Update(id ,strings.Join(args, " "))
		fmt.Printf("Updated id %v\n" , id)
    	},
  	}
}
