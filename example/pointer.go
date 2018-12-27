package main

import "fmt"

type Abc struct {
	Shell 	string
	Fish 	string

	Sub 	*Sub
}

type Sub struct {
	name 	string
}

func (s Sub) Input() {

}

func main() {
	sub := &Sub{
		name: 	"sub name",
	}

	ta := &Abc{
		Shell: 	"shellcontent",
		Fish: 	"fishcontent",
	}

	ta.Sub = sub

	(*ta.Sub).Input()
	ta.Sub.Input()

	fmt.Println()
}
