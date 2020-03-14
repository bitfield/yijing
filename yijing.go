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

type LineTriple [3]Line

// These constants represent the various kinds of Line.
const (
	OldYin    Line = 6
	YoungYang Line = 7
	YoungYin  Line = 8
	OldYang   Line = 9
)

const (
	Yin  = false
	Yang = true
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

func LinesFromBytes(bs ByteSet) (lower, upper LineTriple) {
	coinsets := CoinsFromBytes(bs)
	for i, cs := range coinsets[0:3] {
		lower[i] = LineFromCoins(cs)
	}
	for i, cs := range coinsets[3:6] {
		upper[i] = LineFromCoins(cs)
	}
	return lower, upper
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

type Hexagram int

// Hexagram represents an individual hexagram. The Symbol shows the component
// lines, as a Unicode rune. The name of the hexagram is given in Chinese
// characters, Romanised Chinese and English.
type HexagramInfo struct {
	Symbol                  rune
	Chinese, Roman, English string
}

// Hexagrams represents the 64 I Ching hexagrams, in the King Wen sequence,
// indexed by number from 1-64.
var Hexagrams = []HexagramInfo{
	{},
	{'䷀', "乾", "qián", "The Creative"},
	{'䷁', "坤", "kūn", "The Receptive"},
	{'䷂', "屯", "zhūn", "Difficulty at the Beginning"},
}

var HexagramByTrigrams = map[Trigram]map[Trigram]Hexagram{
	Heaven: map[Trigram]Hexagram{
		Heaven: 1, Earth: 11, Thunder: 34, Water: 5, Mountain: 26, Wind: 9, Flame: 14, Lake: 43,
	},
	Earth: map[Trigram]Hexagram{
		Heaven: 12, Earth: 2, Thunder: 16, Water: 8, Mountain: 23, Wind: 20, Flame: 35, Lake: 45,
	},
	Thunder: map[Trigram]Hexagram{
		Heaven: 25, Earth: 24, Thunder: 51, Water: 3, Mountain: 27, Wind: 42, Flame: 21, Lake: 17,
	},
	Water: map[Trigram]Hexagram{
		Heaven: 6, Earth: 7, Thunder: 40, Water: 29, Mountain: 4, Wind: 59, Flame: 64, Lake: 47,
	},
	Mountain: map[Trigram]Hexagram{
		Heaven: 33, Earth: 15, Thunder: 62, Water: 39, Mountain: 52, Wind: 53, Flame: 56, Lake: 31,
	},
	Wind: map[Trigram]Hexagram{
		Heaven: 44, Earth: 46, Thunder: 32, Water: 48, Mountain: 18, Wind: 57, Flame: 50, Lake: 28,
	},
	Flame: map[Trigram]Hexagram{
		Heaven: 13, Earth: 36, Thunder: 55, Water: 63, Mountain: 22, Wind: 37, Flame: 30, Lake: 49,
	},
	Lake: map[Trigram]Hexagram{
		Heaven: 10, Earth: 19, Thunder: 54, Water: 60, Mountain: 41, Wind: 61, Flame: 38, Lake: 58,
	},
}

func HexagramFromTrigramPair(tp TrigramPair) Hexagram {
	return HexagramByTrigrams[tp.Lower][tp.Upper]
}

func IsYang(line Line) bool {
	return line == OldYang || line == YoungYang
}

var TrigramsByLineTypes = map[Trigram][]bool{
	Heaven:   {Yang, Yang, Yang},
	Earth:    {Yin, Yin, Yin},
	Thunder:  {Yin, Yin, Yang},
	Water:    {Yin, Yang, Yin},
	Mountain: {Yang, Yin, Yin},
	Wind:     {Yang, Yang, Yin},
	Flame:    {Yang, Yin, Yang},
	Lake:     {Yin, Yang, Yang},
}

func LineTypesEqual(a, b []bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TrigramFromLineTriple(input LineTriple) Trigram {
	lines := []bool{
		IsYang(input[0]),
		IsYang(input[1]),
		IsYang(input[2]),
	}
	for t, lt := range TrigramsByLineTypes {
		if LineTypesEqual(lt, lines) {
			return t
		}
	}
	return Heaven // can't happen
}

func HexagramFromBytes(bs ByteSet) Hexagram {
	lower, upper := LinesFromBytes(bs)
	tp := TrigramPair{
		Lower: TrigramFromLineTriple(lower),
		Upper: TrigramFromLineTriple(upper),
	}
	return HexagramFromTrigramPair(tp)
}
