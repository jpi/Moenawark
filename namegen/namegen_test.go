package namegen

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewMarkov(t *testing.T) {
	input := "foo\nbar\nbaz\nquux"
	m, err := NewMarkov(strings.NewReader(input), 2)
	if err != nil {
		t.Errorf("NewMarkov: unexpected error %v", err)
	}
	empty := []rune{'f', 'b', 'b', 'q'}
	if !reflect.DeepEqual(empty, m.dict[""]) {
		t.Errorf("NewMarkov: dict[\"\"] = %v, want %v", m.dict[""], empty)
	}
	oo := []rune{'\n'}
	if !reflect.DeepEqual(oo, m.dict["oo"]) {
		t.Errorf("NewMarkov: dict[\"oo\"] = %v, want %v", m.dict["oo"], oo)
	}
}
