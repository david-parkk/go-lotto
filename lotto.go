package main

import (
	"errors"
	"fmt"
	"sort"
)

type Lotto struct {
	Numbers []int
	Type    LottoType
}

func NewLotto(numbers []int, Type LottoType) (*Lotto, error) {
	if len(numbers) != 6 {
		return nil, errors.New("[ERROR] 로또 번호는 6개여야 합니다.")
	}

	for _, n := range numbers {
		if n < 1 || n > 45 {
			return nil, errors.New("[ERROR] 로또 번호는 1~45 사이여야 합니다.")
		}
	}

	sort.Ints(numbers)
	return &Lotto{Numbers: numbers, Type: Type}, nil
}

func (lotto *Lotto) String() string {
	return fmt.Sprintf("%v", lotto.Numbers)
}
