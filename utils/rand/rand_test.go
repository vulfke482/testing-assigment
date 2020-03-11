package rand

import (
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFastRand(t *testing.T) {
	testcases := map[string]struct {
		min       uint32
		max       uint32
		minResult uint32
		maxResult uint32
	}{
		"simple_usage": {
			min:       15,
			max:       18,
			minResult: 15,
			maxResult: 18,
		},
		"zero_value": {
			min:       0,
			max:       18,
			minResult: 0,
			maxResult: 18,
		},
		"min_over_max": {
			min:       2,
			max:       0,
			minResult: 0,
			maxResult: 0,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := FastRand(test.min, test.max)
			if n < test.minResult || n > test.maxResult {
				t.Fatalf("Unexpected value, must be in range %d-%d: %d", test.minResult, test.maxResult, n)
			}
		})
	}
}

func TestFastRandFloat32(t *testing.T) {
	testcases := map[string]struct {
		precision float32
		min       float32
		max       float32
		minResult float32
		maxResult float32
	}{
		"simple_usage": {
			precision: 2,
			min:       1.5,
			max:       2.4,
			minResult: 1.5,
			maxResult: 2.4,
		},
		"zero_value": {
			precision: 1,
			min:       0,
			max:       18,
			minResult: 0,
			maxResult: 18,
		},
		"min_over_max": {
			precision: 2,
			min:       2.4,
			max:       1.4,
			minResult: 0,
			maxResult: 0,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := FastRandFloat32(test.precision, test.min, test.max)
			if n < test.minResult || n > test.maxResult {
				t.Fatalf("Unexpected value, must be in range %f-%f: %f", test.minResult, test.maxResult, n)
			}
		})
	}
}

func TestFastRandFloat64(t *testing.T) {
	testcases := map[string]struct {
		precision float64
		min       float64
		max       float64
		minResult float64
		maxResult float64
	}{
		"simple_usage": {
			precision: 2,
			min:       1.5,
			max:       2.4,
			minResult: 1.5,
			maxResult: 2.4,
		},
		"zero_value": {
			precision: 1,
			min:       0,
			max:       18,
			minResult: 0,
			maxResult: 18,
		},
		"min_over_max": {
			precision: 2,
			min:       2.4,
			max:       1.4,
			minResult: 0,
			maxResult: 0,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := FastRandFloat64(test.precision, test.min, test.max)
			if n < test.minResult || n > test.maxResult {
				t.Fatalf("Unexpected value, must be in range %f-%f: %f", test.minResult, test.maxResult, n)
			}
		})
	}
}

func TestCryptoRandomNumber(t *testing.T) {
	testcases := map[string]struct {
		min       int64
		max       int64
		minResult int64
		maxResult int64
	}{
		"simple_usage": {
			min:       15,
			max:       24,
			minResult: 15,
			maxResult: 24,
		},
		"zero_value": {
			min:       0,
			max:       18,
			minResult: 0,
			maxResult: 18,
		},
		"min_over_max": {
			min:       24,
			max:       14,
			minResult: 0,
			maxResult: 0,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := CryptoRandomNumber(test.min, test.max)
			if n < test.minResult || n > test.maxResult {
				t.Fatalf("Unexpected value, must be in range %d-%d: %d", test.minResult, test.maxResult, n)
			}
		})
	}
}

func TestRandomNumber(t *testing.T) {
	testcases := map[string]struct {
		min       int64
		max       int64
		minResult int64
		maxResult int64
	}{
		"simple_usage": {
			min:       15,
			max:       24,
			minResult: 15,
			maxResult: 24,
		},
		"zero_value": {
			min:       0,
			max:       18,
			minResult: 0,
			maxResult: 18,
		},
		"min_over_max": {
			min:       24,
			max:       14,
			minResult: 0,
			maxResult: 0,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := RandomNumber(test.min, test.max)
			if n < test.minResult || n > test.maxResult {
				t.Fatalf("Unexpected value, must be in range %d-%d: %d", test.minResult, test.maxResult, n)
			}
		})
	}
}

