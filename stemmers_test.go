package snowballstem_test

import (
	"testing"

	snowballRuntime "github.com/zedsearch/snowballstem"
	"github.com/zedsearch/snowballstem/arabic"
	"github.com/zedsearch/snowballstem/danish"
	"github.com/zedsearch/snowballstem/dutch"
	"github.com/zedsearch/snowballstem/english"
	"github.com/zedsearch/snowballstem/finnish"
	"github.com/zedsearch/snowballstem/french"
	"github.com/zedsearch/snowballstem/german"
	"github.com/zedsearch/snowballstem/hungarian"
	"github.com/zedsearch/snowballstem/irish"
	"github.com/zedsearch/snowballstem/italian"
	"github.com/zedsearch/snowballstem/norwegian"
	"github.com/zedsearch/snowballstem/porter"
	"github.com/zedsearch/snowballstem/portuguese"
	"github.com/zedsearch/snowballstem/romanian"
	"github.com/zedsearch/snowballstem/russian"
	"github.com/zedsearch/snowballstem/spanish"
	"github.com/zedsearch/snowballstem/swedish"
	"github.com/zedsearch/snowballstem/tamil"
	"github.com/zedsearch/snowballstem/turkish"
)

var stemmers = map[string]func(*snowballRuntime.Env) bool{
	"arabic": arabic.Stem, "danish": danish.Stem, "dutch": dutch.Stem,
	"english": english.Stem, "finnish": finnish.Stem, "french": french.Stem,
	"german": german.Stem, "hungarian": hungarian.Stem, "irish": irish.Stem,
	"italian": italian.Stem, "norwegian": norwegian.Stem, "porter": porter.Stem,
	"portuguese": portuguese.Stem, "romanian": romanian.Stem, "russian": russian.Stem,
	"spanish": spanish.Stem, "swedish": swedish.Stem, "tamil": tamil.Stem,
	"turkish": turkish.Stem,
}

func TestStemmersKnownWords(t *testing.T) {
	tests := []struct{ lang, in, want string }{
		{"french", "continuellement", "continuel"},
		{"german", "aufeinanderfolgenden", "aufeinanderfolg"},
		{"spanish", "chiquillos", "chiquill"},
		{"russian", "вавилон", "вавилон"},
		{"italian", "abbandonata", "abbandon"},
		{"porter", "caresses", "caress"},
		{"dutch", "lichamelijk", "licham"},
	}
	for _, tt := range tests {
		env := snowballRuntime.NewEnv(tt.in)
		stemmers[tt.lang](env)
		if got := env.Current(); got != tt.want {
			t.Errorf("%s Stem(%q) = %q, want %q", tt.lang, tt.in, got, tt.want)
		}
	}
}

// Malformed and edge-case input must never panic in any stemmer.
func TestStemmersNoPanic(t *testing.T) {
	inputs := []string{"", "a", "\xff", "ab\xffcd", "\xc3", "é\xff", "ççç", "ﬃ", "ааа\xffааа", "இஇ\x80", "ĞĞğ", "ﻻﻻ"}
	for lang, stem := range stemmers {
		for _, in := range inputs {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s panicked on %q: %v", lang, in, r)
					}
				}()
				stem(snowballRuntime.NewEnv(in))
			}()
		}
	}
}
