package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"

	"github.com/temathc/news-aggregator/pkg/config"
	"github.com/temathc/news-aggregator/pkg/database"
	"github.com/temathc/news-aggregator/pkg/scanner"
)

var conn string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	conn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}

func main() {
	fmt.Println("Ticker started...")
	base := database.NewConnect(conn)
	defer base.Close()
	repoDB := database.NewRepoDB(base)

	config, err := config.GetConf()
	if err != nil {
		fmt.Println("Failed to get configuration from conf.json")
		log.Fatal(err)
		return
	}

	parser := gofeed.NewParser()
	parser.UserAgent = "MyTestAgent"
	scanner := scanner.NewScanner(repoDB, parser)

	ticker := time.NewTicker(time.Duration(config.Timer) * time.Minute)
	for ; ; <-ticker.C {
		wg := &sync.WaitGroup{}
		for i := 0; i < len(config.Links); i++ {
			wg.Add(1)
			go func(i int) {
				err := scanner.ScanRss(config.Links[i], wg)
				if err != nil {
					fmt.Println(err)
				}
			}(i)
		}
		wg.Wait()
	}
}
