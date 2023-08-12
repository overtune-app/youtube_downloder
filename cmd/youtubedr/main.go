package main

import (
	"fmt"
	"os"
)

func main() {
	//rootCmd.Execute()
	exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
