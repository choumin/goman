package main

import "fmt"
import "runtime"
import "os"
import "strings"

func goman() {
	const prompt = ">>> "
	const exit = "exit"
	const quit = "quit"
	fmt.Println(runtime.Version())
	for true {
		fmt.Print(prompt)
		var directive string
		fmt.Scanf("%s", &directive)
		fmt.Println(directive)
		if strings.EqualFold(directive, exit) ||
			strings.EqualFold(directive, quit) {
			os.Exit(0)
		}
	}
}
func main() {
	//fmt.Println("vim-go")
	goman()
}
