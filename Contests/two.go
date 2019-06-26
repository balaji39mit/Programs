package main

import (
	"sort"
)

type customMap struct {
	Key   string
	Value int32
}

func howManyAgentsToAdd(noOfCurrentAgents int32, callsTimes [][]int32) int32 {
	if len(callsTimes) <= 0 {
		return 0
	}

	timeRange := make(map[string]int32, len(callsTimes)*2)

	for _, row := range callsTimes {
		timeRange["start"] = row[0]
		timeRange["end"] = row[1]
	}
	var sortMap []customMap

	for k, v := range timeRange {
		sortMap = append(sortMap, customMap{Key: k, Value: v})
	}

	sort.Slice(sortMap, func(i, j int) bool {
		return sortMap[i].Value < sortMap[j].Value
	})

	overlapCount := int32(0)
	count := int32(0)
	for _, v := range sortMap {
		if v.Key == "start" {
			count++
		}
		if v.Key == "end" {
			count--
		}
		if count > overlapCount {
			overlapCount = count
		}
	}

	if noOfCurrentAgents > overlapCount {
		return 0
	}
	return overlapCount - noOfCurrentAgents
}

func __main() {
	input := [][]int32{
		{0, 20},
		{30, 35},
		{0, 40},
	}
	howManyAgentsToAdd(1, input)
}
