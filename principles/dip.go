package principles

// Dependency Inversion Principle

type Relationship int

// The section bellow violates the Dependency Inversion Principle
const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
}

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

// Low level module (storage) (LLM)
type Relationships struct {
	Relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.Relations = append(r.Relations, Info{parent, Parent, child})
	r.Relations = append(r.Relations, Info{child, Child, parent})
}

// High level module (policy) (HLM)
type Research struct {
	Rels Relationships // This violates the DIP because the HLM depends on a LLM
}

func (r *Research) Investigate(fromName string, link Relationship) bool {
	relations := r.Rels.Relations
	for _, rel := range relations {
		if rel.From.Name == fromName &&
			rel.Relationship == link {
			return true
		}
	}
	return false
}

// The section bellow adheres to the Dependency Inversion Principle
type RelationshipBrowser interface {
	FindAllChildrenOf(personName string) []*Person
}

type RelationshipsDIP struct {
	Relations []Info
}

func (r *RelationshipsDIP) FindAllChildrenOf(parentName string) []*Person {
	result := []*Person{}
	for i, v := range r.Relations {
		if v.Relationship == Parent && v.From.Name == parentName {
			result = append(result, r.Relations[i].To)
		}
	}
	return result
}

func (r *RelationshipsDIP) AddParentAndChild(parent, child *Person) {
	r.Relations = append(r.Relations, Info{parent, Parent, child})
	r.Relations = append(r.Relations, Info{child, Child, parent})
}

type ResearchDIP struct {
	Browser RelationshipBrowser
}

func (r *ResearchDIP) AllChildren(parentName string) []*Person {
	return r.Browser.FindAllChildrenOf(parentName)
}
