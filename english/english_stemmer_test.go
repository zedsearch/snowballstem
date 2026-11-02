package english

import (
	"testing"

	snowballRuntime "github.com/zedsearch/snowballstem"
)

func TestStem(t *testing.T) {
	tests := map[string]string{
		"running":     "run",
		"generously":  "generous",
		"consign":     "consign",
		"consignment": "consign",
		"knightly":    "knight",
		"skies":       "sky",
		"dying":       "die",
		"news":        "news",
		"communism":   "communism",
		"gener":       "gener",
		"arsenal":     "arsenal",
		"caresses":    "caress",
		"ponies":      "poni",
		"café":        "café",
		"":            "",
	}
	for in, want := range tests {
		env := snowballRuntime.NewEnv(in)
		Stem(env)
		if got := env.Current(); got != want {
			t.Errorf("Stem(%q) = %q, want %q", in, got, want)
		}
	}
}
