package hello

// Greet returns a greeting string.
func Greet(name string) string {
	if name == "" {
		name = "world"
	}
	return "hello, " + name
}
