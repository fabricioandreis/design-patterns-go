package creational

import (
	"fmt"
	"strings"
)

// Builder is a creational design pattern that lets you construct complex objects step by step.
// The pattern allows you to produce different types and representations of an object using the same construction code.
// https://refactoring.guru/design-patterns/builder

func CreateParagraph(message string) string {
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(message)
	sb.WriteString("</p>")

	return sb.String()
}

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{rootName, "", []HtmlElement{}}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	b.AddChild(childName, childText)
	return b
}

// Builder Facets
// Create more than one builder to different aspects of an object
type Person struct {
	// Address
	StreetAddress, PostCode, City string
	// Job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostCode(postCode string) *PersonAddressBuilder {
	b.person.PostCode = postCode
	return b
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

type PersonJobBuilder struct {
	PersonBuilder
}

// Builder Parameter
// Using Builder to make sure that and object is created correctly
// Turn email struct as private and the Builder public
type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("from email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("to email should contain @")
	}
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	fmt.Printf("Sending email %v\n", email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

// Functional Builder
// This approach is useful to extend the builder with build actions
// instead of creating new builders that aggregate the current builder
type PersonFunc struct {
	Name, Position string
}

type personModifier func(*PersonFunc)

type PersonBuilderFunc struct {
	actions []personModifier
}

func (b *PersonBuilderFunc) Called(name string) *PersonBuilderFunc {
	b.actions = append(b.actions, func(p *PersonFunc) {
		p.Name = name
	})
	return b
}

func (b *PersonBuilderFunc) Works(position string) *PersonBuilderFunc {
	b.actions = append(b.actions, func(p *PersonFunc) {
		p.Position = position
	})
	return b
}

func (b *PersonBuilderFunc) Build() *PersonFunc {
	p := PersonFunc{}
	for _, action := range b.actions {
		action(&p)
	}
	return &p
}
