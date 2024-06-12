package scraper

// import required libraries
import (
	"slices"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Movie structure to represent object model
type movie struct {
	rank      int
	title     string
	year      int
	duration  int
	audience  string
	rating    float64
	votes     float64
	image_url string
	image_alt string
	movie_url string
}

// Function to parse data from https://www.imdb.com/chart/moviemeter/ (IMDB Most Popular Movies)
func Scrape() [100]movie {
	// Initialize collector, movie list, empty movie indices, counter and offset
	c, movieList, emptyMovieIndices, counter, offset := colly.NewCollector(), [100]movie{}, []int{}, 0, 0

	// Set rank for each movie (Movies are scraped sequentially therefore it is safe to assume that the rank is the index + 1)
	for i := 1; i < 101; i++ {
		movieList[i-1].rank = i
	}

	// OnHTML callback to scrape data
	c.OnHTML(".ipc-metadata-list", func(e *colly.HTMLElement) {

		// Extract image url and alt
		e.ForEach(".ipc-image", func(i int, e *colly.HTMLElement) {
			movieList[i].image_url = e.Attr("src")
			movieList[i].image_alt = e.Attr("alt")
		})

		// Extract movie url
		e.ForEach(".ipc-title-link-wrapper", func(i int, e *colly.HTMLElement) {
			movieList[i].movie_url = e.Attr("href")
		})

		// Extract movie title
		e.ForEach(".ipc-title__text", func(i int, e *colly.HTMLElement) {
			movieList[i].title = e.Text
		})

		// Extract for movie year, duration and audience then parse accordingly, results in format (yyyydurationaudience)
		e.ForEach(".cli-title-metadata", func(i int, e *colly.HTMLElement) {
			// Extract movie year
			year, err := strconv.Atoi(e.Text[0:4])

			// Handle error
			if err != nil {
				panic(err)
			} else {
				movieList[i].year = year
			}

			// Check to see if string contains 'm' (minutes) and parse accordingly
			if strings.Contains(e.Text[4:], "m") {
				// Split string at 'm'
				split := strings.Split(e.Text[4:], "m")

				// Iterate over split string, it will contain: (duration, audience)
				for j, s := range split {

					// Extract duration and convert to minutes
					if j == 0 {
						movieList[i].duration = convertDurationToMinutes(s)
					}

					// Extract audience
					if j == 1 {
						movieList[i].audience = s
					}
				}

			} else {
				// Split string at '\n' (newline)
				split := strings.Split(e.Text[4:], "\n")

				// Extract audience
				movieList[i].audience = split[0]
			}
		})

		// Extract movie rating
		e.ForEach(".ipc-rating-star", func(i int, e *colly.HTMLElement) {
			// Extract aria-label attribute
			label := e.Attr("aria-label")

			// Check to see if the label contains 'This title is currently not ratable', if so set rating to 0 and add index to empty movie indices
			if strings.Contains(label, "This title is currently not ratable") {
				movieList[counter].rating = 0
				emptyMovieIndices = append(emptyMovieIndices, counter)
				counter++
			} else {
				// Extract rating from label (format 'rating: x.x/10')
				ratingLength := len(label) - 3

				// Check to see if rating length is greater than 0
				if ratingLength > 0 {
					// Parse rating and convert to float64
					rating, err := strconv.ParseFloat(label[ratingLength:], 64)

					// Handle error
					if err != nil {
						panic(err)
					} else {
						// Set rating and increment counter
						movieList[counter].rating = rating
						counter++
					}
				}
			}
		})

		// Extract movie votes
		e.ForEach(".ipc-rating-star--voteCount", func(i int, e *colly.HTMLElement) {
			// Check to see if index is in empty movie indices, if so add empty votes count and increment offset
			if slices.Contains(emptyMovieIndices, i) {
				addEmptyVotesCount(i+offset, movieList[:])
				offset += 1
			}

			// Check to see if string contains 'K' or 'M' (thousands or millions) and parse accordingly
			if strings.Contains(e.Text, "K") {
				// Split string at 'K'
				split := strings.Split(e.Text, "K")

				// Convert votes to float64
				votes, err := strconv.ParseFloat(split[0][3:], 64)

				// Handle error
				if err != nil {
					panic(err)
				} else {
					// Set votes
					movieList[i+offset].votes = votes * 1000
				}
			} else if strings.Contains(e.Text, "M") {
				// Split string at 'M'
				split := strings.Split(e.Text, "M")

				// Convert votes to float64
				votes, err := strconv.ParseFloat(split[0][3:], 64)

				// Handle error
				if err != nil {
					panic(err)
				} else {
					// Set votes
					movieList[i+offset].votes = votes * 1000000
				}
			}
		})
	})

	// Visit URL
	c.Visit("https://www.imdb.com/chart/moviemeter/")

	// Return movie list
	return movieList
}

// Helper function to convert duration from scraped data to minutes
func convertDurationToMinutes(duration string) int {
	// Split duration string at 'h '
	minutes, split := 0, strings.Split(duration, "h ")

	// Check to see if split is empty
	if len(split) == 0 {
		// Extract minutes and convert to int
		m, err := strconv.Atoi(duration[0:1])

		// Handle error
		if err != nil {
			panic(err)
		} else {
			return m * 60
		}
	} else {
		// Iterate over split string, it will contain: (hours, minutes)
		for i, s := range split {
			// Do the following for hours (0) and minutes (1)
			if i == 0 {
				// Convert hours to int
				hours, err := strconv.Atoi(s)

				// Handle error
				if err == nil {
					minutes += hours * 60
				} else if len(s) == 0 {
					minutes += 0
				} else {
					tmp := strings.Split(s, "h")
					hours, err = strconv.Atoi(tmp[0])

					if err == nil {
						minutes += hours * 60
					} else {
						panic(err)
					}
				}
			} else if i == 1 {
				// Convert minutes to int
				min, err := strconv.Atoi(s)

				// Handle error
				if err != nil {
					panic(err)
				} else {
					minutes += min
				}
			}
		}
	}

	// Return duration in minutes
	return minutes
}

// Helper function to add empty votes count to movie list
func addEmptyVotesCount(index int, movieList []movie) bool {
	// Check to see if index is out of bounds
	if (index < 0) || (index >= len(movieList)) {
		// Return false
		return false
	} else {
		// Set votes to 0 and return true
		movieList[index].votes = 0.0
		return true
	}
}
