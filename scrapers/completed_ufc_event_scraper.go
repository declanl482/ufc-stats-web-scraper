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

		// Skip the first two table rows. They are empty in the website HTML and will produce empty cells in the output .csv file.
		if e.Index == 0 || e.Index == 1 {
			return
		}

		// Create an empty UFC event structure.
		ufcEvent := models.UFCEvent{}

		// Scrape fields from the HTML page. UFC event data is stored in table rows on the website.

		// Parse the href attribute from the <a> element in the table row.
		ufcEvent.EventURL = e.ChildAttr("a", "href")

		// Parse the text context from the <a> element in the table row.
		ufcEvent.EventTitle = e.ChildText("a")

		// Parse the text from the "b-statistics__date" class in the table row.
		ufcEvent.EventDate = e.ChildText(".b-statistics__date")

		// Parse the text from the 2nd table data element in the table row.
		ufcEvent.EventLocation = e.ChildText("td:nth-child(2)")

		// Append the current UFC event to the slice of UFC events.
		ufcEvents = append(ufcEvents, ufcEvent)
	})

	// Create an event handler for when a page has been scraped. We will be writing the scraped data to .csv files.
	c.OnScraped(func(r *colly.Response) {

	})

}
