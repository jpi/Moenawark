
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
}
