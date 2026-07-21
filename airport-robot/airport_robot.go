package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

// As a first step, define an interface Greeter with two methods.
// LanguageName which returns the name of the language (a string) that the robot is supposed to greet the visitor in.
// Greet which accepts a visitor's name (a string) and returns a string with the greeting message in a specific language.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type German struct{}

func (t German) LanguageName() string {
	return "German"
}

func (t German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

type Italian struct{}

func (t Italian) LanguageName() string {
	return "Italian"
}

func (t Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (t Portuguese) LanguageName() string {
	return "Portuguese"
}

func (t Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
