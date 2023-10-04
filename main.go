package main

import (
	"flag"
	"fmt"
	"os"

	"todocli/app"
)

func main() {
	storagePath := getFlags()

	app, err := app.ConfigureApp(storagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.Run()
}

func getFlags() string {
	storagePath := flag.String("storage-path", "", "Specify path of storage")
	flag.Parse()

	return *storagePath
}
