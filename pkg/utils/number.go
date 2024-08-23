package utils

func DeclOfNum(n int, textForms [3]string) string {
	n = abs(n) % 100
	n1 := n % 10
	if n > 10 && n < 20 {
		return textForms[2]
	}
	if n1 > 1 && n1 < 5 {
		return textForms[1]
	}
	if n1 == 1 {
		return textForms[0]
	}
	return textForms[2]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
