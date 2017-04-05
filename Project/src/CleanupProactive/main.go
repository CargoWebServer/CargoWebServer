// CleanupProactive project main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

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
 * Return the list of active map.
 */
func getActives(db *sql.DB) (map[string]string, error) {
	// The map of active result.
	actives := make(map[string]string, 0)

	q := "SELECT Entry_COD,  max(ID_UNIT), UNIT_REF, max(UNIT_TIME)  FROM( \n"
	q += "SELECT rca_one.rca_entry.Entry_COD AS Entry_COD, rca_one.rca_unit.UNIT_REF as UNIT_REF, rca_one.rca_unit.ID_UNIT as ID_UNIT, \n"
	q += "	concat(rca_one.rca_unit.UNIT_DATE,' ',rca_one.rca_unit.UNIT_START_TIME) as UNIT_TIME \n"
	q += "	FROM rca_one.rca_entry \n"
	q += "	INNER JOIN rca_one.rca_unit ON rca_one.rca_entry.ID_Entry = rca_one.rca_unit.ID_UNIT \n"
	q += ") r \n"
	q += "GROUP BY r.Entry_COD, r.UNIT_REF \n"

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return actives, err
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return actives, err
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
			return actives, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		// Keep it in the map with it id.
		actives[result[0]+":"+result[2]] = result[1]
	}

	return actives, nil
}

/**
 * Return the list of piece to deletes.
 */
func getToRemove(db *sql.DB) ([]string, error) {
	// The list of entry to be remove.
	toRemove := make([]string, 0)

	// I will get the list of active values.
	actives, err := getActives(db)
	if err != nil {
		return toRemove, err
	}

	// The list of actual value in the db...
	q := "SELECT rca_one.rca_entry.ID_Entry, rca_one.rca_entry.Entry_COD, rca_one.rca_unit.UNIT_REF\n"
	q += "FROM rca_one.rca_entry\n"
	q += "INNER JOIN rca_one.rca_unit ON rca_one.rca_entry.ID_ENTRY = rca_one.rca_unit.ID_UNIT"

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

		// Keep it in the map with it id.
		if val, ok := actives[result[1]+":"+result[2]]; ok {
			if val != result[0] {
				// In that case I need to remove the data...
				if !Utility.Contains(toRemove, result[0]) {
					toRemove = append(toRemove, result[0])
				}
			} else {
				// Nothing to do here.
				//log.Println("--------> Active: ", result[2], result[0])
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
		}
	}
}
