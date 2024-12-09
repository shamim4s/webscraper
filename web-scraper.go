package main

import (
	"context"      // Import the context package
	"fmt"
	"log"
	"strings"      // Import the strings package
	"github.com/chromedp/chromedp"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Set up the ChromeDP context to interact with headless Chrome
	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("headless", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-gpu", true),
	)
	defer cancel()

	// Create a new context for ChromeDP
	ctx, cancel := chromedp.NewContext(
		allocCtx,
	)
	defer cancel()

	// Define the URL to scrape
	url := "https://sites.google.com/view/bdixftpserverlist/media-ftp-servers"

	// Variable to store the page content
	var htmlContent string

	// Navigate to the page and retrieve the HTML content
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery), // Wait for the body to load
		chromedp.OuterHTML("html", &htmlContent), // Capture the HTML content
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML content using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	// Find all elements with the aria-label attribute
	doc.Find("[aria-label]").Each(func(i int, s *goquery.Selection) {
		ariaLabel := s.AttrOr("aria-label", "")
		if ariaLabel != ""  && strings.HasPrefix(ariaLabel, "http") && !strings.Contains(ariaLabel, "facebook") {

			if strings.HasPrefix(ariaLabel, "http://") {
						ariaLabel = strings.TrimPrefix(ariaLabel, "http://")
					} else if strings.HasPrefix(ariaLabel, "https://") {
						ariaLabel = strings.TrimPrefix(ariaLabel, "https://")
					}
			fmt.Println("\"" + ariaLabel + "\",")
		}
	})
}
