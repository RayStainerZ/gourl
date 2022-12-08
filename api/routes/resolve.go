package routes

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var redurl string

// Resolve function : takes in a url slug and redirects to the original url
func Resolve(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	//db connection
	db, err := sql.Open("mysql", "root:8f#Ne65tKo<z@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	//get original url from slug url
	row := db.QueryRow("SELECT url FROM new_table WHERE slug = ?;", url)
	row.Scan(&redurl)

	//increment clicks by 1 in trackurl table
	increment, err := db.Query("UPDATE trackurl SET clicks = clicks + 1 WHERE url = ?;", redurl)
	if err !=nil {
        panic(err.Error())
    }
	defer increment.Close()

	//get ip of the client
	ip := ctx.IP()
	//insert ip into userlogs table 
	insert, err := db.Query("INSERT INTO userlogs (ip, url) VALUES (?, ?);", ip, redurl)
	if err !=nil {
		panic(err.Error())
	}
	defer insert.Close()
	//redirect to the original url
	return ctx.Redirect(redurl, 301)
}
