package main

import (
	"fmt"
	"msp_lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	const MSP_LANG = `            
	__  __   ____    ____      _          _      _   _    ____ 
	|  \/  | / ___|  |  _ \    | |        / \    | \ | |  / ___|
	| |\/| | \___ \  | |_) |   | |       / _ \   |  \| | | |  _ 
	| |  | |  ___) | |  __/    | |___   / ___ \  | |\  | | |_| |
	|_|  |_| |____/  |_|       |_____| /_/   \_\ |_| \_|  \____|
																
`
	fmt.Printf(MSP_LANG)

	fmt.Printf("Hello %s! This is the MSP Lang programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
