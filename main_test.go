package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestRunLotto(t *testing.T) {
	input := "8000\n1,2,3,4,5,6\n7\n"
	reader := bufio.NewReader(strings.NewReader(input))
	writer := &bytes.Buffer{}

	RunLotto(reader, writer)

	output := writer.String()

	if !strings.Contains(output, "8개를 구매했습니다.") {
		t.Errorf("구매 개수 출력이 잘못됨, got: %s", output)
	}

	for i := 0; i < 8; i++ {
		if !strings.Contains(output, " ") {
			t.Errorf("로또 번호 출력 누락, got: %s", output)
		}
	}

	if !strings.Contains(output, "당첨 통계") {
		t.Errorf("당첨 통계 출력 누락, got: %s", output)
	}

	if !strings.Contains(output, "총 수익률") {
		t.Errorf("총 수익률 출력 누락, got: %s", output)
	}
}
