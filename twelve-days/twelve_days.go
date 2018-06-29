// Package twelve ...
package twelve

import (
	"fmt"
	"strings"
)

const intro = "On the %s day of Christmas my true love gave to me"

// kv maps a lyric component to its count.
type kv struct {
	k string
	v string
}

// verses is the mapping of lyric component and its count.
var verses = []kv{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song ...
func Song() string {
	lyrics := make([]string, len(verses))
	for i := 0; i < len(verses); i++ {
		lyrics[i] = Verse(i+1) + "\n"
	}
	return strings.Join(lyrics, "")
}

// Verse ...
func Verse(line int) string {
	line--

	var lyrics []string
	lyrics = append(lyrics, fmt.Sprintf(intro, verses[line].k))

	if line == 0 {
		lyrics = append(lyrics, verses[0].v)
		return strings.Join(lyrics, ", ")
	}

	for line > 0 {
		lyrics = append(lyrics, verses[line].v)
		line--
	}

	lyrics = append(lyrics, "and "+verses[0].v)
	return strings.Join(lyrics, ", ")
}
