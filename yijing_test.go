package yijing_test

import (
	"testing"

	y "github.com/bitfield/yijing"
	"github.com/google/go-cmp/cmp"
)

func TestLine(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		coins y.CoinSet
		want  y.Line
	}{
		{y.CoinSet{y.Tails, y.Tails, y.Tails}, y.OldYin},
		{y.CoinSet{y.Tails, y.Tails, y.Heads}, y.YoungYang},
		{y.CoinSet{y.Heads, y.Tails, y.Heads}, y.YoungYin},
		{y.CoinSet{y.Heads, y.Heads, y.Heads}, y.OldYang},
	}
	for _, tc := range tcs {
		got := y.LineFromCoins(tc.coins)
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}

func TestLinesFromBytes(t *testing.T) {
	t.Parallel()
	// These bytes represent the following bit sequence:
	// 000 100 001 011 100 101 (plus trailing waste bits)
	input := y.ByteSet{0b00010000, 0b10111001, 0b01000000}
	wantLower := y.LineTriple{
		y.OldYin,
		y.YoungYang,
		y.YoungYang,
	}
	wantUpper := y.LineTriple{
		y.YoungYin,
		y.YoungYang,
		y.YoungYin,
	}
	gotLower, gotUpper := y.LinesFromBytes(input)
	if !cmp.Equal(wantLower, gotLower) {
		t.Error(cmp.Diff(wantLower, gotLower))
	}
	if !cmp.Equal(wantUpper, gotUpper) {
		t.Error(cmp.Diff(wantUpper, gotUpper))
	}
}

func TestHexagramFromBytes(t *testing.T) {
	t.Parallel()
	// These bytes represent the following bit sequence:
	// 000 100 001 011 100 101 (plus trailing waste bits)
	input := y.ByteSet{0b00010000, 0b10111001, 0b01000000}
	want := y.Hexagram(60)
	got := y.HexagramFromBytes(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCoinsFromBytes(t *testing.T) {
	t.Parallel()
	// These bytes represent the following bit sequence:
	// 000 100 001 011 100 101 (plus trailing waste bits)
	input := y.ByteSet{0b00010000, 0b10111001, 0b01000000}
	want := y.CoinSet6{
		y.CoinSet{y.Tails, y.Tails, y.Tails},
		y.CoinSet{y.Heads, y.Tails, y.Tails},
		y.CoinSet{y.Tails, y.Tails, y.Heads},
		y.CoinSet{y.Tails, y.Heads, y.Heads},
		y.CoinSet{y.Heads, y.Tails, y.Tails},
		y.CoinSet{y.Heads, y.Tails, y.Heads},
	}
	var got y.CoinSet6 = y.CoinsFromBytes(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestHexagramByTrigrams(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input y.TrigramPair
		want  y.Hexagram
	}{
		{y.TrigramPair{y.Thunder, y.Water}, 3},
		{y.TrigramPair{y.Mountain, y.Wind}, 53},
		{y.TrigramPair{y.Flame, y.Thunder}, 55},
		{y.TrigramPair{y.Water, y.Earth}, 7},
		{y.TrigramPair{y.Heaven, y.Lake}, 43},
		{y.TrigramPair{y.Earth, y.Heaven}, 12},
		{y.TrigramPair{y.Wind, y.Wind}, 57},
		{y.TrigramPair{y.Water, y.Flame}, 64},
	}
	for _, tc := range tcs {
		got := y.HexagramByTrigrams[tc.input.Lower][tc.input.Upper]
		if !cmp.Equal(tc.want, got) {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}

func TestTrigramFromLineTriple(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input y.LineTriple
		want  y.Trigram
	}{
		{y.LineTriple{y.OldYin, y.YoungYang, y.YoungYang}, y.Lake},
		{y.LineTriple{y.OldYang, y.OldYang, y.YoungYin}, y.Wind},
		{y.LineTriple{y.YoungYin, y.YoungYang, y.YoungYin}, y.Water},
		{y.LineTriple{y.OldYin, y.OldYang, y.YoungYang}, y.Lake},
		{y.LineTriple{y.YoungYang, y.OldYin, y.OldYin}, y.Mountain},
		{y.LineTriple{y.OldYin, y.OldYin, y.OldYang}, y.Thunder},
		{y.LineTriple{y.YoungYin, y.YoungYin, y.YoungYin}, y.Earth},
		{y.LineTriple{y.OldYang, y.OldYang, y.YoungYang}, y.Heaven},
		{y.LineTriple{y.YoungYang, y.OldYin, y.OldYang}, y.Flame},
	}
	for _, tc := range tcs {
		got := y.TrigramFromLineTriple(tc.input)
		if !cmp.Equal(tc.want, got) {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}
