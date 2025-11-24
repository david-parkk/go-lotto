package main

import (
	"math/rand"
	"time"
)

type LottoMachine struct {
	end    int
	length int
}

func NewLottoMachine(end int, length int) *LottoMachine {
	return &LottoMachine{end: end, length: length}
}

func (lottoMachine *LottoMachine) Generates(count int, lottoType LottoType) []Lotto {
	lotteries := make([]Lotto, 0, count)
	for i := 0; i < count; i++ {
		lotteries = append(lotteries, *lottoMachine.Generate(lottoType))
	}
	return lotteries
}

func (lottoMachine *LottoMachine) Generate(lottoType LottoType) *Lotto {
	rand.Seed(time.Now().UnixNano())

	numMap := make(map[int]bool)
	numbers := make([]int, 0, 6)

	for len(numbers) < 6 {
		num := rand.Intn(lottoMachine.end) + 1
		if !numMap[num] {
			numMap[num] = true
			numbers = append(numbers, num)
		}
	}

	lotto, _ := NewLotto(numbers, lottoType)
	return lotto
}
