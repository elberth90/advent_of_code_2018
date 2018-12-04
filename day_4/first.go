package day_4

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

func First(input []string) (string, error) {
	timetable := map[time.Time]string{}
	var keys []time.Time

	for _, l := range input {
		spl := strings.Split(l, "]")
		k, err := time.Parse("2006-01-02 15:04", strings.TrimLeft(spl[0], "["))
		if err != nil {
			return "", err
		}
		timetable[k] = spl[1]
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].Before(keys[j])})

	guardTime := map[string][]uint{}
	guardID := "-1"
	var start time.Time
	var end time.Time
	for _, t := range keys {
		l := timetable[t]
		if strings.Contains(l, "#") {
			spl := strings.Split(l, " ")
			guardID = spl[2][1:]
			continue
		}

		if strings.Contains(l, "falls asleep") {
			start = t
			continue
		}

		if strings.Contains(l, "wakes up") {
			end = t
			_, s, _ := start.Clock()
			_, e, _ := end.Clock()

			if _, ok := guardTime[guardID]; !ok {
				guardTime[guardID] = make([]uint, 60, 60)
			}
			for i:=s;i<e;i++ {
				guardTime[guardID][i]++
			}
			continue
		}
	}

	most := uint(0)
	guardID = "-1"
	for key, val := range guardTime {
		c := uint(0)
		for _, v := range val {
			c+=v
		}
		if c > most {
			most = c
			guardID = key
		}
	}

	most = uint(0)
	key := 0
	for k, v := range guardTime[guardID] {
		if v > most {
			most = v
			key = k
		}
	}
	ID, err := strconv.Atoi(guardID)
	if err != nil {
		return "", err
	}
	strconv.Itoa(ID * key)
	return "", nil
}
