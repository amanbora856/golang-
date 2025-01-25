package main

import "fmt"

// Separate interfaces
type Printer interface {
	Print(doc string)
}

type Scanner interface {
	Scan(doc string)
}

type Faxer interface {
	Fax(doc string)
}

// Basic printer only implements what it needs
type BasicPrinter struct{}

func (bp BasicPrinter) Print(doc string) {
	fmt.Println("Printing:", doc)
}

// Advanced printer supports all features
type AdvancedPrinter struct{}

func (ap AdvancedPrinter) Print(doc string) {
	fmt.Println("Printing:", doc)
}

func (ap AdvancedPrinter) Scan(doc string) {
	fmt.Println("Scanning:", doc)
}

func (ap AdvancedPrinter) Fax(doc string) {
	fmt.Println("Faxing:", doc)
}

func main() {
	basicPrinter := BasicPrinter{}
	basicPrinter.Print("Document 1")

	advancedPrinter := AdvancedPrinter{}
	advancedPrinter.Print("Document 2")
	advancedPrinter.Scan("Document 2")
	advancedPrinter.Fax("Document 2")
}
