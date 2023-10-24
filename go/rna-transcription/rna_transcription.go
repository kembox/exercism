package strand

import "strings"

func ToRNA(dna string) string {
	var sb strings.Builder
	for _, c := range dna {
		switch c {
		case 'G':
			sb.WriteString("C")
		case 'C':
			sb.WriteString("G")
		case 'T':
			sb.WriteString("A")
		case 'A':
			sb.WriteString("U")
		default:
			sb.WriteString(string(c))
		}
	}
	return sb.String()
}
