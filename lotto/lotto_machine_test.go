package lotto

import (
	"testing"
)

func TestLottoMachine_Generate(t *testing.T) {
	end := 45
	length := 6
	machine := NewLottoMachine(end, length)
	machine.end = 45

	lotto := machine.Generate(UserLotto)

	if len(lotto.Numbers) != length {
		t.Errorf("로또 번호 개수 오류: got %d, want 6", len(lotto.Numbers))
	}

	numMap := make(map[int]bool)
	for _, n := range lotto.Numbers {
		if n < 1 || n > end {
			t.Errorf("로또 번호 범위 오류: got %d", n)
		}
		if numMap[n] {
			t.Errorf("중복된 번호 발견: %d", n)
		}
		numMap[n] = true
	}

	if lotto.Type != UserLotto {
		t.Errorf("로또 타입 오류: got %v, want %v", lotto.Type, UserLotto)
	}
}

func TestLottoMachine_Generates(t *testing.T) {
	end := 45
	length := 6
	machine := NewLottoMachine(end, length)
	machine.end = 45
	count := 10

	lottos := machine.Generates(count, UserLotto)

	if len(lottos) != count {
		t.Errorf("생성된 로또 개수 오류: got %d, want %d", len(lottos), count)
	}

	for i, lotto := range lottos {
		if len(lotto.Numbers) != 6 {
			t.Errorf("[%d] 로또 번호 개수 오류: got %d, want 6", i, len(lotto.Numbers))
		}
	}
}
