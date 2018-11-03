package cmd

// IsMatch regexp with ".", "*" supports
func IsMatch(text, pattern string) bool {
	if text == "" {
		// '' and ''
		if pattern == "" {
			return true
		}
		// '' and 'a*'
		return len(pattern) == 2 && pattern[1] == '*'
	}

	// 'a' and ''
	if pattern == "" {
		return false
	}
	//  'a' and 'b*'
	if len(pattern) > 1 && pattern[1] == '*' {
		return IsMatch(text, pattern[2:]) || IsMatch(text[1:], pattern)
	}

	//  'a' and 'a'
	if text[0] == pattern[0] {
		return IsMatch(text[1:], pattern[1:])
	}

	//  'a' and '.'
	if pattern[0] == '.' {
		return IsMatch(text[1:], pattern[1:])
	}

	//  'a' and 'b'
	return false
}
