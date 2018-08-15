package scale

import (
	"strings"
)

func makeMap(notes []string) map[string]int {
	m := make(map[string]int)
	for i, n := range notes {
		m[n] = i
	}
	return m
}

var sharpNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var sharpNoteMap = makeMap(sharpNotes)

var flatNotes = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var flatNoteMap = makeMap(flatNotes)

func makeBoolMap(notes []string) map[string]bool {
	m := make(map[string]bool)
	for _, n := range notes {
		m[n] = true
	}
	return m
}

var useFlatNotes = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
var useFlatNoteMap = makeBoolMap(useFlatNotes)

// Scale it return scale
func Scale(tonic, interval string) []string {

	var scale []string

	if useFlatNoteMap[tonic] {
		scale = makeNoteSlice(flatNoteMap, flatNotes, strings.Title(tonic))
	} else {
		scale = makeNoteSlice(sharpNoteMap, sharpNotes, strings.Title(tonic))
	}

	if interval != "" {
		intervalScale := make([]string, len(interval))
		count := 0
		for i, v := range interval {
			intervalScale[i] = scale[count]
			if v == 'M' {
				count += 2
			} else if v == 'A' {
				count += 3
			} else {
				count++
			}
		}
		return intervalScale
	}

	return scale
}

func makeNoteSlice(noteMap map[string]int, notes []string, tonic string) []string {
	i := noteMap[tonic]
	if i == 0 {
		return notes
	}
	return append(notes[i:len(notes)], notes[:i]...)
}
