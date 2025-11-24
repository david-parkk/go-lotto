package lotto

type WinningMachine struct {
	winningLotto *Lotto
	bonusNum     int
	winningInfos []WinningInfo
}

func NewWinningMachine(winningLotto *Lotto, bonusNum int, winningInfos []WinningInfo) *WinningMachine {
	return &WinningMachine{winningLotto: winningLotto, bonusNum: bonusNum, winningInfos: winningInfos}
}

func (winningMachine *WinningMachine) CheckWinning(lotto Lotto) *WinningInfo {
	matchCount := 0
	for _, num := range lotto.Numbers {
		for _, winNum := range winningMachine.winningLotto.Numbers {
			if num == winNum {
				matchCount++
			}
		}
	}

	hasBonus := false
	for _, num := range lotto.Numbers {
		if num == winningMachine.bonusNum {
			hasBonus = true
			break
		}
	}

	for _, info := range winningMachine.winningInfos {
		if info.MatchCount == matchCount {

			if info.HasBonus == hasBonus {
				return copyWinningInfo(info)
			}

			if !info.HasBonus {
				return copyWinningInfo(info)
			}
		}
	}

	return nil
}

func (winningMachine *WinningMachine) CheckWinnings(lottos []Lotto) []*WinningInfo {
	results := make([]*WinningInfo, 0, len(lottos))
	for _, lotto := range lottos {
		result := winningMachine.CheckWinning(lotto)
		results = append(results, result)
	}
	return results
}

func copyWinningInfo(src WinningInfo) *WinningInfo {
	newInfo := src
	return &newInfo
}
