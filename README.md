SIU Campus Tour
===============

#REST API

##Objects
* App
* User
* Goal
* Object
* ObjectType
* Reward
* RewardType

###App
```
{
	"Id":{integer},
	"Name":{string},
	"Token":{256-bit hash},
	"Objectives":[{Objective}],
	"Users":[{User}]
}
```

###User
```
{
	"Id":{integer},
	"Fname":{string},
	"Lname":{string},
	"Birthday":{date},
	"UserID":{string},
	"Apps":[{App}]
}
```

###Objective
```
{
	"Id":{integer},
	"Name":{string},
	"Description":{string},
	"Objects":[{Object}],
	"Rewards":[{Reward}]
}
```

###Object
```
{
	"Id":{integer},
	"Name":{string},
	"Type":{ObjectType},
	"Value":{string}
}
```

###ObjectType
```
{
	"Id":{integer},
	"Name":{string}
}
```

###Reward
```
{
	"Id":{integer},
	"Name":{string},
	"Type":{RewardType},
	"Value":{string}
}
```

###RewardType
```
{
	"Id":{integer},
	"Name":{string}
}
```
