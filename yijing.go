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

type Trigram int

const (
	Heaven Trigram = iota
	Earth
	Thunder
	Water
	Mountain
	Wind
	Flame
	Lake
)

type TrigramPair struct {
	Lower, Upper Trigram
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
	{'䷁', "坤", "kūn", "The Receptive"},
	{'䷂', "屯", "zhūn", "Difficulty at the Beginning"},
}

var HexagramNumByTrigrams = map[Trigram]map[Trigram]int{
	Heaven: map[Trigram]int{
		Heaven: 1, Earth: 11, Thunder: 34, Water: 5, Mountain: 26, Wind: 9, Flame: 14, Lake: 43,
	},
	Earth: map[Trigram]int{
		Heaven: 12, Earth: 2, Thunder: 16, Water: 8, Mountain: 23, Wind: 20, Flame: 35, Lake: 45,
	},
	Thunder: map[Trigram]int{
		Heaven: 25, Earth: 24, Thunder: 51, Water: 3, Mountain: 27, Wind: 42, Flame: 21, Lake: 17,
	},
	Water: map[Trigram]int{
		Heaven: 6, Earth: 7, Thunder: 40, Water: 29, Mountain: 4, Wind: 59, Flame: 64, Lake: 47,
	},
	Mountain: map[Trigram]int{
		Heaven: 33, Earth: 15, Thunder: 62, Water: 39, Mountain: 52, Wind: 53, Flame: 56, Lake: 31,
	},
	Wind: map[Trigram]int{
		Heaven: 44, Earth: 46, Thunder: 32, Water: 48, Mountain: 18, Wind: 57, Flame: 50, Lake: 28,
	},
	Flame: map[Trigram]int{
		Heaven: 13, Earth: 36, Thunder: 55, Water: 63, Mountain: 22, Wind: 37, Flame: 30, Lake: 49,
	},
	Lake: map[Trigram]int{
		Heaven: 10, Earth: 19, Thunder: 54, Water: 60, Mountain: 41, Wind: 61, Flame: 38, Lake: 58,
	},
}

func HexagramFromTrigramPair(tp TrigramPair) Hexagram {
	return Hexagrams[HexagramNumByTrigrams[tp.Lower][tp.Upper]]
}
