SIU Campus Tour
===============

#Conventions

Our URI's format follows a specific format as follows. Note that these conventions do NOT apply to the JSON objects passed as input/output.

`{value}` indicates that value is a variable and will be replaced with a specific value, such as a string or an integer.

`[value]` indicates that value is optional, and may be left out of a query.

In our JSON Objects, we specify the type of a property's value instead of listing an actual value. This is done using the following convention:

`{type}`

Do not confuse this with an object definition.

#Response Codes:

Our server returns the following response codes:

`200 - OK` on success. This may or may not return data, check with the URI specification to determine if you should expect a response body.

`201 - Created` on the creation of an object (often using POST). This may optionally return the id of the created object.

`400 - Bad Request` when the request just doesn't make sense... This may optionally return an Error object defined below.

`404 - Not Found` when the requested URI is not an access  point on our API. The response of this will be empty.

`500 - Internal Server Error` when our server crashes and burns. Our bad if you receive this. This may optionally return an Error object defined below, if so please forward it to us!

Throughout this documentation, we will not specifiy which URIs can return which response codes as it is pretty well common sense.
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

*NOTE*: UserID!=Id. UserID is a string representing the user's unique username, an example could be an email address. The Id is the unique integer assigned to the user by the server.

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
```
{
	"Message":{string}
}
```

##REST Access Points

###/app

Supports:
* `POST`
* `GET`

`POST /app`

POST registers a new application with the server by name. The new name must be unique or an error will be generated.

Input:
```
{
	"Name":{string}
}
```
Output:
```
{
    "Result":{
        "Id":{int}
    }
}
```

`GET /app[/{id}]`

This query returns an array of apps. The id parameter is optional. If the id parameter is provided and matches the id of a registered app, this call will return an array of length one (1) whose sole member is that app. If no app has that registered id, an empty array will be returned.

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
			},...]
}
```

###/user
Supports
* `POST`
* `GET`
* `PATCH`
* `DELETE`


`POST /user`

Post registers a new user with the server.

Input:
```
{
	"Fname":{string},
	"Lname":{string},
	"Birthday":{date},
	"UserID":{string}
}
```

Output:
```
{
    "Result":{
        "Id":{int}
    }
}
```

`GET /user[/{id}]`

This query returns an array of users. The id parameter is optional. If the id parameter is provided and matches the id of a registered user, this call will return an array of length one (1) whose sole member is that user. If no app has that registered id, an empty array will be returned.

Input:
```

```

Output:
```
{
    "Users":[{
        "Id":{integer},
    	"Fname":{string},
    	"Lname":{string},
    	"Birthday":{date},
    	"UserID":{string},
    	"Apps":[{App}]
        },...]
}
```

`PATCH /user/{id}`

Patch updates the specified user with the provided properties. All properties in the json input are optional. Any properties not defined in the input will remain unaltered on the server.

Input:
```
{
    "Fname":{string},
    "Lname":{string},
	"Birthday":{date},
	"UserID":{string}
}
```

Output:
```

```

`DELETE /user/{id}`

This will delete a user from the server. Note, accounts are never truely deleted, the data is simply deactivated.

Input:
```

```

Output
```

```

###/objective
Supports:
* `POST`
* `GET`
* `PUT`

`POST /objective`

Post registers a new objective with the server

Input:
```
{
    "Id":{integer},
    "Name":{string},
    "Description":{string}
}
```
Output:
```
{
    "Result":{
        "Id":{int}
    }
}
```

`GET /objective[/{id}]`

This query returns an array of objectives. The id parameter is optional. If the id parameter is provided and matches the id of a registered objective, this call will return an array of length one (1) whose sole member is that objective. If no objective has that registered id, an empty array will be returned.

Input:
```

```

Output:
```
{
    "Objectives":[{
        "Id":{integer},
        "Name":{string},
        "Description":{string},
        "Objects":[{Object}],
        "Rewards":[{Reward}]
    },...]
}
```

###/app/{id}/user
Supports:
* `POST`
* `DELETE`

`POST /app/{id}/user/{id}`

Registers a user with an app. The first {id} is the App's id. The second {id} is the User's id.

Input:
```

```

Output:
```

```

`DELETE /app/{id}/user/{id}`

Removes a user from an app. Note that this does not actually delete any data, only deactivates it on the server. The first {id} is the App's id. The second {id} is the User's id.

Input:
```

```

Output:
```

```

###/app/{id}/objective/{id}
Supports:
* `POST`
* `DELETE`

`POST /app/{id}/objective/{id}`

Registers an objective with an app. The first {id} belongs to the app and the second {id} belongs to the objective.

Input:
```

```

Output
```

```

`DELETE /app/{id}/objective/{id}`

Removes an objective from an application. The first {id} belongs to the app, the second {id} belongs to the objective. This does not delete an objective's mapping to the app, only deactivates it so that users can no longer complete the objective through this application.

Input:
```

```

Output:
```

```