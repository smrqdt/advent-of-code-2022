package main

import (
	"math/big"
	"testing"
)

func parseMonkey() Monkey {
	lines := []string{
		"Monkey 1:",
		"		Starting items: 61, 70, 97, 64, 99, 83, 52, 87",
		"		Operation: new = old * 8",
		"		Test: divisible by 2",
		"			If true: throw to monkey 7",
		"			If false: throw to monkey 6",
	}
	return ParseMonkey(lines)
}

func TestParseMonkey(t *testing.T) {
	parseMonkey()
}

func TestNext(t *testing.T) {
	m := parseMonkey()
	// two := big.NewInt(2)
	// eight := big.NewInt(8)
	worryTrue := big.NewInt(9874664854842)
	worryFalse := big.NewInt(1561865156353101)
	if m.Next(worryTrue) != 7 {
		t.Log("Wrong true monkey")
		t.Fail()
	}
	if m.Next(worryFalse) != 6 {
		t.Log("Wrong false monkey")
		t.Fail()
	}
}

func TestOp(t *testing.T) {
	m := parseMonkey()

	num := big.NewInt(8)
	worry := big.NewInt(4165841635413654132)
	expected := big.NewInt(0).Mul(worry, num)
	if result := m.Op(worry); expected.Cmp(result) != 0 {
		t.Logf("got %v, expected %v", result, expected)
		t.Fail()
	}
}
