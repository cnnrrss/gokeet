package feeder

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"sync"
)

// feedCmd represents the feed command
var FeedCmd = &cobra.Command{
	Use:   "feed",
	Short: "Read recently published articles from subscribed RSS feeds",
	Run: run,
}

const nWorkers = 5

var wg sync.WaitGroup

func load(feed *FeedConfig) {
	fp := gofeed.NewParser()
	parsedFeed, err := fp.ParseURL(feed.Host)
	if err != nil {
		fmt.Printf("%s | Error: %v", feed.Host, err)
		return
	}

	fmt.Printf("%s | %s | %s\n", parsedFeed.Title, feed.Host, parsedFeed.FeedType)

	for _, page := range parsedFeed.Items {
		fmt.Printf("%s | %s\n", page.Title, page.Link)
	}
	fmt.Println()
}

func worker(n int, sites <-chan FeedConfig) {
	//fmt.Printf("started worker %d\n", n)

	for work := range sites {
		load(&work)
	}

	//fmt.Printf("finished worker %d\n", n)

	wg.Done()
}

func startWorkers(n int) chan FeedConfig {
	wg.Add(n)
	ch := make(chan FeedConfig, n)
	for id := 0; id < n; id++ {
		go worker(id, ch)
	}
	return ch
}

func run(_ *cobra.Command, _ []string) {
	fmt.Println("feeder started.")

	sites := []string{
		"http://feeds.feedburner.com/AmazonWebServicesBlog",
		"https://hnrss.org/best",
		"https://hoopshype.com/feed/",
	}

	c := startWorkers(nWorkers)
	for _, site := range sites {
		c <- FeedConfig{ Host: site }
	}

	close(c)
	wg.Wait()

	fmt.Println("feeder finished.")
}
