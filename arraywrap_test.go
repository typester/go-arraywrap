package arraywrap

import (
	"testing"
)

func TestLen(t *testing.T) {
	list_s := []string{"foo", "bar", "bazz"}

	a := New(list_s)

	if a.Len() != 3 {
		t.Errorf("invalid length expected:%d got:%d", 3, a.Len())
	}

	list_i := []int{1, 2, 3, 4, 5}
	a = New(list_i)

	if a.Len() != 5 {
		t.Errorf("invalid length expected:%d got:%d", 5, a.Len())
	}
}

func TestIndex(t *testing.T) {
	list_s := []string{"foo", "bar", "bazz"}

	a := New(list_s)

	v0, ok := a.Index(0).(string)
	if !ok {
		t.Errorf("should be string")
	}
	if v0 != "foo" {
		t.Errorf("unexpected value: %s", v0)
	}

	list_i := []int{1, 2, 3, 4, 5}
	a = New(list_i)

	v2, ok := a.Index(2).(int)
	if !ok {
		t.Errorf("should be int")
	}
	if v2 != 3 {
		t.Errorf("unexpected value: %d", v2)
	}
}

func TestNegativeIndex(t *testing.T) {
	list_s := []string{"foo", "bar", "bazz"}

	a := New(list_s)

	v, ok := a.Index(-1).(string)

	if !ok {
		t.Errorf("should be string")
	}
	if v != "bazz" {
		t.Errorf("unexpected value: %s", v)
	}
}
