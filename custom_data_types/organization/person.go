package organization

import (
	"fmt"
	"strings"
)

// TwitterHandler is an alias for string
// copies the fields and the methods sets over
// type TwitterHandler = string

// TwitterHandler is a type-definition
// copies the fileds of object over
type TwitterHandler string

// RedirectURL returns Twitter handler's homesite
func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

// Identifiable is an interface
type Identifiable interface {
	ID() string
}

type Name struct {
	given  string
	family string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.given, n.family)
}

// Employee struct
type Employee struct {
	Name // embeded type
}

// Person is a struct
type Person struct {
	Name           // embeded type
	twitterHandler TwitterHandler
}

// NewPerson is
func NewPerson(givenName, familyName string) Person {
	return Person{
		Name: Name{
			given:  givenName,
			family: familyName,
		},
	}
}

// ID is a function of Person
func (p Person) ID() string {
	return "12345"
}

// SetTwitterHandler is
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) != 0 && !strings.HasPrefix(string(handler), "@") {
		return fmt.Errorf("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

// TwitterHandler is
func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
