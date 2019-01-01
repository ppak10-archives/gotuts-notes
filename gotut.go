package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

/*
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/local.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/sports.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/national.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/world.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/business.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/technology.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/lifestyle.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/entertainment.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
</loc>
</sitemap>
</sitemapindex>
 */
type SitemapIndex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

type NewsMap struct {
  Keyword string
  Location string
}

func main () {
  var s SitemapIndex
  var n News
  news_map := make(map[string]NewsMap)
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)

  for _, Location := range s.Locations {
    resp, _ := http.Get(Location)
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
    for idx, _ := range n.Titles {
      news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
    }
  }

  for idx, data := range news_map {
    fmt.Println("\n\n\n",idx)
    fmt.Println("\n", data.Keyword)
    fmt.Println("\n", data.Location)
  }
}
