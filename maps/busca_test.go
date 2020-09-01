package Mapa

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		if resultado == nil {
			t.Fatal("é esperado que um erro seja obtido.")
		}
	})
}

func TestAdiciona(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		novaDefinicao := "nova definição"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErro(t, err, ErrPalavraInexistente)
	})
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s',  esperado '%s'", resultado, definicao)
	}
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}
func TestUpdate(t *testing.T) {
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario := Dicionario{palavra: definicao}
	novaDefinicao := "nova definição"

	dicionario.Atualiza(palavra, novaDefinicao)

	comparaDefinicao(t, dicionario, palavra, novaDefinicao)
}
func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}
