package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	}
	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("Dollar: %s\n", cotacao.Bid)); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	}

	fmt.Printf("Data for cotacao file written successfully.\n")

}
