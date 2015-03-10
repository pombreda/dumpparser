package hash

import (
    "math"
    "math/rand"
    "testing"
)

func TestCountMin(t *testing.T) {
    sketch := New(210, 1300)
    sketch1 := New(210, 1300)
    freq := make(map[uint32]uint32)

    rng := rand.New(rand.NewSource(42))
    for i := 0; i < 10000; i++ {
        h := rng.Uint32()
        sketch.Add(h, 1)
        sketch1.Add1(h)
        freq[h] += 1
    }

    // XXX Should test if error is within margin with some probability.
    for k, v := range freq {
        if math.Abs(float64(sketch.Get(k) - v)) > 4 {
            t.Errorf("difference too big: got %d, want %d", sketch.Get(k), v)
        }
        if sketch.Get(k) != sketch1.Get(k) {
            t.Errorf("different counts for Add and Add1",
                     sketch.Get(k), sketch1.Get(k))
        }
    }
}

func TestCountMinSum(t *testing.T) {
    a := New(25, 126)
    b := New(25, 126)
    sum := New(25, 126)
    rng := rand.New(rand.NewSource(42))

    for i := 0; i < 2000; i++ {
        h := rng.Uint32()
        a.Add1(h)
        b.Add(h, h % 100)
        sum.Add(h, h % 100 + 1)
    }
    a.Sum(b)

    for i := uint32(0); i < 126; i++ {
        if a.Get(i) != sum.Get(i) {
            t.Errorf("expected %d, got %d", sum.Get(i), a.Get(i))
        }
    }
}

func BenchmarkCountMinAdd(b *testing.B) {
    sketch := New(256, 256)

    rng := rand.New(rand.NewSource(42))
    for i := 0; i < b.N; i++ {
        for j := 0; j < 2000000; j++ {
            sketch.Add1(rng.Uint32())
        }
    }
}
