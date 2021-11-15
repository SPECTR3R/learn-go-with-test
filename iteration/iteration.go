package iteration

func Repeat(character string, val int) string {
	var repeated string

	for i := 0; i < val; i++ {
		repeated += character
	}
	return repeated
}
