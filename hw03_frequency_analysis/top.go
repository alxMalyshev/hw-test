package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Counter struct {
	FrequencyAnalysis map[string]int
}

func NewCounter() *Counter {
	var c Counter
	c.FrequencyAnalysis = make(map[string]int)
	return &c
}

func (c *Counter) Count(key string) {
	switch _, ok := c.FrequencyAnalysis[key]; ok {
	case true:
		c.FrequencyAnalysis[key]++
	default:
		c.FrequencyAnalysis[key] = 1
	}
}

func (c *Counter) Sort() []string {
	capacity := 10
	sliceOfWords := make([]string, 0)
	result := make([]string, 0)

	for word := range c.FrequencyAnalysis {
		sliceOfWords = append(sliceOfWords, word)
	}

	sort.Slice(sliceOfWords, func(i, j int) bool {
		return c.FrequencyAnalysis[sliceOfWords[i]] > c.FrequencyAnalysis[sliceOfWords[j]]
	})

	if len(sliceOfWords) < capacity {
		capacity = len(sliceOfWords)
	}
	// TODO It is ugly, but it works =)
	for i := 0; i <= capacity-1; {
		tempSlice := []string{}
		freq := c.FrequencyAnalysis[sliceOfWords[i]]
		for _, val := range sliceOfWords[i:capacity] {
			if c.FrequencyAnalysis[val] == freq {
				tempSlice = append(tempSlice, val)
			}
		}
		sort.Strings(tempSlice)
		result = append(result, tempSlice...)
		i += len(tempSlice)
	}
	return result
}

func Top10(s string) []string {
	counter := NewCounter()
	if s != "" {
		for _, word := range strings.Fields(s) {
			counter.Count(word)
		}
		return counter.Sort()
	}
	return nil
}