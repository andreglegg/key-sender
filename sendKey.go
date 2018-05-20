package main

import (
	"os"
	"github.com/go-vgo/robotgo"
)

/*

 ______
|_   _ `.
  | | `. \_ .--. .---.
  | |  | [ `/'`\] /__\\
 _| |_.' /| |   | \__.,
|______.'[___]   '.__.'

 */

func main() {
	//argLength := len(os.Args[2])-1;

	//allArgs := os.Args[1:]
	key := os.Args[1]
	action := os.Args[2]
	modifier := os.Args[3]

	/*fmt.Println(argLength)
	fmt.Println(allArgs)
	fmt.Println(key)
	fmt.Println(action)
	fmt.Println(modifier)*/
	robotgo.KeyToggle(key, action, modifier)
}
