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
	"github.com/nleeper/goment"
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
			rows, err := conn.Query(context.Background(), `SELECT * FROM "contacts"`)
			if err != nil {
				log.Fatal("Failed at querying", err)
			}

			entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Contact])
			if err != nil {
				log.Fatal("Failed at collecting", err)
			}

			for index, entry := range entries {
				g, _ := goment.New(entry.UpdatedAt)

				fmt.Printf("%d. (Last updated: %v)\n", index, g.FromNow())
				fmt.Printf("Name: %v\n", entry.Name)
				if entry.Dob != nil {
					fmt.Printf("Dob: %v\n", entry.Dob.Format("2006-01-02"))
				}
				if entry.Address != nil {
					fmt.Printf("Address: %v\n", *entry.Address)
				}
				if entry.Phone != nil {
					fmt.Printf("Phone: %v\n", *entry.Phone)
				}
				if entry.Email != nil {
					fmt.Printf("Email: %v\n", *entry.Email)
				}
				fmt.Println()
			}

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
