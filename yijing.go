package yijing

// Tails represents a coin toss resulting in tails.
const Tails = 2

// Heads represents a coin toss resulting in heads.
const Heads = 3

// Coin represents the result of a coin toss (Heads or Tails).
type Coin int

// CoinSet represents the three Coins required to produce a Line.
type CoinSet [3]Coin

// Line represents a hexagram line.
type Line int

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
