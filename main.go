package main

import (
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	money := ReadMoney(os.Stdin)

	count := money / 1000
	PrintCount(count)

	machine := NewLottoMachine(45, 6)
	lottos := machine.Generates(count, UserLotto)

	PrintLottos(lottos)

	numbers := ReadWinningNumbers(os.Stdin)
	winningLotto, _ := NewLotto(numbers, WinningLotto)

	bonusNum := ReadBonusNumber(os.Stdin)
	winningMachine := NewWinningMachine(winningLotto, bonusNum, WinningInfos)

	winningInfos := winningMachine.CheckWinnings(lottos)

	PrintWinningStats(winningInfos, WinningInfos)
	totalPrize := SumPrizes(winningInfos)
	PrintProfitRate(money, totalPrize)
}
