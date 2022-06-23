package hw03frequencyanalysis

import (
	// "fmt"
	// "sort"
	// "strings"
	"regexp"
	"sort"
)

var whiteSpace = regexp.MustCompile(`\s`)

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

func (c *Counter) Sort(slice []string) []string {
	// sort.Slice(c.FrequencyAnalysis, func(i, j int)bool{
	// 	return 
	// })
	return nil
}

func Top10(s string) []string {
	counter := NewCounter()
	slice := whiteSpace.Split(s, -1)
	for _, word := range slice {
		counter.Count(word)
	}
	return nil
}
