package pingfed

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

func TestExtractMfaFormData(t *testing.T) {
	data, err := ioutil.ReadFile("example/mfapage.html")
	require.Nil(t, err)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	require.Nil(t, err)

	mfaForm, actionURL, err := extractMfaFormData(doc, "#form1")
	require.Nil(t, err)
	require.Equal(t, "https://authenticator.pingone.com/pingid/ppm/auth/poll", actionURL)
	require.Equal(t, url.Values{"csrfToken": []string{"fc80998c-34d8-4dd2-925c-3b3be8a0dee8"}}, mfaForm)
}

var extractAuthSubmitURLTests = []struct {
        f        string // input html file
        expected string // expected url
}{
	{"example/loginpage.html", "https://example.com/relative/login"},
	{"example/loginpage_absolute.html", "https://other.example.com/login"},
}

func TestExtractAuthSubmitURL(t *testing.T) {
	for _, tt := range extractAuthSubmitURLTests {
		data, err := ioutil.ReadFile(tt.f)
		require.Nil(t, err)

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
		require.Nil(t, err)

		url, err := extractAuthSubmitURL("https://example.com", doc)
		require.Nil(t, err)
		require.Equal(t, tt.expected, url)
	}
}
