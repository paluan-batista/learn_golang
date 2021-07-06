package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenarioDeTeste := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rua das flores", "Rua"},
		{"Estrada Paulista", "Estrada"},
		{"Rodovia Imigrantes", "Rodovia"},
		{"Alameda Paulista", "tipo invalido"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVenida Paulista", "Avenida"},
		{"estrada Paulista", "Estrada"},
		{"", "tipo invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.retornoEsperado)
		}
	}
}
