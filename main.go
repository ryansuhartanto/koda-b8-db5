package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

		fmt.Println()

		switch selection {
		case 1:
			list(conn, scanner)
		case 2:
			add(conn, scanner)
		case 4:
			delete(conn, scanner)

		case 0:
			break loop
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal("Buffer error", err)
	}
}

func list(conn *pgx.Conn, scanner *bufio.Scanner) {
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
}

func add(conn *pgx.Conn, scanner *bufio.Scanner) {
	var (
		name string
		dob  *time.Time

		address,
		phone,
		email *string
	)

	var input string

	fmt.Print("Name: ")
	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())
	name = input

	for {
		fmt.Print("Dob (2006-01-02): ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())
		if input == "" {
			break
		}

		time, err := time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Fprint(os.Stderr, "Failed parse", err)
			continue
		}

		dob = &time
		break
	}

	fmt.Print("Address: ")
	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())
	if input != "" {
		value := input
		address = &value
	}

	fmt.Print("Phone: ")
	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())
	if input != "" {
		value := input
		phone = &value
	}

	fmt.Print("Email: ")
	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())
	if input != "" {
		value := input
		email = &value
	}

	fmt.Println()

	args := pgx.NamedArgs{
		"name":    name,
		"dob":     dob,
		"address": address,
		"phone":   phone,
		"email":   email,
	}

	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO "contacts" ("name", "dob", "address", "phone", "email") `+
			`VALUES (@name, @dob, @address, @phone, @email)`,
		args,
	)
	if err != nil {
		log.Fatalln("Failed at executing", err)
	}

	fmt.Println("Success!")
	fmt.Print("Enter to continue... ")
	scanner.Scan()
}

func delete(conn *pgx.Conn, scanner *bufio.Scanner) {
	var id int64

	for {
		fmt.Print("ID: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, "Failed parse", err)
			continue
		}

		id = value
		break
	}

	var exists bool
	err := conn.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM "contacts" WHERE id = $1)`, id).Scan(&exists)
	if err != nil {
		log.Fatalln("Failed at ID checking", err)
	}

	fmt.Println()

	if !exists {
		fmt.Println("ID does not exist.")
		scanner.Scan()
		return
	}

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err = conn.Exec(
		context.Background(),
		`DELETE FROM "contacts" WHERE id = @id`,
		args,
	)
	if err != nil {
		log.Fatalln("Failed at executing", err)
	}

	fmt.Println("Success!")
	fmt.Print("Enter to continue... ")
	scanner.Scan()
}
