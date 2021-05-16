package funcsql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/url")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Database!")
	return db, nil
}

func GetAllPath(db *sql.DB) (map[string]string, error) {
	pathFromURL := []mapURL{}
	rows, err := db.Query("SELECT * FROM mapURL")
	if err != nil {
		return buildMap(pathFromURL), err
	}
	defer rows.Close()
	for rows.Next() {
		var aURL mapURL
		err = rows.Scan(&aURL.Path, &aURL.URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		pathFromURL = append(pathFromURL, aURL)
	}
	return buildMap(pathFromURL), nil
}

func InsertPath(db *sql.DB, path string, url string) error {
	insert, err := db.Query(fmt.Sprintf("INSERT into mapURL value ('%s','%s')", path, url))
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func buildMap(pathURLs []mapURL) map[string]string {
	ret := make(map[string]string)
	for _, item := range pathURLs {
		ret[item.Path] = item.URL
	}
	return ret
}

type mapURL struct {
	Path string
	URL  string
}
