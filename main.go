package main

import (
	"fmt"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var money int
	fmt.Println("구입 금액을 입력해 주세요.")
	_, err := fmt.Scanln(&money)
	for err != nil {
		log.Println("[ERROR] 숫자를 입력해 주세요.")
	}
	count := money / 1000
	fmt.Println(count, "개를 구매했습니다.")
	machine := NewLottoMachine(45, 6)
	lottos := machine.Generates(count, UserLotto)

	for _, lotto := range lottos {
		fmt.Println(lotto.String())
	}
}
