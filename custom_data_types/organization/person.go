package organization

import (
	"fmt"
	"strconv"
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

// Citizen need comment
type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

// NewSocialSecurityNumber needs comment
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

// NewEuropeanUnionIdentifier needs comment
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch value := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      value,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(value),
			country: country,
		}
	case europeanUnionIdentifier:
		return value
	case Person:
		euID, _ := value.Citizen.(europeanUnionIdentifier)
		return euID
	default:
		panic("Using an invalid type to initialze EU identifier")
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

// Name needs comment
type Name struct {
	given  string
	family string
}

// FullName is
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
	Citizen
}

// NewPerson is
func NewPerson(givenName, familyName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			given:  givenName,
			family: familyName,
		},
		Citizen: citizen,
	}
}

// ID is a function of Person
func (p Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
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
