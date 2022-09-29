package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionalPrinter struct {
}

func (m MultiFunctionalPrinter) Print(d Document) {}

func (m MultiFunctionalPrinter) Fax(d Document) {}

func (m MultiFunctionalPrinter) Scan(d Document) {}

type OldFashionedPrinter struct {
}

func (m OldFashionedPrinter) Print(d Document) {}

// Depreceted : ...
func (m OldFashionedPrinter) Fax(d Document) {}

// Deprecated :....
func (m OldFashionedPrinter) Scan(d Document) {}

func main() {

}
