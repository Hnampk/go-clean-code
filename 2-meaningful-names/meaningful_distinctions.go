package chapter2

// copyCharsBadNaming use number-series naming
func copyCharsBadNaming(a1 []string, a2 []string) {
	for i := 0; i < len(a1); i++ {
		a2[i] = a1[i]
	}
}

func copyCharsBetterNaming(source []string, destination []string) {
	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}
}
