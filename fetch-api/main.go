package main

import (
	"fetch-api/nhlApi"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Error on opening the file rosters.txt: %v", err)
	}

	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()

	if err != nil {
		log.Fatalf("error while fetching teams %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(teams))

	// unbuffered channel
	results := make(chan []nhlApi.Roster)

	for _, team := range teams {
		go func(team nhlApi.Team) {
			roster, err := nhlApi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("error getting roster: %v", err)
			}

			results <- roster

			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v", time.Now().Sub(now))
}

func display(results chan []nhlApi.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("----------------")
			log.Printf("Name %s\n", ros.Person.FullName)
			log.Printf("Position %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey Number %s\n", ros.JerseyNumber)
			log.Println("----------------")
		}
	}

}
