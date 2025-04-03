package piscine

func SplitWhiteSpaces(s string) []string {
	i := 0
	p := ""
	len := 0

	for i, value := range s {
		if (value == ' ' || value == '\n' || value == '\t') && (i != 0 && (s[i-1] != ' ' && s[i-1] != '\n' && s[i-1] != '\t')) {
			len++
		}
	}

	array := make([]string, len+1)
	for _, value := range s {
		if value == ' ' || value == '\n' || value == '\t' {
			if p != "" {
				array[i] = p
				p = ""
				i++
			}
		} else {
			p = p + string(value)
		}
	}

	if p != "" {
		array[i] = p
	}
	return array
}
