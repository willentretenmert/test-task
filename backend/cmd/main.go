package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"backend/internal/models"
	"backend/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "postgres://user1:zxc@host.docker.internal:5432/mydb")
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	log.Printf("database connection pool established")
	defer pool.Close()

	storage, err := storage.InitStorage(pool)
	if err != nil {
		log.Fatalf("error initializing local storage: %v", err)
	}

	server := &fasthttp.Server{
		Handler: requestHandler(pool, storage),
	}

	log.Printf("server started on port 8111")
	if err := server.ListenAndServe(":8111"); err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}

func requestHandler(pool *pgxpool.Pool, storage *storage.Storage) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/bin":
			bin, err := strconv.Atoi(string(ctx.QueryArgs().Peek("bin")))

			issuer := getIssuerFromDatabase(bin, storage)
			if len(issuer) == 0 {
				issuer, err = getIssuerFromBinList(bin)
				if err != nil {
					log.Printf("error getting issuer name from binlist: %v", err)
					return
				}
				go func() {
					if err = setNewIssuer(pool, bin, issuer); err != nil {
						log.Printf("error inserting issuer name: %v", err)
						return
					}
				}()
			}
			if err != nil {
				log.Printf("error getting issuer name: %v", err)
				return
			}

			_, err = ctx.WriteString(issuer)
			if err != nil {
				log.Printf("error writing issuer name: %v", err)
				return
			}
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetContentType("text/plain; charset=utf-8")
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
}

func getIssuerFromDatabase(bin int, storage *storage.Storage) (issuer string) {
	issuer = storage.GetIssuer(bin)
	return
}

func getIssuerFromBinList(bin int) (issuer string, err error) {
	var cardInfo models.CardInfo
	baseUrl := "https://binlist.io/lookup/" // очень долгая ручка, ответ 500-10000мс
	response, err := http.Get(baseUrl + strconv.Itoa(bin))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil && len(body) == 0 {
		log.Printf("error requesting binlist, empty response")
		return "", err
	}

	err = json.Unmarshal(body, &cardInfo)
	if err != nil {
		log.Printf("error parsing json")
		return "", err
	}
	if cardInfo.Issuer.Name == "UNKNOWN" {
		return "", errors.New("issuer name is unknown")
	}
	return cardInfo.Issuer.Name, nil
}

func setNewIssuer(pool *pgxpool.Pool, bin int, issuer string) error {
	_, err := pool.Query(context.Background(), `
		INSERT INTO beans (bin, issuer)
		VALUES ($1, $2)
		ON CONFLICT (bin)
		DO UPDATE SET issuer = EXCLUDED.issuer
	`, bin, issuer)

	if err != nil {
		return err
	}
	return nil
}
