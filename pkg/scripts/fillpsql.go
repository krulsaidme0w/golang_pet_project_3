package scripts

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
)

const (
	tablesCount = 3
)

func FillLibraryDB(db *sql.DB, count int) error {
	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(tablesCount)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := fillUsersBulkLoad(db, count); err != nil {
			log.Fatal(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := fillAuthorBulkLoad(db, count); err != nil {
			log.Fatal(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := fillBookBulkLoad(db, count); err != nil {
			log.Fatal(err)
		}
	}(wg)

	wg.Wait()

	log.Print("Time elapsed to fill data: ", time.Since(start))

	return nil
}

func fillUsersBulkLoad(db *sql.DB, count int) error {
	const query = "INSERT INTO library_user(id, username, email, password) VALUES %s;"

	values := make([]string, 0, count)

	for i := 0; i < count; i++ {
		username := faker.Username()
		email := faker.Email()
		password := Hash(faker.Password())
		id := Hash(email)

		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s')", id, username, email, password))
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
