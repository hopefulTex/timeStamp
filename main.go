package main

import (
	"fmt"
	ui "timeStampFormatter/ui"
)

func main() {
	p := ui.New()
	m, err := p.Run()
	if err != nil {
		fmt.Printf("\nerror: %s\n", err.Error())
	} else {
		fmt.Println(m.View())
	}
}
