package io

import (
	"fmt"
	"go-lotto/lotto"
	"io"
)

func PrintCount(count int, writer io.Writer) {
	fmt.Fprintf(writer, "%d개를 구매했습니다.\n", count)
}

func PrintLottos(lottos []lotto.Lotto, writer io.Writer) {
	for _, l := range lottos {
		fmt.Fprintln(writer, l.String())
	}
}

func PrintWinningStats(winningInfos []*lotto.WinningInfo, allInfos []lotto.WinningInfo, writer io.Writer) {
	stat := make(map[string]int)

	for _, r := range winningInfos {
		if r == nil {
			continue
		}
		key := formatWinningKey(r)
		stat[key]++
	}

	fmt.Fprintln(writer, "당첨 통계")
	fmt.Fprintln(writer, "---")
	for _, info := range allInfos {
		key := formatWinningKey(&info)
		fmt.Fprintf(writer, "%s - %d개\n", key, stat[key])
	}
}

func PrintProfitRate(writer io.Writer, totalSpent int, totalPrize int) {
	if totalPrize == 0 {
		fmt.Fprintln(writer, "총 수익률은 0%입니다.")
		return
	}

	rate := float64(totalPrize) / float64(totalSpent) * 100
	fmt.Fprintf(writer, "총 수익률은 %.1f%%입니다.\n", rate)
}

func formatWinningKey(info *lotto.WinningInfo) string {
	if info.MatchCount == 5 && info.HasBonus {
		return fmt.Sprintf("5개 일치, 보너스 볼 일치 (%d원)", info.Prize)
	}
	return fmt.Sprintf("%d개 일치 (%d원)", info.MatchCount, info.Prize)
}
