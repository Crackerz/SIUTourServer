package tests

import (
	"testing"
)

/**
 * Begin APP tests
 */

func TestCreateApp(t *testing.T) {
    _,err := createApp("test");
    if(err!=nil) {
	    t.Fatal(err.Error());
    }
}

func TestCreateUser(t *testing.T) {
    	_,err:=createUser("Test", "Testy", "10201992", "SIU850955916");
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCreateAction(t *testing.T) {
	_,err:=createAction("Testing","a test action", 0)
	if err!=nil {
		t.Fatal(err.Error())
	}
}

func TestCreateObjective(t *testing.T) {
	_,err:=createObjective("Objective","Description")
	if err!=nil {
		t.Fatal(err.Error())
	}
}

func TestCreateReward(t *testing.T) {
	_,err:=createReward("Reward","Description",0)
	if err!=nil {
		t.Fatal(err.Error())
	}
}

func TestCreateActionType(t *testing.T) {
	_,err:=createActionType("ActionType","Description")
	if err!=nil {
		t.Fatal(err.Error())
	}
}

func TestCreateRewardType(t *testing.T) {
	_,err:=createRewardType("ActionType","Description")
	if err!=nil {
		t.Fatal(err.Error())
	}
}
