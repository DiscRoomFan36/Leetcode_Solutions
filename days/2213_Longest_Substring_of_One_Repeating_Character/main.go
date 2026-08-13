package main

import "math/bits"

// i just copied this solution from this:
//
//      10 ms using two ordered sets and frequency count
//
// i care very little about leetcode anymore, spent all day on the other project.

const MAXN = 100000;

var boundaries = &HiBitset{};         // positions where s[i] != s[i-1]
var lengths     = &HiBitset{};        // existing run lengths
var count       = [MAXN + 1]uint32{}; // frequency of lengths

func addlen(length int) {
    count[length] += 1;
    if count[length] == 1 { lengths.Set(length); }
}

func removelen(length int) {
    count[length] -= 1;
    if count[length] == 0 { lengths.Unset(length); }
}

func longestRepeating(ss string, chars string, indices []int) []int {
    s := []byte(ss);
    n := len(s);

    // initialize boundaries
    boundaries.Set(0);
    boundaries.Set(n);
    prev := 0;
    for i := 1; i < n; i++ {
        if s[i-1] != s[i] {
            boundaries.Set(i);
            addlen(i - prev);
            prev = i;
        }
    }
    addlen(n - prev);

    // add/remove boundary at i
    update := func(i int, set bool) {
        if set {
            prev, next := boundaries.Prev(i), boundaries.Next(i+1);
            removelen(next - prev);
            addlen(i - prev);
            addlen(next - i);
            boundaries.Set(i);
        } else {
            prev, next := boundaries.Prev(i), boundaries.Next(i+1);
            removelen(i - prev);
            removelen(next - i);
            addlen(next - prev);
            boundaries.Unset(i);
        }
    }

    result := indices; // reuse slice for the output
    for i, idx := range indices {
        c := chars[i];
        if idx > 0 {
            old := s[idx-1] != s[idx];
            new := s[idx-1] != c;
            if old != new { update(idx, new); }
        }
        if idx+1 < n {
            old := s[idx] != s[idx+1];
            new := c != s[idx+1];
            if old != new { update(idx+1, new); }
        }
        s[idx] = c;
        result[i] = lengths.GetMax();
    }

    // cleanup for next execution
    boundaries.Clear();
    lengths.Clear();
    clear(count[:result[len(result)-1]+1]);

    return result;
}

const SHIFT = 6;
const MASK  = (1 << SHIFT) - 1;
const L2    = (MAXN + MASK) >> SHIFT;
const L1    = (L2 + MASK) >> SHIFT;
const L0    = (L1 + MASK) >> SHIFT;

type HiBitset struct {
    l2   [L2]uint64;
    l1   [L1]uint64;
    l0   uint64;
    maxv int;
}

func (h *HiBitset) Set(v int) {
    h.maxv      = max(h.maxv, v);
    idx2       := v >> SHIFT;
    idx1       := idx2 >> SHIFT;
    h.l2[idx2] |= 1 << (v & MASK);
    h.l1[idx1] |= 1 << (idx2 & MASK);
    h.l0       |= 1 << (idx1 & MASK);
}

func (h *HiBitset) Unset(v int) {
    idx2        := v >> SHIFT;
    idx1        := idx2 >> SHIFT;
    h.l2[idx2] &^= 1 << (v & MASK);
    if h.l2[idx2] == 0 {
        h.l1[idx1] &^= 1 << (idx2 & MASK);
        if h.l1[idx1] == 0 {
            h.l0 &^= 1 << (idx1 & MASK);
        }
    }
}

// smallest value >= v
func (h *HiBitset) Next(v int) int {
    idx2 := v >> SHIFT;
    idx1 := idx2 >> SHIFT;

    if next := h.l2[idx2] & (^uint64(0) << (v & MASK)); next != 0 {
        return (v &^ MASK) | bits.TrailingZeros64(next);
    }

    if next := h.l1[idx1] & (^uint64(0) << ((idx2 & MASK) + 1)); next != 0 {
        next2 := (idx1 << SHIFT) | bits.TrailingZeros64(next);
        return (next2 << SHIFT) | bits.TrailingZeros64(h.l2[next2]);
    }

    if next := h.l0 & (^uint64(0) << ((idx1 & MASK) + 1)); next != 0 {
        next1 := bits.TrailingZeros64(next);
        next2 := (next1 << SHIFT) | bits.TrailingZeros64(h.l1[next1]);
        return (next2 << SHIFT) | bits.TrailingZeros64(h.l2[next2]);
    }

    return -1;
}

// largest value < v
func (h *HiBitset) Prev(v int) int {
    idx2 := v >> SHIFT;
    idx1 := idx2 >> SHIFT;

    if prev := h.l2[idx2] & ((1 << (v & MASK)) - 1); prev != 0 {
        return (v &^ MASK) | (bits.Len64(prev) - 1);
    }

    if prev := h.l1[idx1] & ((1 << (idx2 & MASK)) - 1); prev != 0 {
        prev2 := (idx1 << SHIFT) | (bits.Len64(prev) - 1);
        return (prev2 << SHIFT) | (bits.Len64(h.l2[prev2]) - 1);
    }

    if prev := h.l0 & ((1 << idx1) - 1); prev != 0 {
        prev1 := bits.Len64(prev) - 1;
        prev2 := (prev1 << SHIFT) | (bits.Len64(h.l1[prev1]) - 1);
        return (prev2 << SHIFT) | (bits.Len64(h.l2[prev2]) - 1);
    }

    return -1;
}

func (h *HiBitset) GetMax() int {
    // if bt.l0 == 0 {
    // 	return -1;
    // }
    idx0 := bits.Len64(h.l0) - 1;
    idx1 := bits.Len64(h.l1[idx0]) - 1;
    idx2 := bits.Len64(h.l2[(idx0<<SHIFT)+idx1]) - 1;
    return (idx0<<(2*SHIFT) | idx1<<SHIFT | idx2);
}

func (h *HiBitset) Clear() {
    idx := h.maxv >> SHIFT;
    clear(h.l2[:idx+1]);
    clear(h.l1[:idx>>SHIFT+1]);
    h.l0 = 0;
    h.maxv = 0;
}
