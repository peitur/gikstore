package main

import (
	"fmt"
	"giks/db"
)

func main() {
	fmt.Println("trest")

	dbFile := "test.db"
	if ! db.Exists(dbFile) {
		db.InitDb( dbFile )
	}

	fmt.Println("done ...")
}
