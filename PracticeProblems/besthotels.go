package practice_problems

import (
	"fmt"
	"sort"
)

/*
PROBLEM STATEMENT : Given a hotel id and it's score, Print the hotels in the following conditions.
1. Sort the hotels in descending order of an average score.
2. If two hotels has same average score, pick the hotel with less id.
*/

type customSort struct {
	res map[int]float64
	key []int //since we want to sort by value
}

func (c *customSort) Len() int {
	return len(c.res)
}

func (c *customSort) Less(i, j int) bool {

	//edge case : there can be possibility that the key does not exist.
	//HEre is it possible? No. Coz,
	if c.res[c.key[i]] == c.res[c.key[j]] {
		return c.key[i] < c.key[j]
	}
	return c.res[c.key[i]] > c.res[c.key[j]]
}

func (c *customSort) Swap(i, j int) {
	c.key[i], c.key[j] = c.key[j], c.key[i]
}

func best_hotels() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	res := make(map[int]map[int]float64, n)
	for i := 0; i < n; i++ {
		var id int
		var score float64
		_, _ = fmt.Scanf("%d %f", &id, &score)
		if val, ok := res[id]; ok {
			//This should not run more than once.
			for key, value := range val {
				key = key + 1
				value = value + score
				res[id] = map[int]float64{
					key: value,
				}
			}
		} else {
			res[id] = map[int]float64{1: score}
		}
	}

	cs := new(customSort)
	cs.res = make(map[int]float64, len(res))
	cs.key = make([]int, 0, len(res))

	for key, val := range res {
		for k, v := range val {
			//fmt.Println("Entering once for : ", key, k, v)
			score := v / float64(k)
			cs.res[key] = score
			cs.key = append(cs.key, key)
		}
	}

	sort.Sort(cs)

	for _, key := range cs.key {
		fmt.Println(key, cs.res[key])
	}

	//Now we have the result
}
