package main

import (
	"fmt"
	"os"
	"os/user"
	"interpreter/repl"
)

func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the NOURUN programming language!\n", user.Username)
	fmt.Printf("Your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}