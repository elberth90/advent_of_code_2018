package day_3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type fabric struct {
	id   int
	left int
	top  int
	wide int
	tall int

	leftTop point
	rightDown point
}

func First(input []string) (string, error) {
	var m []fabric
	for _, l := range input {
		f, err := buildFabric(l)
		if err != nil {
			return "", err
		}
		m = append(m, f)
	}
	fmt.Printf("%+v \n", m)

	var counter int
	var overlap float64
	for i:=0; i<len(m); i++ {
		for j:=0; j<len(m); j++ {
			if i == j {
				continue
			}
			if doOverlap(m[i], m[j]) {
				overlap += getOverlapArea(m[i], m[j])
				counter++
			}
		}
	}

	fmt.Println(overlap)
	return strconv.Itoa(counter), nil
}

func doOverlap(f1, f2 fabric) bool {
	if f1.leftTop.x > f2.rightDown.x || f2.leftTop.x > f1.rightDown.x {
		return false
	}

	if f1.leftTop.y < f2.rightDown.y || f2.leftTop.y < f1.rightDown.y {
		return false
	}

	return true
}

func getOverlapArea(f1, f2 fabric) float64 {
	return (math.Min(float64(f1.rightDown.x), float64(f2.rightDown.x)) - math.Max( float64(f1.leftTop.x),  float64(f2.leftTop.x))) * (math.Min( float64(f1.rightDown.y),  float64(f2.rightDown.y)) -	math.Max( float64(f1.leftTop.y),  float64(f2.leftTop.y)))
}

func buildFabric(line string) (fabric, error) {
	a := strings.Split(line, " ")

	distance := strings.Split(strings.TrimRight(a[2], ":"), ",")
	size := strings.Split(a[3], "x")

	id, err := strconv.Atoi(strings.TrimLeft(a[0], "#"))
	if err != nil {
		return fabric{}, err
	}

	left, err := strconv.Atoi(distance[0])
	if err != nil {
		return fabric{}, err
	}
	top, err := strconv.Atoi(distance[1])
	if err != nil {
		return fabric{}, err
	}
	wide, err := strconv.Atoi(size[0])
	if err != nil {
		return fabric{}, err
	}
	tall, err := strconv.Atoi(size[1])
	if err != nil {
		return fabric{}, err
	}

	f := fabric{
		id:   id,
		left: left,
		top:  top,
		wide: wide,
		tall: tall,

		leftTop: point{left, top},
		rightDown: point{left+wide, top-tall},
	}
	return f, nil
}
