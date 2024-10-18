package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func Characters() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		if name != "" {

			urlPerson := fmt.Sprintf("http://heroes:8085/persons/?name=%s", name)
			respPerson, err := http.Get(urlPerson)
			if err != nil {
				log.Fatalf("Failed to get person data: %v", err)
			}
			defer respPerson.Body.Close()

			if respPerson.StatusCode != http.StatusOK {
				log.Fatalf("Error: received non-200 status code: %v", respPerson.StatusCode)
			}

			var person Person
			if err := json.NewDecoder(respPerson.Body).Decode(&person); err != nil {
				log.Fatalf("Failed to decode person data: %v", err)
			}

			urlPoints := fmt.Sprintf("http://upsell:8080/customers?id=%d", person.ID)
			respPoints, err := http.Get(urlPoints)
			if err != nil {
				log.Fatalf("Failed to get points: %v", err)
			}
			defer respPoints.Body.Close()

			if respPoints.StatusCode != http.StatusOK {
				log.Fatalf("Error: received non-200 status code: %v", respPoints.StatusCode)
			}

			var pointsData map[string]int
			bodyPoints, err := io.ReadAll(respPoints.Body)
			if err != nil {
				log.Fatalf("Failed to read points response: %v", err)
			}

			if err := json.Unmarshal(bodyPoints, &pointsData); err != nil {
				log.Fatalf("Failed to decode points data: %v", err)
			}

			person.Points = pointsData["Points"]

			text := formatPlayer(person)
			fmt.Println(text)
		}
		return nil
	}
}

func formatPlayer(p Person) string {
	return fmt.Sprintf("Name: %s\nID: %d\nPoints: %d\n", p.Name, p.ID, p.Points)
}
