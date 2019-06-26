package main

import (
	"sort"
	"strings"
)

type sortMap struct {
	Key   int32
	Value int32
}

// Complete the sort_hotels function below.
func sort_hotels(keywords string, hotel_ids []int32, reviews []string) []int32 {
	hotels := make(map[int32]int32, 0)
	for _, val := range hotel_ids {
		if _, ok := hotels[val]; !ok {
			hotels[val] = 0
		}
	}
	keyWords := strings.Split(keywords, " ")
	for i, val := range reviews {
		count := int32(0)
		for _, v := range keyWords {
			count += int32(strings.Count(val, v))
		}
		hotels[hotel_ids[i]] = count
	}
	var res []sortMap

	for k, v := range hotels {
		res = append(res, sortMap{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Value == res[j].Value {
			return res[i].Key < res[j].Key
		}
		return res[i].Value > res[j].Value
	})
	output := make([]int32, len(res))
	for k, v := range res {
		output[k] = v.Key
	}
	return output
}
