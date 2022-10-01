package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestAddSlashes(t *testing.T) {
	t.Run("double quote", func(t *testing.T) {
		if res := PHP.Addslashes(`nano"nan"ya`); res != `nano\"nan\"ya` {
			t.Errorf(`expected nano\"nan\"ya, got %s`, res)
		}
	})
	t.Run("slash", func(t *testing.T) {
		if res := PHP.Addslashes(`nano\nan\ya`); res != `nano\\nan\\ya` {
			t.Errorf(`expected nano\\nan\\ya, got %s`, res)
		}
	})
}
