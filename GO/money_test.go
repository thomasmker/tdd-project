package main

import (
	"reflect"
	s "tdd/stocks"
	"testing"
)

var bank s.Bank

func initExchangeRates()() {
	bank = s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	fiverDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiverDollars)
	portfolio = portfolio.Add(tenDollars)
	portfoilioInDollars, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, fifteenDollars, *portfoilioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, expectedValue, *actualValue)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)

	expectedValue := s.NewMoney(2200, "KRW")
	actualValue, err := portfolio.Evaluate(bank, "KRW")

	assertNil(t, err)
	assertEqual(t, expectedValue, *actualValue)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	initExchangeRates()
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	value, actualError := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func TestConversion(t *testing.T) {
	initExchangeRates()
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "USD")

	assertNil(t, err)
	assertEqual(t, s.NewMoney(12, "USD"), *actualConvertedMoney)

	bank.AddExchangeRate("EUR", "USD", 1.3)
	actualConvertedMoney, err = bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, s.NewMoney(13, "USD"), *actualConvertedMoney)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	initExchangeRates()
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	if actualConvertedMoney != nil {
		t.Errorf("Expected money to be nil, found: [%+v]", *actualConvertedMoney)
	}
	assertEqual(t, "EUR->Kalganid", err.Error())
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected error to be nil, found: [%+v]", actual)
	}
}
