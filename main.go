package main

import (
	"fmt"
	"giks"
	"giks/db"
)

func main() {
	fmt.Println("started")

	dbFile := "test.db"
	if !db.Exists(dbFile) {
		db.InitDb(dbFile)
	}

	pwd := giks.FileChecksum("/etc/hostname", "sha512")
	fmt.Printf("Passwd: %s\n", pwd)

	c, e := giks.GeneratePrivateKeyRSA(4096)
	p, e := giks.KeyPEM(c, pwd)

	fmt.Println(p)
	fmt.Println(e)

	fmt.Println("done ...")
}
