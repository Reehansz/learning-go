package stringutil

func RevString(str string) string {
	var temp string
	for i := len(str) - 1; i >= 0; i-- {
		temp += string(str[i])
	}
	return temp
}
