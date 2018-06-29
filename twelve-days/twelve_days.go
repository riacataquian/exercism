// Package twelve output the lyrics to 'The Twelve Days of Christmas'.
package twelve

import (
	"fmt"
)

const intro = "On the %s day of Christmas my true love gave to me, "

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

// Song returns the entire song.
func Song() string {
	var lyrics string
	for i := 0; i < len(verses); i++ {
		v := Verse(i+1) + "\n"
		lyrics += v
	}
	return lyrics
}

// Verse returns the song lyric based on the verse number provided.
func Verse(i int) string {
	i--
	f := verses[0].v

	lyrics := fmt.Sprintf(intro, verses[i].k)
	if i == 0 {
		return lyrics + f
	}

	for i > 0 {
		lyrics += verses[i].v + ", "
		i--
	}

	return lyrics + "and " + f
}
