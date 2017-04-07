// CleanupProactive project main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	// The addresse of the server.
	DB_HOST = "tcp(10.2.128.121:3306)"
	DB_NAME = "RCA_ONE"
	DB_USER = "toto"
	DB_PASS = ""
)

/**
 * Remove a string from a slice.
 */
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

/**
 * Return the list of piece to deletes.
 */
func getToRemove(db *sql.DB) ([]string, error) {

	// The list of entry to be remove.
	activeTimeMap := make(map[string]time.Time, 0)
	activeIdMap := make(map[string]string, 0)
	toRemove := make([]string, 0)

	// The list of actual value in the db...
	q := "SELECT DISTINCT rca_one.rca_entry.Entry_COD, rca_one.rca_entry.ID_Entry, rca_one.rca_unit.UNIT_REF, UNIX_TIMESTAMP(concat(rca_one.rca_unit.UNIT_DATE,' ',rca_one.rca_unit.UNIT_START_TIME)) as UNIT_TIME\n"
	q += "FROM rca_one.rca_entry\n"
	q += "JOIN rca_one.rca_unit ON rca_one.rca_entry.ID_ENTRY = rca_one.rca_unit.ID_UNIT"

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return toRemove, err
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return toRemove, err
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return toRemove, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		serialRegex, _ := regexp.Compile("^\\S{3}\\d{4}\\S$")
		if serialRegex.MatchString(result[2]) {
			// Keep it in the map with it id.
			featureId := result[0] + result[2]
			unixTime, _ := strconv.Atoi(strings.Split(result[3], ".")[0])
			featureTime := time.Unix(int64(unixTime), 0)

			if val, ok := activeTimeMap[featureId]; ok {

				if val.Before(featureTime) {
					// So here we found a newer piece...
					activeTimeMap[featureId] = featureTime
					if activeIdMap[featureId] == "1112" {
						log.Println("OUUP---------------> ", featureId)
					}
					if !Utility.Contains(toRemove, activeIdMap[featureId]) {
						toRemove = append(toRemove, activeIdMap[featureId])
					}
					// I will remove the actual feature.
					activeIdMap[featureId] = result[1]
				} else {
					// In that case the encounter piece is older than the
					// actual one, so I will put in to remove.
					if !Utility.Contains(toRemove, result[1]) {
						toRemove = append(toRemove, result[1])
					}
				}
			} else {
				activeTimeMap[featureId] = featureTime
				activeIdMap[featureId] = result[1]
			}
		} else {
			if !Utility.Contains(toRemove, result[1]) {
				toRemove = append(toRemove, result[1])
			}
		}
	}
	return toRemove, nil
}

func main() {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	getToRemove(db)

	// The list of values to remove.
	toRemove, err := getToRemove(db)

	if err == nil {
		for i := 0; i < len(toRemove); i++ {
			// I will remove the entry from the rca_one.rca_unit with key ID_UNIT
			//Delete
			var stmt0 *sql.Stmt
			var err error
			stmt0, err = db.Prepare("DELETE FROM rca_one.rca_unit WHERE ID_UNIT=?")
			if err == nil {
				res, err := stmt0.Exec(toRemove[i])
				if err == nil {
					affect, err := res.RowsAffected()
					// Print the affected rows...
					if err == nil {
						fmt.Println(affect)
					}
				} else {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}

			// Now the list of feature associated with that id
			var stmt1 *sql.Stmt
			stmt1, err = db.Prepare("DELETE FROM rca_one.rca_entry WHERE ID_ENTRY=?")
			if err == nil {
				res, err := stmt1.Exec(toRemove[i])
				if err == nil {
					affect, err := res.RowsAffected()
					// Print the affected rows...
					if err == nil {
						fmt.Println(affect)
					}
				} else {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}
			log.Println("-----------> remove number: ", toRemove[i])
		}
	}

}
