package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		expectVal int
	}{
		{"TestingWhenRegular", 10000, 10000},
		{"TestingWhenNegValues", -1, 0},
		{"TestingWhenOneElement", 1, 1},
		{"TestingToEmpty", 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testVal := generateRandomElements(test.size)
			assert.Len(t, testVal, test.expectVal)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expectVal int
	}{
		{"TestingWhenRegular", []int{1, 2, 3, 4, 5}, 5},
		{"TestingWhenValuesTheSame", []int{1, 1, 1, 1, 1}, 1},
		{"TestingWhenOneElement", []int{1}, 1},
		{"TestingToEmpty", []int{}, 0},
		{"TestingToNil", nil, 0},
	}

	for _, test := range tests {
		testVal := maximum(test.input)
		assert.Equal(t, test.expectVal, testVal)
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expectVal int
	}{
		{"TestingWhenRegular", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"TestingWhenTheSame", []int{1, 1, 1, 1, 1, 1, 1, 1}, 1},
		{"TestingWhenOneElement", []int{1}, 1},
		{"TestingToEmpty", []int{}, 0},
		{"TestingToNil", nil, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testVal := maxChunks(test.input)
			assert.Equal(t, test.expectVal, testVal)
		})
	}
}
