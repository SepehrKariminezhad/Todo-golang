package cmd

import (
	"Todo_app/dbmanegment"
	"Todo_app/tools"
	"fmt"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
)

func CmdInsert(mydb *dbmanegment.Mydb)*cobra.Command{
	return &cobra.Command{
			Use:"insert [a unique ID] [log to insert]",
    		Short: "Create a row in your DB",
    		Long: "",
    		Args: cobra.MinimumNArgs(2),
    		Run: func(cmd *cobra.Command, args []string) {
    		id, err := strconv.ParseUint(args[0] , 10 , 64) 
			if err != nil {
				fmt.Println("Please provide a valid ID.")
				return
			}
    		tools.RemoveFromStr(args , 0)
			mydb.Insert(id ,strings.Join(args, " "))
		    fmt.Printf("Log %v has been inserted\n" , strings.Join(args, " "))
		},
	}	
}
