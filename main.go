package main

import (
	"fmt"
	"os"
)

func main() {

	// if err:= cmd.rootCmd.Execute();err!=nil{
	// 	fmt.Println(os.Stderr, err);
	// 	os.Exit(1);
	// }

	if err := cmd.rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
