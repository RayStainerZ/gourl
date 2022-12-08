package routes

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var redurl string
func Resolve(ctx *fiber.Ctx) error {
	url := ctx.Params("url")
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	row := db.QueryRow("SELECT url FROM new_table WHERE slug = ?;", url)
	row.Scan(&redurl)
	// print("----->")
	// print(redurl)

	return ctx.Redirect(redurl, 301)
}
