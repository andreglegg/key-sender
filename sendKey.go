package main

import (
	"github.com/namsral/flag"
	"github.com/go-vgo/robotgo"
	//"github.com/micmonay/keybd_event"
	"github.com/micmonay/keybd_event"
)

/*

 ______
|_   _ `.
  | | `. \_ .--. .---.
  | |  | [ `/'`\] /__\\
 _| |_.' /| |   | \__.,
|______.'[___]   '.__.'

 */

type Map map[string]int

var keycode = Map{
"a":  keybd_event.VK_A,
"s":  keybd_event.VK_S,
"d":  keybd_event.VK_D,
"f":  keybd_event.VK_F,
"g":  keybd_event.VK_G,
"h":  keybd_event.VK_H,
"j":  keybd_event.VK_J,
"k":  keybd_event.VK_K,
"l":  keybd_event.VK_L,
"z":  keybd_event.VK_Z,
"x":  keybd_event.VK_X,
"c":  keybd_event.VK_C,
"v":  keybd_event.VK_V,
"b":  keybd_event.VK_B,
"n":  keybd_event.VK_N,
"m":  keybd_event.VK_M,
"q":  keybd_event.VK_Q,
}

func sendKey (key string, modifier string){

	robotgo.KeyTap(key, modifier)
}

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
	sendKey(key, modifier)
	/*kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}


	//set keys
	kb.SetKeys(keycode[key])

	//set shif is pressed
	//kb.HasSHIFT(true)

	//launch
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
	//Ouput : AB*/
}


