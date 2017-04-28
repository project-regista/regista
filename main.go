package main

import (
	"fmt"
	"log"

	neo4j "github.com/project-regista/neo4go/client"
	"github.com/project-regista/neo4go/storage"
	"github.com/project-regista/regista/util"
	scout "github.com/project-regista/scout/client"
	"github.com/project-regista/scout/request"
)

func main() {

	/* INITIAL STUFF */

	// Get Scout client
	scout, err := scout.New("resources/conf/.scout.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Get Neo4j DB client
	n, err := neo4j.New("resources/conf/.neo4j-local.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Get Neo4j connection
	if err := n.Connection(); err != nil {
		log.Fatal(err)
	}
	defer n.CloseConn()

	/* REQUESTS */

	// COMPETITION COUNTRY
	competitionCountry, err := request.GetCompetitionCountry(scout, "43")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("competition-country: %+v\n", competitionCountry)

	// Getting Competition Statement
	stmt, err := storage.CompetitionStmt(n.Conn)
	if err != nil {
		log.Fatal(err)
	}

	// Insert Competition in DB
	err = storage.Insert(stmt, competitionCountry)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	// COMPETITIONS COUNTRY
	competitionsCountry, err := request.GetCompetitionsCountry(scout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("competitions-countries: %+v\n", competitionsCountry)

	for _, v := range competitionsCountry.Data {
		stmt, err := storage.CompetitionStmt(n.Conn)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.Insert(stmt, v)
		if err != nil {
			log.Fatal(err)
		}
		stmt.Close()
	}

	// COUNTRY
	country, err := request.GetCountry(scout, "13")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("country: %+v\n", country)

	stmt, err = storage.CountryStmt(n.Conn)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Insert(stmt, country)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	// COUNTRIES
	countries, err := request.GetCountries(scout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("countries: %+v\n", countries)

	for _, v := range countries.Data {
		stmt, err = storage.CountryStmt(n.Conn)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.Insert(stmt, v)
		if err != nil {
			log.Fatal(err)
		}
		stmt.Close()
	}

	// SEASON COMPETITION
	seasonCompetition, err := request.GetSeasonCompetition(scout, "350")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("season competition: %+v\n", seasonCompetition)

	stmt, err = storage.SeasonStmt(n.Conn)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Insert(stmt, seasonCompetition)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()

	// SEASONS COMPETITION
	seasonsCompetition, err := request.GetSeasonsCompetition(scout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("season competition: %+v\n", seasonCompetition)

	for _, v := range seasonsCompetition.Data {
		stmt, err := storage.SeasonStmt(n.Conn)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.Insert(stmt, v)
		if err != nil {
			log.Fatal(err)
		}
		stmt.Close()
	}

	// MATCH SEASON
	matchSeason, err := request.GetMatchSeason(scout, "795379")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("match season: %+v\n", matchSeason)

	t := util.DateIntervals("01 Jan 17", "28 Jun 17")
	for k, v := range t {
		fmt.Println(k, v)
	}
}
