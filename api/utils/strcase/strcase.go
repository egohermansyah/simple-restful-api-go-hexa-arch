package strcase

func IsUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

func ToLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func CamelCaseToSnakeCase(s string) string {
	result := make([]rune, 0)
	delimiter := []rune("_")
	for index, char := range s {
		if IsUpper(char) {
			if index > 0 {
				result = append(result, delimiter[0])
			}
		}
		result = append(result, ToLower(char))
	}
	return string(result)
}
