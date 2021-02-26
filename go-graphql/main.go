package main

import (
  "go-grapqhql/feeds"
  "go-grapqhql/itunes"
  "log"
  "strconv"
)

func main() {
  ias := itunes.NewItunesApiServices()
  
  res, err := ias.Search("Full stack Radio")
  
  if err != nil {
    log.Fatalf("error on itunes podcast search: %v", err)
  }
  
  
  
  for _, item := range res.Results {
    log.Println("------------------------")
    log.Printf("Podcaster Name: %s", item.ArtistName)
    log.Printf("Podcast Name: %s", item.TrackName)
    log.Printf("Feed Url: %s", item.FeedURL)
    
    feed, err := feeds.GetFeed(item.FeedURL)
    
    if err != nil {
      log.Fatalf("error on itunes feed search: %v", err)
    }
    
    for _, pod := range feed.Channel.Item {
      duration, _ := strconv.Atoi(pod.Duration)
      log.Println("------------------------")
      log.Printf("Title: %s", pod.Title)
      log.Printf("Duration: %d minutes", duration/60)
      log.Printf("Description: %s", pod.Description)
      log.Printf("Url: %s", pod.Enclosure.URL)
      log.Println("------------------------")
    }
    log.Println("------------------------")
  }
}
