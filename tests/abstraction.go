package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"io/ioutil"
	"strconv"
)

var client = &http.Client{}
var url = "http://127.0.0.1:80"

func QueryServer(method,resource,body string) (responseBody []byte, responseCode int, err error) {
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

func createApp(name string) (assetID int, err error) {
	type AppRequest struct {
		Name string
	}

	type AppResult struct {
		Result *struct {
			Id int
		}
	}

    var input = AppRequest{name}
    var output AppResult
    respBody,err := makeRequest(input,&output, "POST", "/app")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createUser(fname, lname, birthday, userID string) (assetID int, err error) {
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

    var input = UserRequest{fname,lname,birthday,userID}
    var output UserResult
    respBody,err := makeRequest(input,&output, "POST", "/user")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createAction(name, description string, actionType int) (assetID int, err error) {
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

    var input = ActionRequest{name,actionType,description}
    var output ActionResult
    respBody,err := makeRequest(input,&output, "POST", "/action")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createObjective(name,description string) (assetID int, err error) {
    type ObjectiveRequest struct {
        Name string
        Description string
    }

	type ObjectiveResult struct {
		Result *struct {
			Id int
		}
    }

    var input = ObjectiveRequest{name,description}
    var output ObjectiveResult
    respBody,err := makeRequest(input,&output, "POST", "/objective")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createReward(name, value string, rewardType int) (assetId int, err error) {
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

    var input = RewardRequest{name,rewardType,value}
    var output RewardResult
    respBody,err := makeRequest(input,&output, "POST", "/reward")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createActionType(name, description string) (assetId int, err error) {
    type ActionTypeRequest struct {
        Name string
        Description string
    }

	type ActionTypeResult struct {
		Result *struct {
			Id int
		}
    }

    var input = ActionTypeRequest{name,description}
    var output ActionTypeResult
    respBody,err := makeRequest(input,&output, "POST", "/type/action")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return output.Result.Id, nil;
}

func createRewardType(name, description string) (assetId int, err error) {
    type RewardTypeRequest struct {
        Name string
        Description string
    }

	type RewardTypeResult struct {
		Result *struct {
			Id int
		}
    }

    var input = RewardTypeRequest{name,description}
    var output RewardTypeResult
    respBody,err := makeRequest(input,&output, "POST", "/type/reward")

    if err != nil {
        return 0, errors.New("Message Body: "+string(respBody)+"\n"+err.Error());
    }
    if output.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

	return output.Result.Id, nil;
}

func makeRequest(input interface{}, output interface{}, method, url string) ([]byte, error) {
    body,err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("Could not marshal JSON in test: CreateActionType.\n"+ err.Error())
	}

	respBody, code, err := QueryServer(method, url, string(body))
	if err != nil {
		return respBody, errors.New("Could not make request.\n"+ err.Error())
	}
	if code != 201 {
		return respBody, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    err = json.Unmarshal(respBody,&output)

    if err != nil {
        return respBody, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    }
    return respBody, nil;
}
