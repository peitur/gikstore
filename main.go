package main

import (
	"os"
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

	pwd := giks.FileChecksum("build/giks", "sha512")
	fmt.Printf("Passwd: %s\n", pwd)

	c, e := giks.GeneratePrivateKeyRSA(4096)
	if e != nil{
		fmt.Println(e)
		os.Exit(1)
	}

	p1, e := giks.KeyPrivatePEM(c, pwd)
	if e != nil {
		fmt.Println( e )
		os.Exit(2)
	}

	p2, e := giks.KeyPublicPEM(c, pwd)
	if e != nil {
		fmt.Println( e )
		os.Exit(2)
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(e)

	fmt.Println("done ...")
}
