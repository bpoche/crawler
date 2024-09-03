package get_urls_from_html

//"errors"

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(rawBaseURL, htmlBody string) ([]string, error) {
	// return all url links found in the html body
	urls := []string{}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	//search for url links in html body
	tokenizer := html.NewTokenizer(strings.NewReader(htmlBody))
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			}
			return nil, tokenizer.Err()
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						//parse the url
						linkURL, err := url.Parse(attr.Val)
						if err != nil {
							return nil, err
						}
						//check if the url is relative or absolute
						if linkURL.IsAbs() {
							urls = append(urls, linkURL.String())
						} else {
							//make the url absolute
							linkURL = baseURL.ResolveReference(linkURL)
							urls = append(urls, linkURL.String())
						}
					}
				}
			}
		}
	}

	return urls, nil
}
