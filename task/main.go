package main

import (
	"path/filepath"
	"task/cmd"
	"task/db"
	"github.com/mitchellh/go-homedir"
)

func main(){
	home, err := homedir.Dir()
	if err!=nil{
		panic(err)
	} 
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
	if err!=nil{
		panic(err)
	}
	cmd.RootCmd.Execute()
}


