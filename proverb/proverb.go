package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var proverb []string
	template := "For want of a %s the %s was lost."
	if len(rhyme) == 0 {
		return proverb
	}
	for idx := 0; idx < len(rhyme)-1; idx++ {
		proverb = append(proverb, fmt.Sprintf(template, rhyme[idx], rhyme[idx+1]))
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverb
}
