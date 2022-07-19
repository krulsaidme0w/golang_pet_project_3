package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	count                 = 1000
	maxOpenConns          = 1000
	tablesCount           = 6
	tablesWithForeignKeys = 3
	configPath            = "./"
	configName            = "./configs/config"
	baseDateFormat        = "2006-01-02"
)

func main() {
	if err := initConfig(configPath, configName); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("psql.host"),
		viper.GetString("psql.port"),
		viper.GetString("psql.user"),
		viper.GetString("psql.password"),
		viper.GetString("psql.dbname")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(maxOpenConns)

	err = fillLibraryDB(db, count)
	if err != nil {
		log.Fatal(err)
	}
}

func fillLibraryDB(db *sql.DB, count int) error {
	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(tablesCount - tablesWithForeignKeys)

	errs, _ := errgroup.WithContext(context.TODO())

	errs.Go(func() error {
		return fillUsersBulkLoad(db, count)
	})

	errs.Go(func() error {
		return fillAuthorBulkLoad(db, count)
	})

	errs.Go(func() error {
		return fillBookBulkLoad(db, count)
	})

	if err := errs.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := fillBookCopyBulkLoad(db, count); err != nil {
		log.Fatal(err)
	}

	if err := fillBookAuthorBulkLoad(db, count); err != nil {
		log.Fatal(err)
	}

	if err := fillHistoryBulkLoad(db, count); err != nil {
		log.Fatal(err)
	}

	log.Print("Time elapsed to fill data: ", time.Since(start))

	return nil
}

func fillUsersBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO library_user(username, email, password) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		username := faker.Username()
		email := faker.Email()
		password := hash(faker.Password())

		values = append(values, fmt.Sprintf("('%s', '%s', '%s')", username, email, password))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil {
		log.Fatal(err)
	}

	return nil
}

func fillAuthorBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO author(firstname, lastname, birthdate) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		firstname := faker.FirstName()
		lastname := faker.LastName()
		birthdate := faker.Date()

		values = append(values, fmt.Sprintf("('%s', '%s', '%s')", firstname, lastname, birthdate))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil {
		log.Fatal(err)
	}

	return nil
}

func fillBookBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO book(title, release_date) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		title := faker.Name()
		releaseDate := faker.Date()

		values = append(values, fmt.Sprintf("('%s', '%s')", title, releaseDate))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil {
		log.Fatal(err)
	}

	return nil
}

func fillBookCopyBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO book_copy(book_id) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		bookID := rand.Intn(count) + 1

		values = append(values, fmt.Sprintf("('%d')", bookID))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil {
		log.Fatal(err)
	}

	return nil
}

func fillBookAuthorBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO book_author(book_id, author_id) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		bookID := rand.Intn(count) + 1
		authorID := rand.Intn(count) + 1

		values = append(values, fmt.Sprintf("('%d', '%d')", bookID, authorID))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil && err.Error() != `pq: duplicate key value violates unique constraint "book_author_id"` {
		fmt.Println(err.Error())
	}

	return nil
}

func fillHistoryBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO history(library_user_id, book_copy_id, date_from, date_to) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		userID := rand.Intn(count) + 1
		bookCopyID := rand.Intn(count) + 1
		dateFrom := time.Unix(rand.Int63n(time.Now().Unix()), 0)
		dateTo := time.Unix(rand.Int63n(time.Now().Unix()), 0)

		if dateTo.Before(dateFrom) {
			tmp := dateFrom
			dateFrom = dateTo
			dateTo = tmp
		}

		dateFromStr := dateFrom.Format(baseDateFormat)
		dateToStr := dateTo.Format(baseDateFormat)

		values = append(values, fmt.Sprintf("('%d', '%d', '%s', '%s')", userID, bookCopyID, dateFromStr, dateToStr))
	}

	valuesStr := strings.Join(values, ", ")

	if _, err := db.Exec(fmt.Sprintf(query, valuesStr)); err != nil {
		log.Fatal(err)
	}

	return nil
}

func hash(str string) string {
	hash := sha256.Sum256([]byte(str))

	return hex.EncodeToString(hash[:])
}

func initConfig(path, name string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	return viper.ReadInConfig()
}
