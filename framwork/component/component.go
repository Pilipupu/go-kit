package component

import (
	"fmt"
	"github.com/Pilipupu/go-kit/datastruct/maps/hashmap"
)

// Component /**
type Component interface {
	Start() (func() error, error)
	Stop()
	Name() string
}

var components = hashmap.New[string, Component]()
var cleanFunctions = hashmap.New[string, func() error]()

func Register(component Component) {
	components.Put(component.Name(), component)
}

func Start() {
	for _, component := range components.Values() {
		cleanFunc, err := component.Start()
		if err != nil {
			panic(fmt.Errorf("failed to start component[%s]: %w", component.Name(), err))
		}
		cleanFunctions.Put(component.Name(), cleanFunc)
	}
}

func Stop() {
	for _, f := range cleanFunctions.Values() {
		f()
	}
}
