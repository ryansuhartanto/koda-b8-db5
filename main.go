package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/ryansuhartanto/koda-b8-db5/models"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// PostgreSQL
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer conn.Close(context.Background())

	scanner := bufio.NewScanner(os.Stdin)

loop:
	for {
		fmt.Print("\033c")
		fmt.Print("" +
			"1. List\n" +
			"2. Add\n" +
			"3. Edit\n" +
			"4. Delete\n" +
			"\n" +
			"0. Exit\n" +
			"\n",
		)

		fmt.Print("Input: ")
		scanner.Scan()
		input := scanner.Text()

		selection, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		switch selection {
		case 1:
			rows, _ := conn.Query(context.Background(), `SELECT * FROM "contacts"`)
			entries, _ := pgx.CollectRows(rows, pgx.RowToStructByName[models.Contact])

			fmt.Println(entries)

			fmt.Println()
			fmt.Print("Enter to continue... ")
			scanner.Scan()

		case 0:
			break loop
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal("Buffer error", err)
	}
}
