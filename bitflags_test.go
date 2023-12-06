package bitflags_test

import (
	"testing"
	"github.com/imacks/bitflags-go"
)

type fruits int

// different types of fruits
const (
	apple fruits = 1 << iota
	orange
	banana
	durian
)

func TestSet(t *testing.T) {
	var basket fruits

	if bitflags.Has(basket, apple) {
		t.Fatal("should not have apple in empty basket")
	}
	basket = bitflags.Set(basket, apple)
	if !bitflags.Has(basket, apple) {
		t.Fatal("should have apple in basket")
	}
	// all operations are idempotent
	basket = bitflags.Set(basket, apple)
	if !bitflags.Has(basket, apple) {
		t.Fatal("should have apple in basket")
	}

	if bitflags.Has(basket, orange) {
		t.Fatal("should not have orange in basket")
	}
	basket = bitflags.Set(basket, orange)
	if !bitflags.Has(basket, orange) {
		t.Fatal("should have orange in basket")
	}

	basket = bitflags.Del(basket, apple)
	if bitflags.Has(basket, apple) {
		t.Fatal("should not have apple in basket")
	}
	// all operations are idempotent
	basket = bitflags.Del(basket, apple)
	if bitflags.Has(basket, apple) {
		t.Fatal("should not have apple in basket")
	}

	basket = bitflags.Set(basket, banana, durian)
	if !bitflags.Has(basket, banana) || !bitflags.Has(basket, durian) {
		t.Fatal("should have banana and durian in basket")
	}
	basket = bitflags.Toggle(basket, banana)
	if bitflags.Has(basket, banana) {
		t.Fatal("should not have banana in basket")
	}
	basket = bitflags.Toggle(basket, banana)
	if !bitflags.Has(basket, banana) {
		t.Fatal("should have banana in basket")
	}
}