package snowballstem

import "testing"

func TestEqSRespectsLimit(t *testing.T) {
	env := NewEnv("abcdef")
	env.Limit = 3
	if env.EqS("abcd") {
		t.Fatal("EqS matched past Limit")
	}
	if env.Cursor != 0 {
		t.Fatalf("cursor moved on failed match: %d", env.Cursor)
	}
	if !env.EqS("abc") || env.Cursor != 3 {
		t.Fatalf("EqS(abc) failed, cursor %d", env.Cursor)
	}
	if !env.EqS("") || env.Cursor != 3 {
		t.Fatalf("EqS(\"\") at limit should succeed, cursor %d", env.Cursor)
	}
}

func TestEqSMultiByte(t *testing.T) {
	env := NewEnv("héllo")
	if !env.EqS("hé") || env.Cursor != 3 {
		t.Fatalf("EqS(hé) failed, cursor %d", env.Cursor)
	}
}

var vowels = []byte{17, 65, 16, 1} // a e i o u y, min 'a' max 'y'

func TestGroupingBInvalidUTF8KeepsCursor(t *testing.T) {
	env := NewEnv("ab\xff")
	env.Cursor = len("ab\xff")
	if env.InGroupingB(vowels, 'a', 'y') {
		t.Fatal("InGroupingB matched invalid byte")
	}
	if env.Cursor != 3 {
		t.Fatalf("InGroupingB moved cursor to %d", env.Cursor)
	}
	if env.OutGroupingB(vowels, 'a', 'y') {
		t.Fatal("OutGroupingB matched invalid byte")
	}
	if env.Cursor != 3 {
		t.Fatalf("OutGroupingB moved cursor to %d", env.Cursor)
	}
}

func TestGroupingB(t *testing.T) {
	env := NewEnv("ba")
	env.Cursor = 2
	if !env.InGroupingB(vowels, 'a', 'y') || env.Cursor != 1 {
		t.Fatalf("InGroupingB(a) failed, cursor %d", env.Cursor)
	}
	if env.InGroupingB(vowels, 'a', 'y') || env.Cursor != 1 {
		t.Fatalf("InGroupingB(b) should fail without moving, cursor %d", env.Cursor)
	}
	if !env.OutGroupingB(vowels, 'a', 'y') || env.Cursor != 0 {
		t.Fatalf("OutGroupingB(b) failed, cursor %d", env.Cursor)
	}
}

func TestReplaceSAndInsert(t *testing.T) {
	env := NewEnv("hello")
	env.Cursor = 5
	env.Bra, env.Ket = 1, 3
	env.SliceFrom("EE")
	if env.Current() != "hEElo" || env.Cursor != 5 || env.Limit != 5 {
		t.Fatalf("got %q cursor %d limit %d", env.Current(), env.Cursor, env.Limit)
	}
	env.Insert(0, 0, "xx")
	if env.Current() != "xxhEElo" || env.Bra != 3 || env.Ket != 5 {
		t.Fatalf("got %q bra %d ket %d", env.Current(), env.Bra, env.Ket)
	}
}
