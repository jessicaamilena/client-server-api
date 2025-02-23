package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	err = initDatabase(db)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		cotacaoHandler(w, r, db)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func initDatabase(db *sql.DB) error {
	createCotacoesTable := `create table if not exists cotacoes (id integer primary key autoincrement, bid varchar, created_at datetime default current_timestamp);`
	_, err := db.Exec(createCotacoesTable)
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ApiCtx, cancelApiCall := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancelApiCall()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ApiCtx, http.MethodGet, url, nil)
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

	bid := cotacaoResponse.USDBRL.Bid
	dbCtx, cancelDbReq := context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer cancelDbReq()

	if err := insertCotacao(dbCtx, db, bid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"bid": bid}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func insertCotacao(dbCtx context.Context, db *sql.DB, bid string) error {
	_, err := db.ExecContext(dbCtx, "insert into cotacoes (bid) values (?)", bid)
	return err
}
