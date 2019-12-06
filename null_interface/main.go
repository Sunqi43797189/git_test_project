package main

import "fmt"

type Bird struct {
}

type Pig struct {
}

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

func (b *Bird) Fly() {
	fmt.Println("bird fly")
}

func (b *Bird) Walk() {
	fmt.Println("Bird Walk")
}

func (p *Pig) Walk() {
	fmt.Println("pig walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}

	for _, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)
		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}
}
