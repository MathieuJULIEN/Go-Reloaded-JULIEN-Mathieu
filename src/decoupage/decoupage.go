package decoupage

var ListeMot []string

func Decoupage(text string) {
	ListeMot = nil
	var MotProvi string
	for a := range text {
		if rune(text[a]) == ' ' {
			ListeMot = append(ListeMot, MotProvi)
			MotProvi = ""
		} else {
			MotProvi += string(text[a])
		}
		if len(MotProvi) != 0 && a == len(text)-1 {
			ListeMot = append(ListeMot, MotProvi)
		}
	}
}
