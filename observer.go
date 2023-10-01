package main

import (
	"fmt"
)

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

type Observer interface {
	Update(data string)
}

type ConcreteSubject struct {
	observers []Observer
	data      string
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Unregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.data)
	}
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(data string) {
	fmt.Printf("Observer %s received an update: %s\n", o.name, data)
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.data = "New data"
	subject.Notify()

	subject.Unregister(observer1)

	subject.data = "Updated data"
	subject.Notify()
}
