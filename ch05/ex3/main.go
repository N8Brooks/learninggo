package main

func Prefixer(prefix string) func(string) string {
	return func(suffix string) string {
		return prefix + " " + suffix
	}
}

func main() {
	pre := Prefixer("pre")
	println(pre("fix"))
}
