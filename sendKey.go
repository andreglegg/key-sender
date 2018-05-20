package main

import (
	"github.com/namsral/flag"
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
	var key string
	//var action string
	var modifier string
	flag.StringVar(&key, "key", "", "key to press")
	//flag.StringVar(&action, "action", "", "press down or up?")
	flag.StringVar(&modifier, "modifier", "null", "shift, alt or ctrl")
	flag.Parse()
	/*fmt.Print("key:", key)
	fmt.Println("modifier: ", modifier)*/
	/*fmt.Println(argLength)
	fmt.Println(allArgs)
	fmt.Println(key)
	fmt.Println(action)
	fmt.Println(modifier)*/
	//robotgo.KeyToggle(key, action, modifier)
	//robotgo.KeyToggle(key, action, modifier)
	if modifier == "null" {
		robotgo.KeyTap(key)
	} else {
		robotgo.KeyTap(key, modifier)
	}
}
