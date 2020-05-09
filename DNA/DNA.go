package DNA

import (
	"strings"
)

type ID int

const NucleotidesAmount int = 17 // number of nucleotides in DNA

type DNA struct {
	ID                 ID     `json:"id"`
	NucleotidesFormula string `json:"nucleotidesFormula"`
}

type MutationChecker interface {
	CompareNucleotidesFormulas(*DNA) (int, error)
}

func (comparedDNA *DNA) CompareNucleotidesFormulas(DNA *DNA) (differencesCounter int, err error, optionalConsoleView string) { // compare two DNAs of different genomes with a common ancestor
	if len(DNA.NucleotidesFormula) == NucleotidesAmount && len(comparedDNA.NucleotidesFormula) == NucleotidesAmount {
		bytes1 := []byte(strings.ToUpper(DNA.NucleotidesFormula))
		bytes2 := []byte(strings.ToUpper(comparedDNA.NucleotidesFormula))
		for i := 0; i < NucleotidesAmount; i++ {
			if bytes1[i] != bytes2[i] { // checking each rune of strings
				differencesCounter++
				optionalConsoleView += "^"
			} else {
				optionalConsoleView += " "
			}
		}
		return
	} else {
		return 999999999, run(), "" // if length of strings are different
	}
}

type MyError struct {
	ErrorDescription string
}

func (e *MyError) Error() string {
	return e.ErrorDescription
}

func run() error {
	return &MyError{
		"different length of nucleotides formulas",
	}
}
