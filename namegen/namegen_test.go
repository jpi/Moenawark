package namegen

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewMarkov(t *testing.T) {
	input := "foo\nbar\nbaz\nquux"
	dict, err := loadDict(strings.NewReader(input), 2)
	if err != nil {
		t.Errorf("NewMarkov: unexpected error %v", err)
	}
	empty := []rune{'f', 'b', 'b', 'q'}
	if !reflect.DeepEqual(empty, dict[""]) {
		t.Errorf("NewMarkov: dict[\"\"] = %v, want %v", dict[""], empty)
	}
	oo := []rune{'\n'}
	if !reflect.DeepEqual(oo, dict["oo"]) {
		t.Errorf("NewMarkov: dict[\"oo\"] = %v, want %v", dict["oo"], oo)
	}
	ba := []rune{'r', 'z'}
	if !reflect.DeepEqual(ba, dict["ba"]) {
		t.Errorf("NewMarkov: dict[\"ba\"] = %v, want %v", dict["ba"], ba)
	}
}
