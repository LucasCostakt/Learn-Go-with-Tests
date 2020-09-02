package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até Vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Vai!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}
func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	tempoSpy := &TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Sleep}
	sleeper.Sleep()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}

type SleeperSpy struct {
	Chamadas int
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func (t *TempoSpy) Sleep(duracao time.Duration) {
	t.duracaoPausa = duracao
}
