package io

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadMoney(reader io.Reader, writer io.Writer) int {
	var money int
	for {
		fmt.Fprintln(writer, "구입 금액을 입력해 주세요.")
		_, err := fmt.Fscanf(reader, "%d\n", &money)
		if err != nil {
			fmt.Fprintln(writer, "[ERROR] 숫자를 입력해 주세요.")
			var discard string
			fmt.Fscanln(reader, &discard)
			continue
		}
		if money%1000 != 0 {
			fmt.Println("[ERROR] 1000원 단위로 입력해 주세요.")
			continue
		}
		return money
	}
}

func ReadWinningNumbers(reader io.Reader, writer io.Writer) []int {
	new_reader := bufio.NewReader(reader)
	for {
		fmt.Fprintln(writer, "당첨 번호를 입력해 주세요.")
		input, err := new_reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(writer, "[ERROR] 입력을 읽을 수 없습니다.")
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.Split(input, ",")
		if len(parts) != 6 {
			fmt.Fprintln(writer, "[ERROR] 숫자 6개를 입력해야 합니다.")
			continue
		}

		numbers := make([]int, 0, 6)
		valid := true
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Fprintln(writer, "[ERROR] 숫자 형식이 올바르지 않습니다.")
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

func ReadBonusNumber(reader io.Reader, writer io.Writer) int {
	new_reader := bufio.NewReader(reader)

	for {
		fmt.Fprintln(writer, "보너스 번호를 입력해 주세요: ")
		input, err := new_reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(writer, "[ERROR] 입력을 읽을 수 없습니다.")
			continue
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(writer, "[ERROR] 숫자를 입력해 주세요.")
			continue
		}

		return number
	}
}
