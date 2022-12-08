package routes

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/raystainerz/gourl/helpers"
)

type request struct {
	URL         string        `json:"url"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
}
var id string

// Shorten function : takes in a url and returns a shortened url
func Shorten(ctx *fiber.Ctx) error {
	
	//db connection
	db, err := sql.Open("mysql", "root:8f#Ne65tKo<z@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	body := &request{}

	//parse the body, throw error if not JSON
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	//check if url is valid
	if !govalidator.IsURL(body.URL) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Can't do that :)"})
	}

	// add HTTPS if not present
	body.URL = helpers.EnforceHTTP(body.URL)
	
	//create a new random string for the url
	id = helpers.String(7)

	//insert new url into new_table
	insert, err := db.Query("INSERT INTO new_table (slug, url ) VALUES(?,?)", id, body.URL)
    if err !=nil {
        panic(err.Error())
    }
    defer insert.Close()
    fmt.Println("New url added!")
	
	//check if url already exists in trackurl table
	var redurl string
	row := db.QueryRow("SELECT url FROM trackurl WHERE url = ?;", body.URL)
	row.Scan(&redurl)
	//if url doesnt exists in trackurl table, add it
	if redurl == "" {
		trackurl, err := db.Query("INSERT INTO trackurl (url, clicks, type ) VALUES(?,?,?)", body.URL, 0, "short")
		if err !=nil {
			panic(err.Error())
		}
		defer trackurl.Close()
	}
	//response
	resp := response{
		URL:             body.URL,
		CustomShort:     "",
	}

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
