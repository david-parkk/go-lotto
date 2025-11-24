package main

type WinningInfo struct {
	Rank       int
	MatchCount int
	HasBonus   bool
	Prize      int64
}

var WinningInfos = []WinningInfo{
	{Rank: 1, MatchCount: 6, HasBonus: false, Prize: 2000000000},
	{Rank: 2, MatchCount: 5, HasBonus: true, Prize: 30000000},
	{Rank: 3, MatchCount: 5, HasBonus: false, Prize: 1500000},
	{Rank: 4, MatchCount: 4, HasBonus: false, Prize: 50000},
	{Rank: 5, MatchCount: 3, HasBonus: false, Prize: 5000},
}

func NewWinningInfo(rank int, matchCount int, hasBonus bool, prize int64) *WinningInfo {
	return &WinningInfo{
		Rank:       rank,
		MatchCount: matchCount,
		HasBonus:   hasBonus,
		Prize:      prize,
	}
}
