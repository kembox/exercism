package protein

import (
	"errors"
	"strings"
)

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
	var codons []string
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	var b strings.Builder
	for i := 0; i < len(rna); i++ {
		b.WriteString(string(rna[i]))
		//fmt.Println(b.String())
		if i%3 == 2 {
			codons = append(codons, b.String())
			b.Reset()
		}
	}

	var sret []string
	for _, codon := range codons {
		s, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				break
			} else {
				return nil, err
			}
		} else {
			sret = append(sret, s)
		}
	}
	return sret, nil
}

var p_map map[string][]string = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
	"STOP":          {"UAA", "UAG", "UGA"},
}

func contains(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
func FromCodon(codon string) (string, error) {
	for k, mv := range p_map {
		if contains(mv, codon) {
			if k == "STOP" {
				return "", ErrStop
			} else {
				return k, nil
			}
		}
	}
	return "", ErrInvalidBase
}
