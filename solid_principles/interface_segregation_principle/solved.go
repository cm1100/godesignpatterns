package main

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

type PhotoCopier struct {
}

func (p PhotoCopier) Scan(d Document) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoCopier) Print(d Document) {
	//TODO implement me
	panic("implement me")
}

type MultiPrinter interface {
	Scanner
	Printer
}
type MultiMachine struct {
	printter Printer
	Scanner  Scanner
}
