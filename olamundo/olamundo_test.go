package main

import (
	"testing"
)

var testador = []struct {
	nome     string
	innome   string
	inidioma string
	out      string
}{
	{"diz olá para as pessoas", "Chris", "", "Olá, Chris"},
	{"'Mundo' como padrão para 'string' vazia", "", "", "Olá, Mundo"},
	{"em espanhol", "Elodie", espanhol, "Hola, Elodie"},
	{"em francês", "Chico", frances, "Bonjour, Chico"},
}

func verificaMensagemCorreta(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestOla(t *testing.T) {

	for _, tt := range testador {
		t.Run(tt.nome, func(t *testing.T) {
			s := Ola(tt.innome, tt.inidioma)
			verificaMensagemCorreta(t, s, tt.out)
		})
	}

}
