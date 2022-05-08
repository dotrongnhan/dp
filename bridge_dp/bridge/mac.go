package bridge

import "fmt"

type Mac struct {
	Printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.Printer = printer
}
