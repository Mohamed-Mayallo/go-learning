package main

import "fmt"

func hello(who, lang string) string {
	if who == "" {
		who = "world"
	}

	prefix := getPrefix(lang)

	return prefix + who
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case "english":
		prefix = "Hello, "
	case "spanish":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(hello("Mohamed", ""))
}
