package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result representa os resultados de cada requisição
type Result struct {
	StatusCode int
}

func main() {
	// Parâmetros CLI
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de requisições simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: a URL é obrigatória.")
		flag.Usage()
		return
	}

	// Canal para coletar os resultados
	results := make(chan Result, *requests)

	// Sincronizar as goroutines
	var wg sync.WaitGroup

	// Medir o tempo total
	startTime := time.Now()

	// Criar goroutines para enviar as requisições
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// Obter a próxima requisição
				reqCount := *requests - len(results)
				if reqCount <= 0 {
					break
				}

				// Enviar requisição HTTP
				resp, err := http.Get(*url)
				if err != nil {
					results <- Result{StatusCode: 0}
					continue
				}
				results <- Result{StatusCode: resp.StatusCode}
				resp.Body.Close()
			}
		}()
	}

	// Aguardar todas as goroutines terminarem
	wg.Wait()
	close(results)

	// Calcular o tempo total
	totalTime := time.Since(startTime)

	// Processar os resultados
	statusCount := make(map[int]int)
	for result := range results {
		statusCount[result.StatusCode]++
	}

	// Relatório
	fmt.Println("\n--- Relatório de Teste de Carga ---")
	fmt.Printf("URL Testada: %s\n", *url)
	fmt.Printf("Número Total de Requests: %d\n", *requests)
	fmt.Printf("Concorrência: %d\n", *concurrency)
	fmt.Printf("Tempo Total: %s\n", totalTime)
	fmt.Printf("Requests com Status 200: %d\n", statusCount[200])
	fmt.Println("\nOutros Status HTTP:")
	for status, count := range statusCount {
		if status != 200 {
			fmt.Printf("  %d: %d\n", status, count)
		}
	}
	fmt.Println("----------------------------------")
}
