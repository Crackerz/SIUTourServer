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

###Error
{
	"Message":{string}
}

##REST Access Points

`/app`

Supports:
* `PUT`
* `GET`

`PUT: /app`

PUT registers a new application with the server by name. The new name must be unique or an error will be generated.

Input:
```
{
	"Name":{string}
}
```
Output:
```
{
	"Token":{256-bit hash}
}
```

`GET /app?token={token}`
This query returns an array of apps. The token parameter is optional. If the token parameter is provided and matches the token of a registered app, this call will return an array of length one (1) whose sole member is that app. If no app has that registered token, an empty array will be returned.

Input:
```
```

Output:
```
{
	"Apps":[{
				"Id":{integer},
				"Name":{string},
				"Token":{256-bit hash},
				"Objectives":[{Objective}],
				"Users":[{User}]
			}]
}
```
