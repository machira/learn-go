package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Fixed number of args", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		expected := []int{}
		got := SumAll([][]int{})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})

	t.Run("single array", func(t *testing.T) {
		expected := []int{6}
		a := [][]int{{1, 2, 3}}
		got := SumAll(a)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})

	t.Run("many arrays", func(t *testing.T) {
		expected := []int{6, 10}
		a := [][]int{{1, 2, 3}, {1, 2, 3, 4}}
		got := SumAll(a)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([][]int{})
		want := []int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("single item slice", func(t *testing.T) {
		got := SumAllTails([][]int{{1}, {2}})
		want := []int{0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([][]int{{}, {3, 4, 5}})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("two element slice", func(t *testing.T) {
		got := SumAllTails([][]int{{1, 2}, {0, 9}})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
