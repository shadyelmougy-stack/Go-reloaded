package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* ---------- main ---------- */

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <inputfileName>.txt <outputFileName>.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	lines := strings.Split(string(data), "\n")
	var result []string

	for _, line := range lines {
		result = append(result, processLine(line))
	}

	final := strings.Join(result, "\n")

	err = os.WriteFile(outputFile, []byte(final), 0644)
	if err != nil {
		fmt.Println("Error writing file")
	}
}

/* ---------- helpers ---------- */

func hexToDec(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(n))
}

func binToDec(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(n))
}

func capitalize(w string) string {
	if len(w) == 0 {
		return w
	}
	return strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
}

func isVowelOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	r := strings.ToLower(string(word[0]))
	return strings.Contains("aeiouh", r)
}

/* ---------- commands ---------- */

func applySimple(result []string, cmd string) {
	if len(result) == 0 {
		return
	}

	i := len(result) - 1

	switch cmd {
	case "(hex)":
		result[i] = hexToDec(result[i])
	case "(bin)":
		result[i] = binToDec(result[i])
	case "(up)":
		result[i] = strings.ToUpper(result[i])
	case "(low)":
		result[i] = strings.ToLower(result[i])
	case "(cap)":
		result[i] = capitalize(result[i])
	}
}

func applyMulti(result []string, cmd string) {
	cmd = strings.Trim(cmd, "()")
	parts := strings.Split(cmd, ",")

	if len(parts) != 2 {
		return
	}

	action := strings.TrimSpace(parts[0])
	n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return
	}

	start := len(result) - n
	if start < 0 {
		start = 0
	}

	for i := start; i < len(result); i++ {
		switch action {
		case "up":
			result[i] = strings.ToUpper(result[i])
		case "low":
			result[i] = strings.ToLower(result[i])
		case "cap":
			result[i] = capitalize(result[i])
		}
	}
}

/* ---------- punctuation ---------- */

func fixPunctuation(text string) string {

	// remove before
	replacer1 := strings.NewReplacer(
		" ,", ",",
		" .", ".",
		" !", "!",
		" ?", "?",
		" :", ":",
		" ;", ";",
	)
	text = replacer1.Replace(text)

	// add after
	replacer2 := strings.NewReplacer(
		",", ", ",
		".", ". ",
		"!", "! ",
		"?", "? ",
	)
	text = replacer2.Replace(text)

	text = strings.Join(strings.Fields(text), " ")

	replacer3 := strings.NewReplacer(
		". . .", "...",
		"...", "... ",
	)
	text = replacer3.Replace(text)

	return text
}

func fixQuotes(words []string) []string {
	var result []string
	insideQuote := false

	for i := 0; i < len(words); i++ {

		if words[i] == "'" {

			if !insideQuote {
				insideQuote = true
				if i+1 < len(words) {
					words[i+1] = "'" + words[i+1]
				}

			} else {
				insideQuote = false
				if len(result) > 0 {
					result[len(result)-1] = result[len(result)-1] + "'"
				}
			}

		} else {
			result = append(result, words[i])
		}
	}

	return result
}

/* ---------- process ---------- */

func processLine(line string) string {

	words := strings.Fields(line)
	var result []string

	// a → an
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" && isVowelOrH(words[i+1]) {
			words[i] = "an"
		}
	}

	for i := 0; i < len(words); i++ {

		w := words[i]

		//  fix problem (cap, 6)
		if strings.HasPrefix(w, "(") && strings.Contains(w, ",") && !strings.HasSuffix(w, ")") {
			if i+1 < len(words) {
				w = w + " " + words[i+1]
				i++
			}
		}

		// command
		if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")") {

			if strings.Contains(w, ",") {
				applyMulti(result, w)
			} else {
				applySimple(result, w)
			}
			continue
		}

		result = append(result, w)
	}

	result = fixQuotes(result)

	output := strings.Join(result, " ")
	output = fixPunctuation(output)

	return output
}
