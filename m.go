package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("measurements.txt") // Abertura do arquivo
	if err != nil {
		panic(err)
	}
	defer measurements.Close() // Fechar o file descriptor para liber os recursos na memória

	dados := make(map[string]Measurement) // Define um slice(array) de map <string, Measurement>

	scanner := bufio.NewScanner(measurements) // usado para iterar entre as linhas do arquivo
	for scanner.Scan() {
		rawData := scanner.Text()                  // linha inteira
		semicolon := strings.Index(rawData, ";")   // procura o index do ;
		location := rawData[:semicolon]            // :semicolon estamos recortando do inicio da string até o index do semicolon
		rawTemp := rawData[semicolon+1:]           // semicolon+1: nesse caso estamos retortando do index para o final da string, o +1 é para ignorar o ;
		temp, _ := strconv.ParseFloat(rawTemp, 64) // convert a string em float o segundo parametro define qual será o tipo float

		measurement, ok := dados[location]
		if !ok { // quando o location não existir no map o measurement deve ser iniciado com os valores padrões do temp
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}

	// order o map
	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", name, measurement.Min, measurement.Sum/float64(measurement.Count), measurement.Max)
	}
	fmt.Printf("}\n")

	fmt.Println(time.Since(start))

	// for name, measurement := range dados { // iterando entre os valores do map
	// 	fmt.Printf("%s: %#+v\n", name, measurement) // %s: %#+v\n faz com que as propriedades do objeto sejam detalhadas no console
	// }
}
