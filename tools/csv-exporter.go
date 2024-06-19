package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://user1:zxc@localhost:5432/mydb")
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
		return
	}

	file, err := os.Open("binlist-data.csv")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if fields[0] == "bin" {
			continue
		}
		conn.Exec(context.Background(), "INSERT INTO beans (bin, issuer) VALUES ($1, $2) ON CONFLICT DO NOTHING", fields[0], fields[4])
	}
}
