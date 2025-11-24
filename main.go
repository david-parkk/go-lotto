package main

import (
	"go-lotto/io"
	"go-lotto/lotto"
	stdio "io"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	RunLotto(os.Stdin, os.Stdout)
}

func RunLotto(reader stdio.Reader, writer stdio.Writer) {
	money := io.ReadMoney(reader, writer)

	count := money / 1000
	io.PrintCount(count, writer)

	machine := lotto.NewLottoMachine(45, 6)
	lottos := machine.Generates(count, lotto.UserLotto)

	io.PrintLottos(lottos, writer)

	numbers := io.ReadWinningNumbers(reader, writer)
	winningLotto, _ := lotto.NewLotto(numbers, lotto.WinningLotto)

	bonusNum := io.ReadBonusNumber(reader, writer)
	winningMachine := lotto.NewWinningMachine(winningLotto, bonusNum, lotto.WinningInfos)

	winningInfos := winningMachine.CheckWinnings(lottos)

	io.PrintWinningStats(winningInfos, lotto.WinningInfos, writer)
	totalPrize := lotto.SumPrizes(winningInfos)
	io.PrintProfitRate(writer, money, totalPrize)
}
