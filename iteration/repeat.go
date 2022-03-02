package iteration

func Repeat(character string, count int) string {
	stringBuilder := ""
	for i := 1; i <= count; i++ {
		stringBuilder += character
	}
	return stringBuilder
}
