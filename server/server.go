package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type CotacaoApiResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Bid string `json:"bid"`
}

func main() {
	db, err := sql.Open("sqlite3", "file:cotacao.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", cotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var cotacaoResponse CotacaoApiResponse
	err = json.NewDecoder(resp.Body).Decode(&cotacaoResponse)
	if err != nil {
		panic(err)
	}

	// bid := cotacaoResponse.USDBRL.Bid

	// dbCtx, cancelDbReq := context.WithTimeout(r.Context(), 10*time.Millisecond)
	// defer cancelDbReq()

}
