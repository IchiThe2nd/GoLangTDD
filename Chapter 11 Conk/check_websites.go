package concurrency

//websitecheker is a func checks a url retunring a bool
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

//Check Websites takes a websitechecker and a slice of urls and returns a map of urlss
// to the result of checkeing each cwebsite with the website checker

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// ignor keys and range through urls passing strings to "u"
	for _, url := range urls {
		go func(u string) {
			// send to resultCHannel // result has string & bool make a result of u and boolean return
			// of website checker
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		//receive from resulty channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
