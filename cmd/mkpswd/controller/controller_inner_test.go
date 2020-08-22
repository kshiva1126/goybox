package controller

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	tests := []struct {
		name            string
		input, expected charFlags
	}{
		{
			"1",
			[]string{"n"},
			[]string{"n"},
		},
		{
			"2",
			[]string{"n", "l"},
			[]string{"l", "n"},
		},
		{
			"3",
			[]string{"n", "l", "s"},
			[]string{"l", "n", "s"},
		},
		{
			"4",
			[]string{"n", "l", "s", "u"},
			[]string{"l", "u", "n", "s"},
		},
		{
			"5",
			[]string{"n", "l", "s", "u", "c"},
			[]string{"l", "u", "n", "s", "c"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.sort()
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Exptected return of %v, but got %v", tt.expected, tt.input)
			}
		})
	}
}

func Test_concatPasswordLetters(t *testing.T) {
	tests := []struct {
		name     string
		input    charFlags
		expected string
	}{
		{
			"1",
			[]string{"n"},
			"0123456789",
		},
		{
			"2",
			[]string{"l", "n"},
			"abcdefghijklmnopqrstuvwxyz0123456789",
		},
		{
			"3",
			[]string{"l", "n", "s"},
			"abcdefghijklmnopqrstuvwxyz0123456789!\"#$%'()*,./:;<=>?@[]^_`{|}~",
		},
		{
			"4",
			[]string{"l", "u", "n", "s"},
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%'()*,./:;<=>?@[]^_`{|}~",
		},
		{
			"5",
			[]string{"l", "u", "n", "s", "c"},
			"abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!\"#$%'()*,./:;<=>?@[]^_`{}~",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			letters := concatPasswordLetters(tt.input)
			if letters != tt.expected {
				t.Errorf("Exptected return of %v, but got %v", tt.expected, letters)
			}
		})
	}
}

func Test_outputsPasswords(t *testing.T) {
	var nc int = 8
	var np int = 1

	tests := []struct {
		name, input string
	}{
		{
			"1",
			"0123456789",
		},
		{
			"2",
			"abcdefghijklmnopqrstuvwxyz0123456789",
		},
		{
			"3",
			"abcdefghijklmnopqrstuvwxyz0123456789!\"#$%'()*,./:;<=>?@[]^_`{|}~",
		},
		{
			"4",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%'()*,./:;<=>?@[]^_`{|}~",
		},
		{
			"5",
			"abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!\"#$%'()*,./:;<=>?@[]^_`{}~",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			outputs, err := outputsPasswords(tt.input, &nc, &np)
			if err != nil {
				t.Errorf("Expected return of nil, but got %v", err)
			}
			if len(outputs) != np {
				t.Errorf("Expected length of outputs is %v, but got %v", np, len(outputs))
			}
			for _, output := range outputs {
				if len(output) != nc {
					t.Errorf("Expected length of output is %v, but got %v", nc, len(output))
				}
			}
		})
	}
}
