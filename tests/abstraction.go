package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
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

	body, err := json.Marshal(AppRequest{name})
	if err != nil {
		return 0, errors.New("Unable to marshal json in test:  CreateApp\n"+err.Error())
	}

	respBody, code, err := QueryServer("POST", "/app", string(body))

	if err != nil {
		return 0, errors.New("Could not make request.\n"+err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

	var respJSON AppResult
	err = json.Unmarshal(respBody,&respJSON)

	if err != nil {
		return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
	} else if respJSON.Result == nil {
		return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
	}

	return respJSON.Result.Id, nil
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

    body,err := json.Marshal(UserRequest{fname,lname,birthday,userID})
	if err != nil {
		return 0, errors.New("Could not marshal json in Test: createUser.\n"+err.Error())
	}

	respBody, code, err := QueryServer("POST", "/user", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON UserResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0, errors.New(err.Error()+"No Result object in response.\nMessage Body:\n"+string(respBody))
    }

    return respJSON.Result.Id, nil
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

    body,err := json.Marshal(ActionRequest{name,actionType,description})
	if err != nil {
		return 0, errors.New("Could not marshal JSON in test: CreateAction.\n"+err.Error())
	}

	respBody, code, err := QueryServer("POST", "/action", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON ActionResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

    return respJSON.Result.Id, nil
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

    body,err := json.Marshal(ObjectiveRequest{name,description})
	if err != nil {
		return 0, errors.New("Could not marshal JSON in test: CreateObjective.\n"+err.Error())
	}

	respBody, code, err := QueryServer("POST", "/objective", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON ObjectiveResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

    return respJSON.Result.Id, nil
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

    body,err := json.Marshal(RewardRequest{name,rewardType,value})
	if err != nil {
		return 0, errors.New("Could not marshal JSON in test: CreateReward.\n"+err.Error())
	}

	respBody, code, err := QueryServer("POST", "/reward", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON RewardResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0 , errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

    return respJSON.Result.Id, nil;
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

    body,err := json.Marshal(ActionTypeRequest{name,description})
	if err != nil {
		return 0, errors.New("Could not marshal JSON in test: CreateActionType.\n"+ err.Error())
	}

	respBody, code, err := QueryServer("POST", "/type/action", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+ err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON ActionTypeResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

	return respJSON.Result.Id, nil;
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

    body,err := json.Marshal(RewardTypeRequest{name,description})
	if err != nil {
		return 0, errors.New("Could not marshal JSON in test: CreateActionType.\n"+ err.Error())
	}

	respBody, code, err := QueryServer("POST", "/type/action", string(body))
	if err != nil {
		return 0, errors.New("Could not make request.\n"+ err.Error())
	}
	if code != 201 {
		return 0, errors.New("Received incorrect error code. Expected 201 and recieved "+strconv.Itoa(code))
	}

    var respJSON RewardTypeResult
    err = json.Unmarshal(respBody,&respJSON)

    if err != nil {
        return 0, errors.New(err.Error()+"\nUnable to parse response as JSON.\nMessage Body:\n"+string(respBody))
    } else if respJSON.Result == nil {
        return 0, errors.New("No Result object in response.\nMessage Body:\n"+string(respBody))
    }

	return respJSON.Result.Id, nil;
}
func main() {
	id,err:=createApp("Testing");
	if err!=nil {
		fmt.Println(err.Error());
	}
	fmt.Println(strconv.Itoa(id));
}
