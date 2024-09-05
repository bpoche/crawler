package main

import (
	"fmt"
	"sort"
)

type pageInfo struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {

	// Print the report header
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", page.Count, page.URL)
	}

}

func sortPages(pages map[string]int) []pageInfo {
	var pageList []pageInfo
	for url, count := range pages {
		pageList = append(pageList, pageInfo{URL: url, Count: count})
	}

	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].Count == pageList[j].Count {
			return pageList[i].URL < pageList[j].URL
		}
		return pageList[i].Count > pageList[j].Count
	})

	return pageList
}
