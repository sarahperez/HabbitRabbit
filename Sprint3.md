# Detail work you've completed in Sprint 3
## Frontend:

Registration Page: Created new component for registration page. The registration page contains text input fields for the user to add their name, email, username, and password. Started working on adding the users information to the database. I was having trouble setting up the register service to call the register function that the backend team has made. Currently the registration page routes back to the login page. In practice the user will be able to register (their information will be added to the database), they would be directed back to the login page where they could then login. 

Login Page: Look of login page was changed to be more cohesive with rest of app, button linking to registration page was added

Side-Nav: Added link to to-do page, close button is now functional

Home: Added message that gets stored user info and greets user with their name

To-Do: Created new component for to-do list, basic list interface that allows the user to type a new task and add it to their list

## Backend:

At the begining of this sprint, we consolidated our database and added a table that could keep track of an unordered to-do list for each user. In addition to setting up a todo_items table, we wrote functions for our api that allow users to add items to their to do lists, update the status of an item (mark it complete), and delete an item from their to do list. We influded a functionality that calculates the percentage of overall tasks completed for each user. We hope that this number will provide users with an insight to their productivity metric. 

Further, we did a lot of research into how we would create unit tests for the back end moving forward, as many of our functions will be heavily reliant on a database. We decided to make a seperate database with the same tables as out main database, except this one would be used exclusivly for testing purpouses and be cleared at the end of each test. For example, at the begining of a test, relevant infromation will be added to the table, nessisary functions will then be called, results will be verified (test will pass or fail) and then whatever was added will be deleted. So far this has been the best way we found to run all unit tests in one testing database without conflicts.

We have also done a lot of research into creating a calandar/task scheduler in the backend. We were unsucessful in finding an applicable library/repository that could help us achieve our goals, so we are deciding to move forward by implementing our own functions to handle task scheduling for each user. We hope to find a library however that can store dates and times and organize objects based on these paramaters moving forward. 

Additionaly, we have identified how we would like to store and keep track of user's friends. We have decided to go with a single-valued column approach which will keep track of the requester user, recieveing user, friend status and spefifier (the last user to update the friendship status) instead of a multivalued column approach and hope to have this implemented soon. Here is the ideology we will be modeling our friend table after: https://dba.stackexchange.com/questions/135941/designing-a-friendships-database-structure-should-i-use-a-multivalued-column

# List frontend unit tests

# List backend unit tests
   Unit Tests for Sprint3:
      -TestPanicHandler: 
         This test makes sure the panic handler  function works correctly, It does this by creating a http handler that is supposed to cause a panic, and compares the outcome to the actual reseult of the panic handler. If both functions return the same error message, that means the function is catching the panic, and is working correctly.
      -TestRegisterFunc:
         Using a database used strictly for testing, a fake user input is loaded into the database. the resulting database configuration is then compared to a hardcoded user struct, and if the two registration results produce the same string, then the test passed. After this is checked, the data is then deleted from the database to avoid issues with further testing. This function currently runs, but is not producing the expected output.
      -TestLoginFunc:
         This function works with the http handlers to make sure the right requests are made when a user logs in correctly. As of right now, this function is not implemented correctly and therefore doesnt run.
      -TestAddingToList:
         This function adds a to do list item to the todo list table within the testing database, and compares that to a hardcoded expected result to make sure the task was added correctly. if the two produce the same result, the test passes. The data is then deleted from the database to avoid issues with further testing. This function runs and passes
      -TestEditToDo:
         This test adds items to the two do list table like the previous test, however this function also tests changing the status of list items, as the user has the ability to complete and delete tasks. Throughout this test many changes and edits are made to the todolist, and those values are checked with the hardcoded values. The data is then deleted from the database to avoid issues with further testing. This function currently runs and passes.

   Unit Tests Completed for Sprint2 (more details listed in Sprint2.md): 
      -TestGoHome
      -TestDisplayCalender
      -TestUsernameValidation
      -TestPasswordValidation
      -TestEmailValidation
      -TestHashAndSalt


# Show updated documentation for your backend API 
all of this infromation can also be found in our README.md in our backend folder

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
This function will be called when the user tries to create an account. When the router routes to this function, the incoming HTTP request should contain a JSON file with a name, username, email, and password. This function will call a helper function to validate the username, name, email and password from the JSON. The username and email will then be checked against the database to see if they are already associated with an account. If the paramaters pass these checks, the user is added to the database. If the username, email or password is invalid, or already associated with an account the server will send a response indicating these issues in a string.

This is an example of what would be returned to the client if the regristration was sucessful: 
```
{
 "data":{"ID":2,"Username":"UserAlexa","Name":"Alexa","Email":"useralexa@habbitrabbit.com"},
 "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2NzgxNjE3MDgsInVzZXJfaWQiOjJ9.HB8gITSa94poZyVktZFXUkJbIQBTyD69ENdS__Xipkk",
 "message":"all is fine"
}
```

---
```EditToDo(w http.ResponseWriter, request *http.Request) ```
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
| PUT               | The passed in task for the coresponding user will be marked as completed.                           |
| DELETE            | The passed in task for the coresponding user will be deleted (used in cases where task is canceled).                |

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
This function returns the to do list of the associated user as well as the percentage of completed to incomplete tasks. This could be used when the user first opens their to do list, and can be used to get the updated to do list associated with a user after the add or complete a task. A request to this function should send the appropriate user ID.

Expected JSON information in request body:

```javascript
{ 
   "user": 1
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST               | Send the client a JSON file with the completed and incomplete items on the to do list of the user.  |

Example of return JSON:
```
{
 "Complete": ["finished task"],
 "Incomplete": ["not finished task","not started task"],
 "Percentage of tasks completed": 33
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



