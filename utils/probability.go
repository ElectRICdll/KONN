package utils

import (
	"math/rand"
)

func GodsJudge(percent int) bool {
	return percent > rand.Int()%100
}
