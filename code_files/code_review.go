package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Conduct a brief Code Review of the given code snippets in the context of the previous Task

/*
 * The function storyStats returns four things:
 * the shortest word
 * the longest word
 * the average word length
 * the list of words that are of length equal to the average word rounded up and down, or empty list if such words do not exist.
 */
func storyStats(str string) (map[string]string, error) {
	//do we really need this if we are only returning error in a if block?
	// var err error

	//variable naming convention could be better
	// m := make(map[string]string)
	// m["shortest"] = ""
	// m["longest"] = ""
	// m["average"] = ""

	statsMap := map[string]string{
		"shortest": "",
		"longest":  "",
		"average":  "",
	}

	if strings.TrimSpace(str) == "" {
		//only using error in if statement and returning
		// err = errors.New("can't work empty string")
		// return nil, err
		return nil, errors.New("can't work empty string")
	}

	// commenting because now i am not using them
	// var shortest int
	// var longest int

	temp_str := strings.Fields(str)

	shortBest := temp_str[0]
	longBest := ""

	for _, element := range temp_str {
		elementLen := len(element)

		if elementLen > len(longBest) {
			longBest = element
		}

		if elementLen < len(shortBest) {
			shortBest = element
		}
	}

	statsMap["shortest"] = shortBest
	statsMap["longest"] = longBest

	//this function is not available so commenting
	// statsMap["average"] = fmt.Sprint(averageNumber(temp_str))

	return statsMap, nil
}

/*
 * The function testValidity returns `true` if the given strings fits the specs,
 * and `false` otherwise.
 */
func testValidity(str string) bool {
	// patern := `[-]?\d[\d]*[\]?[\d{2}]*?[-]`
	// better regex?
	patern := `-?\d+\?\d{2}*-`
	re := regexp.MustCompile(patern)
	return re.MatchString(str)
}

func main() {
	fmt.Println(storyStats("Hello this is a code review"))
}