func TestRandomDigitString(t *testing.T) {
	testcases := map[string]struct {
		length int
	}{
		"simple_usage": {
			length: 15,
		},
		"zero_value": {
			length: 0,
		},
		"min_over_max": {
			length: 1,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			n := RandomDigitString(test.length)
			if len(n) != test.length {
				t.Fatalf("Unexpected len: %d != %d", test.length, len(n))
			}
		})
	}
}

/*
BenchmarkFastRandFloat32-4      100000000               19.7 ns/op             0 B/op          0 allocs/op
BenchmarkMathRandFloat32-4      30000000                56.5 ns/op             0 B/op          0 allocs/op
BenchmarkCryptoRandFloat32-4     2000000               863 ns/op              56 B/op          4 allocs/op
BenchmarkFastRand-4             100000000               20.7 ns/op             0 B/op          0 allocs/op
BenchmarkMathRand-4             30000000                50.3 ns/op             0 B/op          0 allocs/op
BenchmarkRandomDigitString-4     3000000               567 ns/op             382 B/op          4 allocs/op
*/
// benchSinkFloat32 prevents the compiler from optimizing away benchmark loops.
var benchSinkFloat32 float32
var mu sync.Mutex

// benchSinkFloat32 prevents the compiler from optimizing away benchmark loops.
var benchSink uint32
var benchString string

func BenchmarkFastRandFloat32(b *testing.B) {
	var precision float32 = 100
	var min float32 = 0.8
	var max float32 = 1.3

	b.RunParallel(func(pb *testing.PB) {
		s := float32(0)
		for pb.Next() {
			s += FastRandFloat32(precision, min, max)
		}
		mu.Lock()
		benchSinkFloat32 += s
		mu.Unlock()
	})
}

func BenchmarkMathRandFloat32(b *testing.B) {
	var precision float32 = 100
	var min float32 = 0.8
	var max float32 = 1.3

	b.RunParallel(func(pb *testing.PB) {
		s := float32(0)
		for pb.Next() {
			s += mathRandFloat32(precision, min, max)
		}
		mu.Lock()
		benchSinkFloat32 += s
		mu.Unlock()
	})
}

func BenchmarkCryptoRandFloat32(b *testing.B) {
	var precision float32 = 100
	var min float32 = 0.8
	var max float32 = 1.3

	b.RunParallel(func(pb *testing.PB) {
		s := float32(0)
		for pb.Next() {
			s += cryptoRandFloat32(precision, min, max)
		}
		mu.Lock()
		benchSinkFloat32 += s
		mu.Unlock()
	})
}

func BenchmarkFastRand(b *testing.B) {
	var min uint32 = 8
	var max uint32 = 13

	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += FastRand(min, max)
		}
		atomic.AddUint32(&benchSink, s)
	})
}

func BenchmarkMathRand(b *testing.B) {
	var min uint32 = 8
	var max uint32 = 13

	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += mathRand(min, max)
		}
		atomic.AddUint32(&benchSink, s)
	})
}

func BenchmarkRandomDigitString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		b.StartTimer()
		var s []string
		for pb.Next() {
			s = append(s, RandomDigitString(1))
		}
		b.StopTimer()

		mu.Lock()
		benchString = strings.Join(s, ",")
		mu.Unlock()
	})
}

func mathRand(min, max uint32) uint32 {
	if min > max {
		return 0
	}
	return (rand.Uint32() % max) + min
}

func mathRandFloat32(precision, min, max float32) float32 {
	min *= precision
	max *= precision
	return float32(RandomNumber(int64(min), int64(max))) / precision
}

func cryptoRandFloat32(precision, min, max float32) float32 {
	min *= precision
	max *= precision
	return float32(CryptoRandomNumber(int64(min), int64(max))) / precision
}
