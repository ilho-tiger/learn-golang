package organization

import (
	"fmt"
	"strings"
)

// Identifiable is an interface
type Identifiable interface {
	ID() string
}

// Person is a struct
type Person struct {
	givenName      string
	familyName     string
	twitterHandler string
}

func NewPerson(givenName, familyName string) Person {
	return Person{
		givenName:  givenName,
		familyName: familyName,
	}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.givenName, p.familyName)
}

// ID is a function of Person
func (p Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) != 0 && !strings.HasPrefix(handler, "@") {
		return fmt.Errorf("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p Person) TwitterHandler() string {
	return p.twitterHandler
}
