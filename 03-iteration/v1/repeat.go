package iteration

func Repeat(character string) string {
	var repeated string

	for i := 1; i < 4; i++ {
		repeated = repeated + character
	}

	return repeated
}
