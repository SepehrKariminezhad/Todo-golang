package cmd

import (
	"Todo_app/dbmanegment"
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)


func CmdDelete(mydb *dbmanegment.Mydb)*cobra.Command{
return  &cobra.Command{
    	Use:   "delete [ID to delete]",
    	Short: "Delete the given ID from the DB",
    	Long: "",
    	Args: cobra.ExactArgs(1),
    	Run: func(cmd *cobra.Command, args []string) {
    	id, err := strconv.ParseUint(args[0] , 10 , 64) 
			if err != nil {
				fmt.Println("Please provide a valid ID.")
				return
			}
			mydb.Delete(id)
			fmt.Println("Delete successful")
		  
		},
	}
}