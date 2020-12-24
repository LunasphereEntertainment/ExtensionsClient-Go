package RestClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type APIClient struct {
	BaseURL *url.URL
	AuthToken string
	ExtensionID int

	Client *http.Client
}

/*func (c *APIClient) SelectExtension(extensionId int) {
	jar := http.CookieJar(nil)

	var cookies []http.Cookie

	cookies = append(cookies, http.Cookie{Name: "extId", Value: strconv.Itoa(extensionId)})

	jar.SetCookies(c.BaseURL, &cookies)

	jar.Cookies(c.BaseURL)

	c.Client.Jar = jar
}*/

func (c *APIClient) Build(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse("api/" + path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	log.Printf("- %s - %s\n", method, u.String())

	var buf io.Reader
	if body != nil {
		jData, _ := json.Marshal(body)
		fmt.Println("Submitting: " + string(jData))
		buf = bytes.NewBuffer(jData)

		/*buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)*/
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.ExtensionID > 0 {
		req.AddCookie(&http.Cookie{Domain: c.BaseURL.String(), Name: "extId", Value: strconv.Itoa(c.ExtensionID)})
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("User-Agent", "HNExtensionsClient-Golang")

	return req, nil
}

func (c *APIClient) Do (req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	/*dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("- Response - ", string(dat))

	if v != nil {
		err = json.Unmarshal(dat, v)
	}*/
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}