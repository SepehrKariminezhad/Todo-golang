package main

import (
	"Todo_app/cmd"
	"Todo_app/dbmanegment"
)



func main(){
	var dbrepo = new(dbmanegment.Mydb)
	dbrepo.Connect()
	cmd.RootCmd.AddCommand(cmd.CmdInsert(dbrepo) , cmd.CmdUpdate(dbrepo) , cmd.CmdRead(dbrepo) , cmd.CmdReadall(dbrepo) , cmd.CmdDelete(dbrepo))
	cmd.RootCmd.Execute()
}
