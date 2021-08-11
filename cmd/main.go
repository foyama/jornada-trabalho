package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


const (
	layoutISO = "15:04:05"
	tempoJornada = 528
)

func main() {
	entrada, _ := in("Entrada")
	saidaAlmoco, _ := in("Saida Almoço")
	entradaAlmoco, _ := in("Entrada Almoço")

	timeEntrada, _ := time.Parse(layoutISO, addSec(entrada))
	timeSaidaAlmoco, _ := time.Parse(layoutISO, addSec(saidaAlmoco))
	timeEntradaAlmoco, _ := time.Parse(layoutISO, addSec(entradaAlmoco))

	hour, minute, second := time.Now().Clock()
	horaAtual := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	timeHoraAtual, _ := time.Parse(layoutISO, horaAtual)

	segundosTrabalhados := timeSaidaAlmoco.Unix() - timeEntrada.Unix() + timeHoraAtual.Unix() - timeEntradaAlmoco.Unix()
	minutosTrabalhados := segundosTrabalhados / 60
	faltaTrabalhar := tempoJornada - minutosTrabalhados
	if faltaTrabalhar < 0 {
		fmt.Printf("Você já cumpriu sua jornada")
		return
	}

	horasSair := time.Now().Local().Add(time.Minute * time.Duration(faltaTrabalhar))
	hour, minute, _ = horasSair.Clock()
	fmt.Printf("Você sairá às: %02d:%02d", hour, minute)
}

func in(str string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Digite sua %v (HH:mm): ", str)
	return reader.ReadString('\n')
}

func addSec(str string) string {
	return strings.Replace(str, "\n", "", -1) + ":00"
}
