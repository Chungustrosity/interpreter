let map = fn(arr, f){ let iter = fn(arr, accumulated){ if (len(arr) == 0) {return accumulated} else {return iter(rest(arr), puch(accumulated, f(first(iter))))} } return iter(arr, []); }
package main

import (
	"fmt"
	"os"
	"os/user"

	"interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
