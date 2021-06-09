package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sotycharlex/task-manager/cmd"
	"github.com/sotycharlex/task-manager/db"
	
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
