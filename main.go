package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	var (
		path string
		err  error
	)

	if runtime.GOOS == "windows" {
		path, err = os.Getwd()

		if err != nil {
			fmt.Printf("Tidy : %s\n", err)
			os.Exit(1)
		}

	} else {

		flag.StringVar(&path, "path", "/home", "Path where program will be executed.")

		flag.Parse()
	}

	order := NewOrder(path)

	err = order.tidyFolder()

	if err != nil {
		fmt.Printf("Tidy : %s\n", err)
		os.Exit(1)
	}

	time.Sleep(time.Second * 5)

}
