package yijing

// Tails represents a coin toss resulting in tails.
const Tails = 2

// Heads represents a coin toss resulting in heads.
const Heads = 3

// Coin represents the result of a coin toss (Heads or Tails).
type Coin int

// CoinSet represents the three Coins required to produce a Line.
type CoinSet [3]Coin

type CoinSet6 [6]CoinSet

type ByteSet [3]byte

// Line represents a hexagram line.
type Line int

type LineSet [6]Line

// These constants represent the various kinds of Line.
const (
	OldYin    Line = 6
	YoungYang Line = 7
	YoungYin  Line = 8
	OldYang   Line = 9
)

// LineFromCoins takes a CoinSet and returns the equivalent Line.
func LineFromCoins(cs CoinSet) Line {
	return Line(cs[0] + cs[1] + cs[2])
}

func CoinsFromBytes(bs ByteSet) CoinSet6 {
	var coins []Coin
	for _, b := range bs {
		for i := 7; i >= 0; i-- {
			if b>>i&1 == 1 {
				coins = append(coins, Heads)
			} else {
				coins = append(coins, Tails)
			}
		}
	}
	return CoinSet6{
		CoinSet{coins[0], coins[1], coins[2]},
		CoinSet{coins[3], coins[4], coins[5]},
		CoinSet{coins[6], coins[7], coins[8]},
		CoinSet{coins[9], coins[10], coins[11]},
		CoinSet{coins[12], coins[13], coins[14]},
		CoinSet{coins[15], coins[16], coins[17]},
	}
}

func LinesFromBytes(bs ByteSet) LineSet {
	var ls LineSet
	coinsets := CoinsFromBytes(bs)
	for i, cs := range coinsets {
		ls[i] = LineFromCoins(cs)
	}
	return ls
}

// Hexagram represents an individual hexagram. The Symbol shows the component
// lines, as a Unicode rune. The name of the hexagram is given in Chinese
// characters, Romanised Chinese and English.
type Hexagram struct {
	Symbol                  rune
	Chinese, Roman, English string
}

// Hexagrams represents the 64 I Ching hexagrams, in the King Wen sequence,
// indexed by number from 1-64.
var Hexagrams = []Hexagram{
	{},
	{'䷀', "乾", "qián", "The Creative"},
}
