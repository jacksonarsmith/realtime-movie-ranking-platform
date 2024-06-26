package scraper

// import required libraries
import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// Movie structure to represent object model
type Movie struct {
	Title    string
	Rank     int32
	Year     int32
	Duration int32
	Audience string
	Rating   float64
	ImageSrc string
	ImageAlt string
	MovieUrl string
}

// Function to parse data from https://www.imdb.com/chart/Moviemeter/ (IMDB Most Popular Movies)
func Scrape() []Movie {
	// Initialize empty movie slice
	var movies []Movie

	// Log starting data fetching
	log.Print("Starting data fetching...\n")

	// Initialize chrome options
	options := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36"),
		chromedp.Headless,
	}

	// Initialize chrome instance with options
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// Create new context with log
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Initialize nodes slice
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.imdb.com/chart/moviemeter"),
		chromedp.Nodes(".ipc-metadata-list-summary-item", &nodes, chromedp.ByQueryAll),
	)

	// Handle error
	if err != nil {
		log.Fatalf("Failed to execute chromedp tasks: %v", err)
	}

	// Initialize movie data variables
	var title, ranking, image_src, image_alt, movie_url, span_text, audience, rating string
	var year, duration int32

	// Iterate over nodes
	for _, node := range nodes {
		/*
			Extract data from node from the following selectors: title, ranking, image_src,
			image_alt, movie_url, span_text, rating to satisfy movie structure above
		*/
		err := chromedp.Run(ctx,
			chromedp.Text("h3", &title, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".meter-const-ranking", &ranking, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("img", "src", &image_src, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("img", "alt", &image_alt, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("a", "href", &movie_url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".cli-title-metadata", &span_text, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue(".ipc-rating-star", "aria-label", &rating, nil, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		// Handle error
		if err != nil {
			log.Printf("Failed to extract data: %v", err)
			continue // Skip to next movie
		}

		// Do string manipulation to extract ranking
		ranking = strings.Split(ranking, " ")[0] // First element of split string is always the ranking

		// Convert string ranking to integer
		rank, err := strconv.Atoi(ranking)

		// Handle error
		if err != nil {
			panic(err)
		}

		// Extract year, duration and audience from span text calling the helper function extractMovieSpan
		year, duration, audience = extractMovieSpan(span_text)

		// Log movie data extracted from node
		log.Printf("Title: %s, Image src: %s, Image alt: %s, Movie link: %s, Year: %d, Duration: %d, audience: %s, Rating: %f\n", title, image_src, image_alt, movie_url, year, duration, audience, extractRating(rating))

		// Create movie object and append to movies slice
		movie := Movie{Title: title, Rank: int32(rank), ImageSrc: image_src, ImageAlt: image_alt, MovieUrl: movie_url, Year: year, Duration: duration, Audience: audience, Rating: extractRating(rating)}
		movies = append(movies, movie)
	}

	// Log data fetching complete
	log.Println("Data fetching complete")

	// Return movies slice
	return movies
}

// Helper function to extract year, duration and audience from span text
func extractMovieSpan(span_text string) (int32, int32, string) {
	// Split span text at '\n'
	str := strings.Split(span_text, "\n")

	// Initialize variables
	var year, duration int32
	var audience string

	// Iterate over split string
	for i, s := range str {
		// Do the following for year (0), duration (1) and audience (2)
		switch {
		case i == 0:
			// Convert year to int
			y, err := strconv.Atoi(s)

			// Handle error
			if err != nil {
				panic(err)
			}

			year = int32(y)
		case i == 1:
			// Check to see if string contains 'h' or 'm', if so convert duration to minutes
			if strings.Contains(s, "h") || strings.Contains(s, "m") {
				duration = convertDurationToMinutes(s)
			} else {
				duration = 0
			}
		case i == 2:
			// Set audience
			audience = s
		default:
			// Set year, duration and audience to 0 and empty string
			year = 0
			duration = 0
			audience = ""
		}
	}

	// Return year, duration and audience
	return year, duration, audience
}

// Helper function to extract rating from rating string
func extractRating(rating string) float64 {
	// Check to see if the rating contains 'This title is currently not ratable', if so set rating to 0 and add index to empty movie indices
	if strings.Contains(rating, "This title is currently not ratable") {
		return 0.0
	} else {
		// Extract rating from rating (format 'rating: x.x/10')
		ratingLength := len(rating) - 3

		// Check to see if rating length is greater than 0
		if ratingLength > 0 {
			// Parse rating and convert to float64
			ratingFloat, err := strconv.ParseFloat(rating[ratingLength:], 64)

			// Handle error
			if err != nil {
				panic(err)
			} else {
				// Set rating and increment counter
				return ratingFloat
			}
		} else {
			return 0.0
		}
	}
}

// Helper function to convert duration string to minutes
func convertDurationToMinutes(duration string) int32 {
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
			return int32(m) * 60
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
				min, err := strconv.Atoi(s[:len(s)-1])

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
	return int32(minutes)
}
