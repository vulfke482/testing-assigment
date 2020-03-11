package rand

import (
	"testing"

	"git.sansera.com/mtkach/golang-test-assignment/utils/testutil"
)

func TestWeightIndex(t *testing.T) {
	result := map[interface{}]int{}

	for i := 0; i < 10000; i++ {
		r := WeightIndex(map[interface{}]uint64{
			"p1": 50,
			"p2": 50,
		})

		if _, e := result[r]; !e {
			result[r] = 0
		}
		result[r]++
	}
	v := float64(result["p1"]) / float64(result["p2"])

	testutil.TestingAssertEqual(t, v > 0.9, true)
	testutil.TestingAssertEqual(t, v <= 1.1, true)
}

/*
goos: linux
goarch: amd64
pkg: bitbucket.org/espinpro/go-engine/utils/rand
BenchmarkWeightIndex-4           2000000               747 ns/op              56 B/op          4 allocs/op
PASS
*/
var res interface{}

func BenchmarkWeightIndex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			result := WeightIndex(map[interface{}]uint64{
				"p1": 50,
				"p2": 50,
			})
			res = result
		}
	})
}
