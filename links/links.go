package links

import "regexp"

var Url = regexp.MustCompile(`(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`)

func ExtractLinks(text string) []string {
	links := []string{}
	for _, match := range Url.FindAllString(text, -1) {
		links = append(links, match)
	}
	return links
}
