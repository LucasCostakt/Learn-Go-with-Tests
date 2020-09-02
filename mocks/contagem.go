package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

type Sleeper interface {
	Sleep()
}
type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.pausa(s.duracao)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}
func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
