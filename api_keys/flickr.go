package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Function to test Flickr access token
func testFlickrAccessToken(accessToken string) {
	oauth_token := "72157685534112345-1a2b3c4d5e6f7g8h"
	apiURL := "https://www.flickr.com/services/rest/"
	params := url.Values{}
	params.Add("method", "flickr.test.login")
	params.Add("format", "json")
	params.Add("nojsoncallback", "1")
	params.Add("oauth_token", oauth_token)

	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("Error testing access token '%s': %v\n", accessToken, err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response body: %s\n", body)

	// You would typically parse the JSON response here and check if the token is valid
	// Example:
	// if isValid(response) {
	//     fmt.Printf("Access token '%s' is valid!\n", accessToken)
	// } else {
	//     fmt.Printf("Access token '%s' is invalid.\n", accessToken)
	// }
}