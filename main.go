package main

import (
	"os"
	"fmt"
	"giks"
	"giks/db"
)

func main() {
	fmt.Printf("started: %s\n", os.Args[1:] )

	if len( os.Args ) == 1{
		fmt.Println("Need reference file")
		os.Exit(1)
	}

	dbFile := os.Args[1]
	if !db.Exists(dbFile) {
		db.InitDb(dbFile)
	}

	pwd := giks.FileChecksum("build/giks", "sha512")
	fmt.Printf("Passwd: %s\n", pwd)

	c, e := giks.GeneratePrivateKeyRSA(64)
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
