package scrapers

import (
	"ufc-stats-web-scraper/models"

	"github.com/gocolly/colly"
)

func ScrapeCompletedEvents() {

	// Generate a slice of UFC event structures to scrape event details into.
	var ufcEvents []models.UFCEvent

	// Create a Colly Collector instance. This is the main entity of the Colly library, which provides an interface for managing network communications
	// and executing callback functions attached to the collector, while collector jobs are running.
	c := colly.NewCollector()

	// Create an event handler for scraping UFC event data from table rows.
	c.OnHTML("tr.b-statistics__table-row", func(e *colly.HTMLElement) {

	})

	// Create an event handler for when a page has been scraped. We will be writing the scraped data to .csv files.
	c.OnScraped(func(r *colly.Response) {

	})

}
