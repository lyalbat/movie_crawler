package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Movie struct {
	Name    string
	Year    uint32
	InLists []uint32
}

func main() {
	baseUrl := "https://letterboxd.com/"

	maxPages := 10

	for i := 0; i < maxPages; i++ {

		scrapeUrl := baseUrl + "films/popular/page/" + fmt.Sprint(i)
		getMovies(scrapeUrl)

	}

}

// build a database of movies
func getMovies(url string) {

	c := colly.NewCollector(colly.AllowedDomains("www.letterboxd.com", "letterboxd.com"))

	c.OnHTML("h1.title-1.prettify", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scraping: %s\n", err.Error())
	})

	c.Visit(url)
}

func convertToMovie(rawMovie string) {
	movie := Movie{}
	movieInfo := strings.Split(rawMovie, "(")
	fmt.Printf("movieInfo: %q\n", movieInfo[1])
	name := strings.TrimSuffix(movieInfo[0], " ")
	yearAsString := strings.TrimSuffix(movieInfo[1], ")")
	fmt.Println(yearAsString)
	year, second := strconv.ParseInt(yearAsString, 10, 64)
	fmt.Printf("year: %q\n", year)
	fmt.Printf("second info: %q\n", second)
	movie.Name = name
	//movie.Year = year
	fmt.Printf("movie: %q\n", movie)
	//fmt.Printf("movie: %p\n", movie)
	//movie.Year = mov
}
