package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var money int
	fmt.Println("구입 금액을 입력해 주세요.")
	_, err := fmt.Scanln(&money)
	for err != nil {
		log.Println("[ERROR] 숫자를 입력해 주세요.")
		_, err = fmt.Scanln(&money)
	}
	count := money / 1000
	fmt.Println(count, "개를 구매했습니다.")
	machine := NewLottoMachine(45, 6)
	lottos := machine.Generates(count, UserLotto)

	for _, lotto := range lottos {
		fmt.Println(lotto.String())
	}

	fmt.Println("당첨 번호를 입력해 주세요.")
	var numberStr string
	_, err = fmt.Scanln(&numberStr)
	for err != nil {
		log.Println("[ERROR] 입력을 읽을 수 없습니다.")
		_, err = fmt.Scanln(&numberStr)
	}
	split := strings.Split(numberStr, ",")
	numbers := make([]int, 0, 6)

	for _, each := range split {
		number, _ := strconv.Atoi(each)
		numbers = append(numbers, number)
	}
	winningLotto, err := NewLotto(numbers, WinningLotto)

	fmt.Println("보너스 번호를 입력해 주세요.")
	var bonusNum int
	_, err = fmt.Scanln(&bonusNum)
	for err != nil {
		log.Println("[ERROR] 입력을 읽을 수 없습니다.")
		_, err = fmt.Scanln(&bonusNum)
	}
	winningMachine := NewWinningMachine(winningLotto, bonusNum, WinningInfos)

	winningInfos := winningMachine.CheckWinnings(lottos)
	fmt.Println(winningInfos)

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
	for _, info := range WinningInfos { // 미리 정의한 등수 리스트
		key := formatWinningKey(&info)
		fmt.Printf("%s - %d개\n", key, stat[key])
	}
}

func formatWinningKey(info *WinningInfo) string {
	if info.MatchCount == 5 && info.HasBonus {
		return fmt.Sprintf("5개 일치, 보너스 볼 일치 (%d원)", info.Prize)
	}
	return fmt.Sprintf("%d개 일치 (%d원)", info.MatchCount, info.Prize)
}
