package randomstring

func isMember(c, s string) bool {
	output := false

	for i := range s {
		if c == string(s[i]) {
			output = true
		}
	}

	return output
}
