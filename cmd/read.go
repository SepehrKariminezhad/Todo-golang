package cmd

import (
	"Todo_app/dbmanegment"
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

func CmdReadall(mydb *dbmanegment.Mydb)*cobra.Command{
return &cobra.Command{
      Use:   "readall",
      Short: "read all the data from your DB",
      Long: "",
      Run: func(cmd *cobra.Command, args []string) {
		    mydb.Read_all()
      },
  }
}

func CmdRead(mydb *dbmanegment.Mydb)*cobra.Command{
return &cobra.Command{
    Use:   "read [ID to read]",
    Short: "read the data with the given ID",
    Long: "",
	  Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0] , 10 , 64)
		if err != nil {
			fmt.Println("Please provide a valid ID.")
			return
		}
      mydb.Read(id)
    },
  }
}