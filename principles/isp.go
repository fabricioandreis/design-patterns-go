package principles

type Document struct {
}

// The section bellow violates the Interface Segregation Principle
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {

}

func (m *MultiFunctionPrinter) Fax(d Document) {

}

func (m *MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {
	panic("not implemented") // TODO: Implement
}

func (o *OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// This is a bad implementation since the Old Fashioned Printer
// should not need to implement an operation that it doesn't have
func (o *OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// The section bellow adheres to the Interface Segregation Principle
type Printer interface {
	Print(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
}
func (p Photocopier) Print(d Document) {
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Faxer
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
