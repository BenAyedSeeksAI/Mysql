package db

import (
	"database/sql"
	"fmt"
	"mysqlCountryProjects/schema"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB {
	password := "pourlaforme"
	config := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/world", password)
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func GetCountries(limit int) ([]string, error) {
	db := openDB()
	defer db.Close()

	query := `SELECT Name 
	FROM country
	LIMIT ?`

	var countries []string

	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var countryName string
		if err := rows.Scan(&countryName); err != nil {
			return countries, err
		}
		countries = append(countries, countryName)
	}
	if err = rows.Err(); err != nil {
		return countries, err
	}
	return countries, nil
}
func GetGovernmentsPerCountryStats() ([]schema.GovernmentForm, error) {
	db := openDB()
	defer db.Close()

	query := `SELECT country.GovernmentForm, count(country.GovernmentForm) as GovCount
from country
group by country.GovernmentForm
order by GovCount DESC`
	StatsList := make([]schema.GovernmentForm, 0)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stats schema.GovernmentForm
		if err := rows.Scan(&stats.GovernmentType, &stats.NumberOfCountries); err != nil {
			return StatsList, err
		}
		StatsList = append(StatsList, stats)
	}
	return StatsList, nil
}
func GetCountryPerContinentStats() ([]schema.CountCountryByContinent, error) {
	db := openDB()
	defer db.Close()

	query := `SELECT country.Continent, count(country.Continent)as ConCount
from country
group by country.Continent
order by ConCount DESC`
	StatsList := make([]schema.CountCountryByContinent, 0)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stats schema.CountCountryByContinent
		if err := rows.Scan(&stats.Continent, &stats.NumberOfCountries); err != nil {
			return StatsList, err
		}
		StatsList = append(StatsList, stats)
	}
	return StatsList, nil
}
