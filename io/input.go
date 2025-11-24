package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func ReadMoney(r io.Reader) int {
	var money int
	for {
		fmt.Println("구입 금액을 입력해 주세요.")
		_, err := fmt.Fscanf(r, "%d\n", &money)
		if err != nil {
			fmt.Println("[ERROR] 숫자를 입력해 주세요.")
			var discard string
			fmt.Fscanln(r, &discard)
			continue
		}
		if money%1000 != 0 {
			fmt.Println("[ERROR] 1000원 단위로 입력해 주세요.")
			continue
		}
		return money
	}
}

func ReadWinningNumbers(r io.Reader) []int {
	reader := bufio.NewReader(r)
	for {
		fmt.Println("당첨 번호를 입력해 주세요.")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("[ERROR] 입력을 읽을 수 없습니다.")
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.Split(input, ",")
		if len(parts) != 6 {
			log.Println("[ERROR] 숫자 6개를 입력해야 합니다.")
			continue
		}

		numbers := make([]int, 0, 6)
		valid := true
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				log.Println("[ERROR] 숫자 형식이 올바르지 않습니다.")
				valid = false
				break
			}
			numbers = append(numbers, num)
		}

		if valid {
			return numbers
		}
	}
}

func ReadBonusNumber(r io.Reader) int {
	reader := bufio.NewReader(r)

	for {
		fmt.Print("보너스 번호를 입력해 주세요: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("[ERROR] 입력을 읽을 수 없습니다.")
			continue
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Println("[ERROR] 숫자를 입력해 주세요.")
			continue
		}

		return number
	}
}
