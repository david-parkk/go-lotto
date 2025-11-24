package main

type WinningMachine struct {
	winningLotto *Lotto
	bonusNum     int
	winningInfos []WinningInfo
}

func NewWinningMachine(winningLotto *Lotto, bonusNum int, winningInfos []WinningInfo) *WinningMachine {
	return &WinningMachine{winningLotto: winningLotto, bonusNum: bonusNum, winningInfos: winningInfos}
}
