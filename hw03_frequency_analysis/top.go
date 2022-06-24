package hw03frequencyanalysis

import (

	// "regexp"
	"sort"
	"strings"
)

// var whiteSpace = regexp.MustCompile(`\s`)

type Counter struct {
	FrequencyAnalysis map[string]int
}

func NewCounter() *Counter {
	var c Counter
	c.FrequencyAnalysis = make(map[string]int)
	return &c
}

func (c *Counter) Count(key string) {
	if _, ok := c.FrequencyAnalysis[key]; ok {
		c.FrequencyAnalysis[key]++
	} else {
		c.FrequencyAnalysis[key] = 1
	}
}

func (c *Counter) Sort() []string {
	sliceOfWords := make([]string, 0)

	for word := range c.FrequencyAnalysis {
		sliceOfWords = append(sliceOfWords, word)
	}

	sort.Strings(sliceOfWords)

	sort.Slice(sliceOfWords, func(i, j int) bool {
		return c.FrequencyAnalysis[sliceOfWords[i]] > c.FrequencyAnalysis[sliceOfWords[j]]
	})

	if len(sliceOfWords) < 10 {
		return sliceOfWords
	} else {
		return sliceOfWords[:10]
	}
}

func Top10(s string) []string {
	counter := NewCounter()
	// slice := whiteSpace.Split(s, -1)
	slice := strings.Fields(s)
	for _, word := range slice {
		counter.Count(word)
	}
	return counter.Sort()
}
