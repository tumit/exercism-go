package strand

// ToRNA it return RNA of DNA
func ToRNA(dna string) string {

	rna := ""

	for _, s := range dna {
		rna += toRNA(string(s))
	}

	return rna

}

func toRNA(dna string) string {

	if dna == "G" {
		return "C"
	}

	if dna == "C" {
		return "G"
	}

	if dna == "T" {
		return "A"
	}

	if dna == "A" {
		return "U"
	}

	return ""
}
