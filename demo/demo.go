package main

import (
	"github.com/botsphp/crontab"
	"log"
	"sort"
)

func main() {
	cron := []string{
		`1 1 1 * *`,
		`* * * * *`,
		`*/10 */6 * * *`,
	}
	for _, r := range cron {
		parse(r)
	}
}

func parse(c string) {
	entry, err := crontab.ParseEntry(c)
	if err != nil {
		log.Fatalf("Error when parsing line %v: %s", entry, err)
	}
	s := entry.Schedule.Seconds()
	sort.Ints(s)
	log.Printf("len: %d\n", len(s))
	log.Printf("%#v", s)
}
