package pehape

import (
	"math/rand"
	"strings"
	"time"
)

var randStrShuffle *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StrShuffle Randomly shuffles a string
func StrShuffle(str string) string {
	if len(str) <= 1 {
		return str
	}

	var result strings.Builder

	mapHasVisited := make(map[int]struct{})

	length := len(str)

	for {
		if result.Len() == len(str) {
			return result.String()
		}

		index := randStrShuffle.Intn(length)
		if _, ok := mapHasVisited[index]; ok {
			continue
		}
		mapHasVisited[index] = struct{}{}

		result.WriteString(string(str[index]))
	}
}
