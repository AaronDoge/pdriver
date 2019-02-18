package main

import "fmt"

type asyncProducer struct {
	input 	chan *producerMsg

}

type producerMsg struct {

}

func (p *asyncProducer) Input() chan <- *producerMsg {
	return p.input
}

func main() {
	pro := &asyncProducer{
	}
	fmt.Printf("%p\n", pro)

	//fmt.Printf("%p\n", pro.Input())
	//fmt.Printf("%p\n", pro.Input())

	a := pro.Input()
	a <- &producerMsg{}

	//b := pro.Input()

	fmt.Printf("%p\n", a)
	//fmt.Printf("%p\n", b)

	fmt.Println(a)
	//fmt.Println(b)
}
