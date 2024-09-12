package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 100; i++ {
		repeated = repeated + character
	}
	return repeated
}

