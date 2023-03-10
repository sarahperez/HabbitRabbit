# Project Set Up: Go + SQLITE

Start by installing Go: https://go.dev
Since we are running our code in Visual Studio Code, we also installed this extension: https://code.visualstudio.com/docs/languages/go#:~:text=Go%20in%20Visual%20Studio%20Code%20Using%20the%20Go,the%20Go%20extension%20from%20the%20VS%20Code%20Marketplace.

### Preparing to Access the Database

Our app implements a SQLite database. Before running the server, download and install the https://jmeubank.github.io/tdm-gcc/download/ GCC compiler. We opted for the "Minimal online installer." 

We also downloaded the precompiled binary for our system from the SQLite website: https://www.sqlite.org/download.html (we downloaded the bundle option to check updates to the database using sqlite3 https://www.sqlite.org/cli.html). Make sure to unzip that download into a folder and add the path to the folder to the Path environment variables on your system. Close VS code and re-open it if you had the app running during the installation.

### After Downloading + Installing Go and SQLite

Pull the repository and run the necessary "go get" commands to install the necessary go packages.

### Starting the Server

Open a terminal (in VS Code). Navigate into the backend folder. Then run: ```go run main.go```

# Database Organization
 
| Table in the database       | Information Stored                                                                                              |
| -------------               |:-------------:                                                                                                  |
| users                       | ID (primary key  - will serve as userID accross other tables), username, name, email, hashed password           |
| todo_items                  | ID (generated automaticly- counts rows, not important in porgram), user ID, task description, completion status |

# API

Our server relies on the github.com/gorilla/mux package, which helped us to create a router to handle HTTP requests.

### Handling Functions

```LoginFunc(w http.ResponseWriter, request *http.Request)```
This function is called when the user tries to sign in on the angular app. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username and password. This function will call a helper function to validate the username and password from the JSON, check to see if the user exists in the database and return to the server the user account information and an okay message. If the username or password is invalid or the user was not found in the database, the server will respond with a string indicating these issues.

This is an example of what would be returned to the client if the login was sucessful: 
```
{
"data":{"ID":2,"Username":"UserAlexa","Name":"Alexa","Email":"useralexa@habbitrabbit.com"},
"jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2NzgxNjE3MDgsInVzZXJfaWQiOjJ9.HB8gITSa94poZyVktZFXUkJbIQBTyD69ENdS__Xipkk",
"message":"all is fine"
}
```
---

```RegisterFunc(w http.ResponseWriter, r *http.Request)```
This function will be called when the user tries to create an account. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username, email, and password. This function will call a helper function to validate the username, name, email and password from the JSON. The username and email will then be checked against the database to see if they are already associated with an account. If the paramaters pass these checks, the user is added to the database. If the username, email or password is invalid, or already associated with an account the server will send a response indicating these issues in a string.

This is an example of what would be returned to the client if the regristration was sucessful: 
```
{
"data":{"ID":2,"Username":"UserAlexa","Name":"Alexa","Email":"useralexa@habbitrabbit.com"},
"jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2NzgxNjE3MDgsInVzZXJfaWQiOjJ9.HB8gITSa94poZyVktZFXUkJbIQBTyD69ENdS__Xipkk",
"message":"all is fine"
}
```

---
```UpdateToDo(w http.ResponseWriter, request *http.Request) ```
This function controls the to-do list of the program, the request method determines what action will be taken. The request should send in a user ID and a task description.

Expected json information in request body:

```javascript
{ 
   "user": 1, 
   "description": "buy apples"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | The passed in task will be added to the To-Do list with a completion status of false (incomplete).  |
| DELETE            | The passed in task for the coresponding user will be marked as completed.                           |

Examples of return messages:
```
"task completion status now updated to completed"
```
or
```
"item added"
```

---
```ToDoStatus(w http.ResponseWriter, request *http.Request) ```
This function returns the to do list of the associated user. This could be used when the user first opens their to do list, and can be used to get the updated to do list associated with a user after the add or complete a task. A request to this function should send the appropriate user ID.

Expected JSON information in request body:

```javascript
{ 
   "user": 1
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| GET               | Send the client a JSON file with the completed and incomplete items on the to do list of the user.  |

Example of return JSON:
```
{
 "complete": ["finished task"],
 "incomplete": ["not finished task","not started task"]
}
```

---

```GoHome(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

---

```DisplayCalendar(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

# Testing

We have multiple testing files throughout the back end. Most of these are unit tests can can be run by opening a terminal, navigating into the desired folder (for example: HabbitRabbit\backend\users) and then run the following command: ```go test -v```


**Other resources we referenced:**

https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6


