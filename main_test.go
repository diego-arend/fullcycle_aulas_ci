package main

import "testing"

func TestSome(t *testing.T) {
	total := Some(10, 10)

	if total != 20 {
		t.Errorf("Erro teste Some. O resultado deveria ser %d, mas o teste retornou: %d", 20, total)
	}
}
