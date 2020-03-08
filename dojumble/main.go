package main

import (
	"flag"
	"fmt"
	"strings"

	jumble "github.com/mglavitsch/jumbler-go"
)

const defaultText = "It doesn't matter in what order the letters in a word are, the only important thing is that the first and last letter be at the right place. The rest can be a total mess and you can still read it without problem. This is because the human mind does not read every letter by itself, but the word as a whole."

func main() {
	var text string

	flag.StringVar(&text, "text", "", "Text to be jumbled")
	flag.Parse()

	if strings.Compare(text, "") == 0 {
		text = defaultText
	}

	jumbledText, _ := jumble.Text(text)
	fmt.Printf("\n%v\n\n", jumbledText)
}
