package main

import (
	"go-lotto/io"
	"go-lotto/lotto"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	money := io.ReadMoney(os.Stdin)

	count := money / 1000
	io.PrintCount(count)

	machine := lotto.NewLottoMachine(45, 6)
	lottos := machine.Generates(count, lotto.UserLotto)

	io.PrintLottos(lottos)

	numbers := io.ReadWinningNumbers(os.Stdin)
	winningLotto, _ := lotto.NewLotto(numbers, lotto.WinningLotto)

	bonusNum := io.ReadBonusNumber(os.Stdin)
	winningMachine := lotto.NewWinningMachine(winningLotto, bonusNum, lotto.WinningInfos)

	winningInfos := winningMachine.CheckWinnings(lottos)

	io.PrintWinningStats(winningInfos, lotto.WinningInfos)
	totalPrize := lotto.SumPrizes(winningInfos)
	io.PrintProfitRate(money, totalPrize)
}
