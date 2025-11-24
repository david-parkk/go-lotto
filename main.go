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
	NewWinningMachine(winningLotto, bonusNum, WinningInfos)
}
