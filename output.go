package main

import (
	"fmt"
)

func PrintCount(count int) {
	fmt.Println(count, "개를 구매했습니다.")
}
func PrintLottos(lottos []Lotto) {
	for _, l := range lottos {
		fmt.Println(l.String())
	}
}

func PrintWinningStats(winningInfos []*WinningInfo, allInfos []WinningInfo) {
	stat := make(map[string]int)

	for _, r := range winningInfos {
		if r == nil {
			continue
		}
		key := formatWinningKey(r)
		stat[key]++
	}

	fmt.Println("당첨 통계")
	fmt.Println("---")
	for _, info := range allInfos {
		key := formatWinningKey(&info)
		fmt.Printf("%s - %d개\n", key, stat[key])
	}
}

func PrintProfitRate(totalSpent int, totalPrize int) {
	if totalPrize == 0 {
		fmt.Println("총 수익률은 0%입니다.")
		return
	}

	rate := float64(totalPrize) / float64(totalSpent) * 100
	fmt.Printf("총 수익률은 %.1f%%입니다.\n", rate)
}

func formatWinningKey(info *WinningInfo) string {
	if info.MatchCount == 5 && info.HasBonus {
		return fmt.Sprintf("5개 일치, 보너스 볼 일치 (%d원)", info.Prize)
	}
	return fmt.Sprintf("%d개 일치 (%d원)", info.MatchCount, info.Prize)
}
