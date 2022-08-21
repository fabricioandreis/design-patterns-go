package structural

import (
	"strings"
	"unicode"
)

// Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.
// https://refactoring.guru/design-patterns/flyweight

// Avoid redundancy when storing data

type FormattedTextHeavy struct {
	plainText  string
	capitalize []bool
}

func NewFormattedTextHeavy(planText string) *FormattedTextHeavy {
	return &FormattedTextHeavy{planText, make([]bool, len(planText))}
}

func (f *FormattedTextHeavy) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedTextHeavy) Capitalize(start, end int) {
	first := start
	if first < 0 {
		first = 0
	}
	for i := start; i <= end && i < len(f.plainText); i++ {
		f.capitalize[i] = true
	}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type FormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText: plainText}
}

func (f *FormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	f.formatting = append(f.formatting, r)
	return r
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		for _, r := range f.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

type UserHeavy struct {
	FullName string
}

func NewUserHeavy(fullName string) *UserHeavy {
	return &UserHeavy{fullName}
}

var allNames []string

type User struct {
	names []uint8
}

func getOrAdd(s string) uint8 {
	for i := range allNames {
		if allNames[i] == s {
			return uint8(i)
		}
	}
	allNames = append(allNames, s)
	return uint8(len(allNames) - 1)
}

func NewUser(fullName string) *User {
	user := User{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		user.names = append(user.names, getOrAdd(p))
	}

	return &user
}

func (u *User) FullName() string {
	var parts []string
	for _, n := range u.names {
		parts = append(parts, allNames[n])
	}
	return strings.Join(parts, " ")
}
