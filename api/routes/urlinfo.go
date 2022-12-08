package routes

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/raystainerz/gourl/helpers"
	"github.com/asaskevich/govalidator"
)

type dbinfo struct {
	URL string `json:"url"`
}
// Urlinfo function : takes in a url and returns its information (clicks, etc)
func Urlinfo(ctx *fiber.Ctx) error {
	
	db, err := sql.Open("mysql", "root:8f#Ne65tKo<z@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
	body := &request{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}


	if !govalidator.IsURL(body.URL) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Can't do that :)"})
	}

	// enforce HTTPS, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
	var redurl string
	var clicks int
	var types string
	// print("here look:")
	// print(body.URL)
	row := db.QueryRow("SELECT url, clicks, type FROM trackurl WHERE url = ?;", body.URL)
	row.Scan(&redurl, &clicks, &types)
	
	return ctx.JSON(fiber.Map{"url": redurl, "clicks": clicks, "type": types})
}
