package proteintranslation

import "errors"

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0)
	for k := 0; k < len(rna); k += 3 {
		if k+3 > len(rna) {
			return nil, ErrInvalidBase
		}
		codon := rna[k : k+3]
		aux, err := FromCodon(codon)
		if errors.Is(err, ErrStop) {
			break
		}
		if err != nil {
			return result, err
		}
		result = append(result, aux)
	}

	return result, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "STOP", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
