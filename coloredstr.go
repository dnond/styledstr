package styledstr

func Red(str string) string {
	return convert(str, "31")
}

func Green(str string) string {
	return convert(str, "32")
}

func Yellow(str string) string {
	return convert(str, "33")
}

func Blue(str string) string {
	return convert(str, "34")
}

func Mazenda(str string) string {
	return convert(str, "35")
}

func Cian(str string) string {
	return convert(str, "36")
}

func White(str string) string {
	return convert(str, "37")
}

func Colors() []string {
	return []string{"Red", "Green", "Yellow", "Blue", "Mazenda", "Cian", "White"}
}

/////

func Underbar(str string) string {
	return convert(str, "4")
}

func Bold(str string) string {
	return convert(str, "4")
}

////

func convert(str, colorCode string) string {
	return "\x1b[" + colorCode + "m" + str + "\x1b[0m"
}
