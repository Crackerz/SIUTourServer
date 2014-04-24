package tests

import(
    "testing"
)

func TestMapUserApp(t *testing.T) {
    userId,err := createUser("User 1","1 User","02191992","william.jblankenship@gmail.com")
    if err != nil {
        t.Fatal(err.Error())
    }

    appId,err := createApp("App 1")
    if err != nil {
        t.Fatal(err.Error())
    }

    err = mapUserApp(userId,appId)
    if err != nil {
        t.Fatal(err.Error())
    }
}

func TestMapAppObjective(t *testing.T) {
    appId,err := createApp("App 2")
    if err != nil {
        t.Fatal(err.Error())
    }

    objectiveId,err := createObjective("Objective 1","Description 1")
    if err!=nil {
        t.Fatal(err.Error())
    }

    err = mapAppObjective(appId, objectiveId)
    if err!=nil {
        t.Fatal(err.Error())
    }
}

func TestMapUserAction(t *testing.T) {
    userId,err := createUser("User 2","2 User","03171992","billyjoe@yahoo.com")
    if err!=nil {
        t.Fatal(err.Error())
    }

    actionId,err := createAction("Action 1","Random Data",0)
    if err!=nil {
        t.Fatal(err.Error())
    }

    err = mapUserAction(userId, actionId)
    if err!=nil {
        t.Fatal(err.Error())
    }
}

func TestMapActionObjective(t *testing.T) {
    actionId, err := createAction("Action 2","More Random Data",0)
    if err!=nil {
        t.Fatal(err.Error())
    }

    objectiveId, err := createObjective("Objective 2","Objective Description 2")
    if err!=nil {
        t.Fatal(err.Error())
    }

    err = mapActionObjective(actionId, objectiveId);
    if err!=nil {
        t.Fatal(err.Error())
    }
}
