package strand

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts dna strand to rna
func ToRNA(dna string) string {
	rna := ""

	for _, nucleotide := range dna {
		rna += string(dnaToRna[nucleotide])
	}
	return rna
}
