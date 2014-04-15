package tests

import (
	"io/ioutil"
//	"math/rand"
	"net/http"
	"testing"
	"strings"
	"errors"
	"encoding/json"
)

/*
 * Begin Standard Resources
 */

var client = &http.Client{}
var url = "http://127.0.0.1:80"

func QueryServer(method, resource, body string) (responseBody string, responseCode int, err error) {
	req,err := http.NewRequest(method,url+resource,strings.NewReader(body))
	if err != nil {
		return "", 0, errors.New("Could not create request client side.\n"+err.Error())
	}

	resp,err := client.Do(req)
	if err != nil {
		return "", 0, errors.New("Failed executing request.\n"+err.Error())
	}

	defer resp.Body.Close()
	rbody,err := ioutil.ReadAll(resp.Body)
	responseBody = string(rbody)
	responseCode = resp.StatusCode
	return
}

func TestApp(t *testing.T) {
	resp,code,err:=QueryServer("GET","/test","");
	if err != nil {
		t.Error("Error contactin server\n", err)
	} else if code != 500 {
		t.Error("Received Status: ", code, "\nMessage Body: ", resp)
	}
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
		Result struct{
				Id int
		}
	}

	body,err := json.Marshal(AppRequest{"Test"});
	if err!=nil {
		t.Fatal("Unable to marshal json in test:  CreateApp");
	}
	_, code, err := QueryServer("POST","/app",string(body))

	if code!=200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ",code);
	}

	/**TODO: check values of returned object**/
}

//func QueryServer(method, resource, body string) (responseBody string, responseCode int, err error)
