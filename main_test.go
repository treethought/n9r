package main

import "testing"

func Test_numeronyze(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"with trailing space", "dog   ", "d1g"},
		{"numeronyzer", "numeronyzer", "n9r"},
		{"kubernetes", "kubernetes", "k8s"},
		{"internationalization", "internationalization", "i18n"},
		{"kim with space", "kim kardashian", "k11n"},
		{"kim no space", "kimkardashian", "k11n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numeronyze(tt.in); got != tt.want {
				t.Errorf("numeronyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
