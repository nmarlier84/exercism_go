package rnatranscription

var Trans = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

func ToRNA(dna string) string {
	res := []rune{}
	for _, myRune := range dna {
		res = append(res, Trans[myRune])
	}
	return string(res)
}
