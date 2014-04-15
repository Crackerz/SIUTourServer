package tests

import (
	"io/ioutil"
	//	"math/rand"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"
)

/*
 * Begin Standard Resources
 */

var client = &http.Client{}
var url = "http://127.0.0.1:80"

func QueryServer(method, resource, body string) (responseBody []byte, responseCode int, err error) {
	req, err := http.NewRequest(method, url+resource, strings.NewReader(body))
	if err != nil {
		return nil, 0, errors.New("Could not create request client side.\n" + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, errors.New("Failed executing request.\n" + err.Error())
	}

	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	responseCode = resp.StatusCode
	return
}

/*
func TestDistribution(t *testing.T) {
	min, max := 0, 200
	r := rand.New(rand.NewSource(20))
	for i := 0; i < 10000; i++ {
		sample := int(r.Float64())*(max-min) + min
		if sample < min || sample > max {
			t.Error(sample)
		}
	}
}
*/

/**
 * Begin APP tests
 */

func TestCreateApp(t *testing.T) {

	type AppRequest struct {
		Name string
	}

	type AppResult struct {
		Result *struct {
			Id int
		}
	}

	body, err := json.Marshal(AppRequest{"Test"})
	if err != nil {
		t.Fatal("Unable to marshal json in test:  CreateApp")
	}
	respBody, code, err := QueryServer("POST", "/app", string(body))

	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

	/**TODO: check values of returned object**/
    var respJSON AppResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result != nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

//func QueryServer(method, resource, body string) (responseBody string, responseCode int, err error)
