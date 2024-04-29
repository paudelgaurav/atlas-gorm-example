package utils

func Exists(s string, slice []string) bool {
	for _, cur := range slice {
		if cur == s {
			return true
		}
	}
	return false
}

func SubSet(s1, s2 []string) bool {
	for _, s := range s1 {
		if !Exists(s, s2) {
			return false
		}
	}
	return true
}
