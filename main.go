package main

import (
	"fmt"
	"giks/db"
	"giks/store"
)

func main() {
	fmt.Println("trest")

	dbFile := "test.db"
	if !db.Exists(dbFile) {
		db.InitDb(dbFile)
	}

	c, e := store.GeratePrivateKeyRSA(64, "test")
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println("done ...")
}
