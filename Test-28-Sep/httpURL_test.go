package main
import (
    "reflect"
    "testing"
)
func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}
func TestCheckWebsites(t *testing.T) {
    websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }
    want := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }
    got := CheckWebsites(mockWebsiteChecker, websites)
    if !reflect.DeepEqual(want, got) {
        t.Fatalf("wanted %v, got %v", want, got)
    }
}
// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool
// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    for _, url := range urls {
        results[url] = wc(url)
    }
    return results
}