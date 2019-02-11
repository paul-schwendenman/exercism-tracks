package strand

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts dna strand to rna
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))

	for pos, nucleotide := range dna {
		rna[pos] = dnaToRna[nucleotide]
	}
	return string(rna)
}
