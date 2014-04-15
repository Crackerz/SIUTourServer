package tests

import (
	"io/ioutil"
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
var url = "http://127.0.0.1:8081"

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

    var respJSON AppResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

func TestCreateUser(t *testing.T) {

    type UserRequest struct {
        Fname string
        Lname string
        Birthday string
        UserID string
    }

	type UserResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(UserRequest{"Testing","Tester","10201992","SIU850955916"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/user", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON UserResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

func TestCreateAction(t *testing.T) {

    type ActionRequest struct {
        Name string
        Type int
        Description string
    }

	type ActionResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(ActionRequest{"Testing",0,"A test action"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/action", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON ActionResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

func TestCreateObjective(t *testing.T) {

    type ObjectiveRequest struct {
        Name string
        Description string
    }

	type ObjectiveResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(ObjectiveRequest{"Testing","A test objective"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/objective", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON ObjectiveResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

func TestCreateReward(t *testing.T) {
    type RewardRequest struct {
        Name string
        Type int
        Value string
    }

	type RewardResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(RewardRequest{"Testing",0,"A test objective"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/reward", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON RewardResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }
}

func TestCreateActionType(t *testing.T) {
    type ActionTypeRequest struct {
        Name string
        Description string
    }

	type ActionTypeResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(ActionTypeRequest{"Testing","A test objective"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/type/action", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON ActionTypeResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }

}

func TestCreateRewardType(t *testing.T) {
    type RewardTypeRequest struct {
        Name string
        Description string
    }

	type RewardTypeResult struct {
		Result *struct {
			Id int
		}
    }

    body,err := json.Marshal(RewardTypeRequest{"Testing","A test objective"})
	if err != nil {
		t.Error("Could not make request.\n", err.Error())
	}

	respBody, code, err := QueryServer("POST", "/type/reward", string(body))
	if code != 200 {
		t.Error("Received incorrect error code. Expected 200 and recieved ", code)
	}

    var respJSON RewardTypeResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        t.Error(err.Error(),"\nUnable to parse response as JSON.\nMessage Body:\n",string(respBody))
    } else if respJSON.Result == nil {
        t.Error("No Result object in response.\nMessage Body:\n",string(respBody))
    }

}
