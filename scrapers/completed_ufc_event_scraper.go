package scrapers

import (
	"encoding/csv"
	"log"
	"os"
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

		// Create a .csv file for the scraped UFC events in the data directory.
		file, err := os.Create("data/completed_ufc_events.csv")

		// Handle any errors in generating the .csv file.
		if err != nil {
			log.Fatalln("Failed to create completed_ufc_events.csv file: ", err)
		}

		// Ensure that the .csv file is closed when the surrounding function returns or errors out.
		defer file.Close()

		// Create a CSV writer object to write data to the .csv file.
		writer := csv.NewWriter(file)

		// Generate and write the column headers for the output file.
		headers := []string{
			"event_url",
			"event_title",
			"event_date",
			"event_location",
		}
		writer.Write(headers)

		// Write the scraped UFC event data to the output file.
		// The range keyword in Go allows us to iterate over structures like slices. It returns the index, element for each element in the slice.
		// We can ignore the index here since it is not relevant to the data we are looking to extract.
		for _, ufcEvent := range ufcEvents {
			record := []string{
				ufcEvent.EventURL,
				ufcEvent.EventTitle,
				ufcEvent.EventDate,
				ufcEvent.EventLocation,
			}
			writer.Write(record)
		}

		// Ensure that all data from the writer object is flushed to the output file when the surrounding function returns or errors out.
		defer writer.Flush()
	})

	// The Visit function starts the life cycle of Colly.
	c.Visit("http://www.ufcstats.com/statistics/events/completed?page=all")

}
