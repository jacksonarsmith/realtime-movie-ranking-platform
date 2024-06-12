package scraper

import (
	"fmt"
	"testing"
)

// TestScrape tests the Scrape function
func TestScrape(t *testing.T) {
	t.Run("scrape", func(t *testing.T) {
		res := Scrape()
		fmt.Println(res)
		for i := 0; i < len(res); i++ {
			if res[i].title == "" {
				t.Errorf("Scrape() = %v; want %v", res[i].title, "not empty")
			}
		}
	})
}

// TestConvertDurationToMinutes tests the convertDurationToMinutes function
func TestConvertDurationToMinutes(t *testing.T) {
	testCases := []struct {
		name     string
		duration string
		want     int
	}{
		{"1h 30", "1h 30", 90},
		{"2h", "2h", 120},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := convertDurationToMinutes(tc.duration)
			if got != tc.want {
				t.Errorf("convertDurationToMinutes(%v) = %v; want %v", tc.duration, got, tc.want)
			}
		})
	}
}

// TestAddEmptyVotesCount tests the addEmptyVotesCount function
func TestAddEmptyVotesCount(t *testing.T) {
	testArr := [2]movie{}

	testArr[0].rank = 1
	testArr[0].title = "Test Movie 1"
	testArr[0].year = 2021
	testArr[0].duration = 90
	testArr[0].audience = "PG-13"
	testArr[0].rating = 7.5
	testArr[0].votes = 0
	testArr[0].image_url = "https://test.com/image1.jpg"
	testArr[0].image_alt = "Test Movie 1"
	testArr[0].movie_url = "https://test.com/movie1"

	t.Run("AddEmptyVotesCount", func(t *testing.T) {
		res := addEmptyVotesCount(1, testArr[:])
		if !res {
			t.Errorf("addEmptyVotesCount() = %v; want %v", res, true)
		}
	})

}
