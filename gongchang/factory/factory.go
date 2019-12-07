package factory

import "fmt"

type Class interface {
	Do()
}

var (
	factoryByName = make(map[string]func() Class)
)

func Register(name string, factory func() Class) {
	fmt.Println("register", name)
	factoryByName[name] = factory
}

func Create(name string) Class {
	fmt.Println("create", factoryByName)
	if factory, ok := factoryByName[name]; ok {
		return factory()
	}
	panic("name not found")
}
