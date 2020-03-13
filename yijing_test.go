package yijing_test

import (
	"testing"

	"github.com/bitfield/yijing"
)

func TestLine(t *testing.T) {
	tcs := []struct {
		coins yijing.CoinSet
		want  yijing.Line
	}{
		{yijing.CoinSet{yijing.Tails, yijing.Tails, yijing.Tails}, yijing.OldYin},
		{yijing.CoinSet{yijing.Tails, yijing.Tails, yijing.Heads}, yijing.YoungYang},
		{yijing.CoinSet{yijing.Heads, yijing.Tails, yijing.Heads}, yijing.YoungYin},
		{yijing.CoinSet{yijing.Heads, yijing.Heads, yijing.Heads}, yijing.OldYang},
	}
	t.Parallel()
	for _, tc := range tcs {
		got := yijing.LineFromCoins(tc.coins)
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}
