def to_rna(dna_strand):
    dna_map = {
        'G': 'C',
        'C': 'G',
        'A': 'U',
        'T': 'A',
    }
    return ''.join(dna_map[dna_item] for dna_item in dna_strand)
