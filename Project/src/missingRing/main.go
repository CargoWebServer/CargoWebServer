// missingRing project main.go
package main

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"strconv"

	_ "github.com/alexbrainman/odbc"
	"golang.org/x/net/html"
)

const (
	// The addresse of the server.
	DB_HOST = "tcp(10.2.128.96:1433)"
	DB_NAME = "DB_Inspection_Dimension"
	DB_USER = "dbprog"
	DB_PASS = "dbprog"
)

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func importData(db *sql.DB, serial string, partNumber string, position string, diameter float64, ovStd string, ovSal string, revision int64, creationTime string) {

	// First of all I will retreive the modele id.
	var query_1 = "SELECT TOP 1 [id] FROM [DB_Inspection_Dimension].[dbo].[Bushing] WHERE [STANDARD] = ? AND [OVERSIZE_NUMBER] = ? ORDER BY REVISION DESC"
	rows, err := db.Query(query_1, ovStd, ovSal)
	var bushingId int
	if err == nil {
		for rows.Next() {
			rows.Scan(&bushingId)
		}
	} else {
		log.Println("---> ", err)
	}

	// Now I will retreive the model id.
	var query_2 = "SELECT [id] FROM [DB_Inspection_Dimension].[dbo].[Modele] WHERE [nom]=?"
	rows, err = db.Query(query_2, partNumber)
	var modelId int
	if err == nil {
		for rows.Next() {
			rows.Scan(&modelId)
		}
	} else {
		log.Println("---> ", err)
	}

	// Now I will try to get the actual diameter.
	var query_3 = "SELECT count(piece_id) FROM [DB_Inspection_Dimension].[dbo].[OversizeDimension] WHERE [position_id] = ? AND [piece_id] = ?"
	var count int
	rows, err = db.Query(query_3, position, serial)
	if err == nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	} else {
		log.Println("---> ", err)
	}

	if count == 0 {
		log.Println("---> import ", partNumber, serial, " position ", position, " diameter ", diameter, "std", ovStd, "sal", ovSal, "rev.", revision)
		// So here I will import the actual value inside the database.
		var query_4 = "insert into [DB_Inspection_Dimension].[dbo].[OversizeDimension] "
		query_4 += "([busing_id],[position_id],[piece_id],[produit_id],[ACTUAL_DIAMETER],[CREATION_DATE],[EMPLOYE_ID]) "
		query_4 += "VALUES (?,?,?,?,?,?,?)"
		db.Exec(query_4, bushingId, position, serial, modelId, diameter, creationTime, "mm006819")
		//log.Println(query_4)
		//log.Println("VALUES ('", bushingId, "','", position, "','", serial, "',", modelId, ",", diameter, ",'"+creationTime, "', 'mm006819')")
	}
}

func main() {
	/** Connect with ODBC here... **/
	connectionString := "driver={sql server};"
	connectionString += "server=mon-sql-v01;"
	connectionString += "database=DB_Inspection_Dimension;"
	connectionString += "uid=dbprog;"
	connectionString += "pwd=dbprog;"
	connectionString += "port=1433;"
	connectionString += "charset=UTF8;"

	db, err := sql.Open("odbc", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	/*dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=Windows1252"
	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatal(err)
	}*/

	// So first of all i will read the list of html file from the folder tmp.
	var files []string

	root := "C:\\Temp\\messages"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		htm, err := ioutil.ReadFile(file)
		if err == nil {
			doc, _ := html.Parse(strings.NewReader(string(htm)))
			body, _ := getBody(doc)
			var creationTime string
			for c := body.FirstChild; c != nil; c = c.NextSibling {
				if c.Data == "div" {
					if len(creationTime) == 0 {
						creationTime = c.FirstChild.Data
					}
				} else if c.Data == "table" {
					var serial string
					var partNumber string
					var actualDiameter float64
					var rev int64
					var ovStd string
					var ovSal string
					var pos string
					row := 0
					tBody := c.FirstChild
					for tr := tBody.FirstChild; tr != nil; tr = tr.NextSibling {
						col := 0
						for td := tr.FirstChild; td != nil; td = td.NextSibling {
							if td.FirstChild != nil {
								if row == 2 && col == 0 {
									partNumber = strings.TrimSpace(td.FirstChild.Data)
								} else if row == 2 && col == 1 {
									rev, _ = strconv.ParseInt(td.FirstChild.Data, 10, 64)
								} else if row == 2 && col == 2 {
									serial = strings.TrimSpace(td.FirstChild.Data)
								} else if row == 2 && col == 3 {
									ovStd = strings.TrimSpace(td.FirstChild.Data)
								} else if row == 2 && col == 4 {
									ovSal = strings.TrimSpace(td.FirstChild.Data)
								} else if row == 2 && col == 5 {
									pos = strings.TrimSpace(td.FirstChild.Data)
								} else if row == 2 && col == 6 {
									actualDiameter, _ = strconv.ParseFloat(td.FirstChild.Data, 64)
								}
								col++
							}
						}
						row++
					}
					if len(serial) > 0 {
						importData(db, serial, partNumber, pos, actualDiameter, ovStd, ovSal, rev, creationTime)
					}
				}
			}
		}
	}

	defer db.Close()
}
