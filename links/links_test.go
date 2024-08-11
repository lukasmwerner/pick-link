package links

import "testing"

func TestExtractLinks(t *testing.T) {
	text := `this is google.com https://google.com and http://a-web-link.com
	and this is a super long link https://abc.com/ljasdfljljasldfjllkjkjasldfjasdfjasdflkjasdfuiqweurqweroiqwerouqweriuobnvbvbncbnbcnvnbdfuii9uqweriuqwerui`
	links := ExtractLinks(text)
	expected_links := []string{"https://google.com", "http://a-web-link.com", "https://abc.com/ljasdfljljasldfjllkjkjasldfjasdfjasdflkjasdfuiqweurqweroiqwerouqweriuobnvbvbncbnbcnvnbdfuii9uqweriuqwerui"}
	if len(links) != len(expected_links) {
		t.Errorf("different lengths %v != %v", len(links), len(expected_links))
		return
	}
	for i, link := range links {
		if link != expected_links[i] {
			t.Errorf("links did not match %s != %s", link, expected_links[i])
		}
	}
}
