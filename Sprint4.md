# Detail work you've completed in Sprint 4

## Frontend:

### Friends Tab 
A new component supporting friendships with other users was added this sprint, along with its service to allow for back-end communication. The friends page is split into two main parts, incoming friend requests and search to add friends. For the incoming friend requests, all of the users requests are retrieved and displayed with two buttons for accept or decline under each one. The user can click either button and the appropriate function to reflect this decision is called from the service. If a user has no requests, a message is displayed saying that. As for the search to add, a user can type in a username and send that person a friend request, also supported with an approriate function. The page was styled to be cohesive with the styling of the to-do list. 

### To-Do List
The to-do list page was introduced last sprint but has been further developed this sprint. Icons for a checkbox and garbage can were added, which the user can click to either mark a task as done and see it crossed out or delete a task and see it completely removed from the list. A service to communicate these changes with the back-end was created as well, allowing us to retrieve a users current tasks, update a task as marked done in the database, and remove a task from the database.

### Calendar
The calendar saw no change visually as its part of an outside library, but the service to allow for back-end communication was added this sprint. We are able to retrieve a list of the current users events, add an event for a user to the database, and delete an event for a user from the database.

### Unit Testing
Finally we created more unit tests to test the new components we had added. We also made a couple more cypress tests to test end to end functionality.

## Backend:
We began the sprint by resolving an bug we had with our established register function. 

Later in this sprint, we created a table that would be able to store the calandar evenets of each user registered in our database. Along with this database, we created functions in our api to communicate with the front end and allow for users to add calandar items, delete calandar items, as well as fetch all calandar items for the requested user. 

Further, we created a friend table that could keep track of all user relations, including accepted friends, friend requests and blocked users. Along with this table, we created functions in our api to communicate with the front end and allow for the user to request to be friends with another user, allow a user to accept a firend request, and allow a user to block another. We also implemented a function that allows the client to request the friend status of a user, and recieve a list of all accepted friends, incoming friend requests and blocked friends. 

In order to create these tables and functions to communicate with the front end, we had to create new migration functions and new database functions to help us manage our database. Further, we added all of the tables we created for our backend database into our testing-only database so we would be able to run our unit tests. 

A good amount of time during this sprint was dedicated to connecting our front end with our backend and resolving issues involved in alligning these two parts of our project. Further, a lot of time went into creating meaningful unit tests for the new tables and functionality we added to our backend. We creadted additional functions for testing use only as well.

---
# List frontend unit and Cypress tests
Unit Tests for Sprint 4: 
   -Todo List Component
      -Adding a task
      -Deleting a task
   Friend Component
      -Accepting a friend
      -Denying a friend

Cypress Tests:
   -Login
   -Calendar
      -Add an event
      -Delete an event
   -Todo List
      -Add a task
      -Delete a task

# List backend unit tests

Unit Tests for Sprint4:
      -TestRegisterFunc(encapsulating LoginFunc and DeleteUser):
         This test function was started in Sprint 3, however it is now fully operation with a few changes since the last sprint. Using a database used strictly for testing, a fake user input is loaded into the database. the resulting database configuration is then passed into the login function to see if the user is able to login. This result is then compared to a hardcoded user struct, and if the two login results produce the same string, then the test passed. After this is checked, the data is then deleted from the database to avoid issues with further testing. This allows us to test both the register function and the login function within the same test. For this reason, we did not find it necessart to continue on with the TestLoginFunc. The TestRegisterFunc currently runs and passes.
      -TestDeleteToDo (encapsulating EditToDo and ToDoStatus): 
         This test function uses the EditToDo function to add tasks to todo_items table in our testing database. After several tasks are added, the tests calls the delete function on an already added item, to check to see if it gets deleted. After the deletetion, the ToDoStatus function is called to check if the item was succesfully deleted. The string produced by that function call is compared to a hardcoded TodoItem struct, and if the two values produced are the same, then the test passes. After this is checked, the data is then deleted from the database to avoid issues with further testing. The TestDeleteToDo function currently runs and passes.
      -TestEditCal (encapsulating CalStatus and DeleteCal): 
        This test function not only tests the abilities of EditCal, but also tests the DeleteCal and CalStatus functions. For this reason, tests were not written for those functions. This function uses the same testing database as described in the previous test, and adds calendar items to the database. Adding duplicate events to the database is tested, along with deleting an event from the calender and database. After these attempted changes, the function calls CalStatus to see if the proper changes have been made. This result is then compared to a hardcoded string output of the calendar items, and if the two values produced are the same, then the test passes. After this is checked, the data is then deleted from the database to avoid issues with further testing. The TestEditCal function currently runs and passes.
      -TestDeleteCal (encapsulating EditCal and CalStatus):
         Similar to the TestDeleteToDo function, this test adds several calendar items to the testing database through the EditCal, which are then deleted using the DeleteCal functions.The string produced by that function call is compared to a hardcoded string of the expected struct output. If the two values produced are the same, then the test passes. After this is checked, the data that wasnt already deleted through the test is then deleted from the database to avoid issues with further testing. This Function currently runs and passes.
      -TestRequestFriend:
         This test first several users into the testing database. After these users are registerd, All the users request the same receiver used in testing by calling RequestFriend. After these requests are sent, the status of the receiver friends is checked and compared to a hardcoded string of the expected friend status struct. If the two values produced are the same, then the test passes. After this is checked, the users that were stored in the testing database is deleted from the database to avoid issues with further testing. This function currently runs and passes.
      -TestAcceptFriend:
         Similar to the beginning of TestRequestFriend, this test first several users into the testing database. After these users are registerd, All the users request the same receiver used in testing by calling RequestFriend. After these requests have been sent, the AcceptFriend function is called to accept those friend requests. After these requests are accepted, the status of the receiver friends is checked and compared to a hardcoded string of the expected friend status struct. If the two values produced are the same, then the test passes. After this, the users that were stored in the testing database are deleted to avoid issues with further testing. This function currently runs and passes.
      -TestBlockFriend:
          Similar to TestAcceptFriend, this test first several users into the testing database. After these users are registerd, All the users request the same receiver used in testing by calling RequestFriend. After these requests have been sent, the BlockFriend function is called to block the users that sent the request. After this is processed, the status of the receiver friends is checked and compared to a hardcoded string of the expected friend status struct. If the two values produced are the same, then the test passes. After this, the users that were stored in the testing database are deleted to avoid issues with further testing. This function currently runs and passes.

    
        

   Unit Tests Completed before Sprit4 (more details listed in Sprint2.md and Sprint3.md):
      -TestPanicHandler 
      -TestAddingToList
      -TestEditToDo: 
      -TestUsernameValidation
      -TestPasswordValidation
      -TestEmailValidation
      -TestHashAndSalt

   The following functions and tests were deleted as they were too basic for the work completed this sprint, and therefore uncessesary:
      -TestGoHome
      -TestDisplayCalender
---
# Show updated documentation for your backend API
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
 
| Table in the database       | Columns                                                                                              |
| -------------               |:-------------:                                                                                                  |
| users                       | ID (primary key  - will serve as userID accross other tables), username, name, email, hashed password           |
| todo_items                  | ID (generated automaticly- counts rows, not important in program), user ID, task description, completion status |
| calendar_items              | ID (generated automaticly- counts rows, not important in program), user ID, event ID, startStr, endStr, title          |
| friend_statuses             | ID (generated automaticly- counts rows, not important in program), requester, reciever, status (R, A, B)       |

*Freind database organized according to the single value column database example given here: https://dba.stackexchange.com/questions/135941/designing-a-friendships-database-structure-should-i-use-a-multivalued-column

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
```DeleteUser(w http.ResponseWriter, request *http.Request)```
This function deletes a user from the database. Used in testing only.

Expected json information in request body, request should be sent as a POST:

```javascript
{ 
   "user": 1, 
}
```

This is an example of what would be returned to the client if the login was sucessful: 
```
"user could not be found, not deleted"
or 
"user deleted"
```
---

```RegisterFunc(w http.ResponseWriter, r *http.Request)```
This function will be called when the user tries to create an account. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username, email, and password. This function will call a helper function to validate the username, name, email and password from the JSON. The username and email will then be checked against the database to see if they are already associated with an account. If the paramaters pass these checks, the user is added to the database. If the username, email or password is invalid, or already associated with an account the server will send a response indicating these issues in a string.

This is an example of what would be returned to the client if the registration was sucessful: 
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

Examples of return messages:
```
"task completion status now updated to completed"
```
or
```
"item added"
```
---
```DeleteToDo(w http.ResponseWriter, request *http.Request) ```
This function will delete a to do task from the database.

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
| POST            | The passed in task for the coresponding user will be deleted (used in cases where task is canceled).                |

Examples of return messages:
```
"Task could not be found, so it could not be deleted"
```
or
```
"Task deleted"
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

```EditCal(w http.ResponseWriter, request *http.Request) ```
This function adds an item to a user's calendar.

Expected json information in request body:

```javascript
{ 
  "user": 24,
  "eventID": 2431,
  "startStr": "2023-10-12T10:30:00",
  "endStr": "",
  "title": "Do laundry"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | The passed in task will be added to the database as a calendar item.                                |

Examples of return messages:
```
"item added"
```
---
```DeleteCal(w http.ResponseWriter, request *http.Request) ```
This function will delete a calendar item from the database.

Expected json information in request body:

```javascript
{ 
  "user": 24,
  "eventID": 2431,
  "startStr": "2023-10-12T10:30:00",
  "endStr": "",
  "title": "Do laundry"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | The passed in task will be deleted from the database.  |

Examples of return messages:
```
"task could not be found, so it could not be deleted"
```
or
```
"task deleted"
```

---
```CalStatus(w http.ResponseWriter, request *http.Request) ```
This function returns the calendar items of the associated user.

Expected JSON information in request body:

```javascript
{ 
   "user": 1
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Send the client a JSON file with all of the tasks associated with the passed in user.               |

Example of return JSON:
```
{
    "items": [{"EventID": 2431, "StartStr": "2023-10-12T10:30:00", "EndStr": "", "Title": "Do laundry"},
             {"EventID": 2432, "StartStr": "2023-20-12T10:30:00", "EndStr": "2023-15-12T10:30:00", "Title": "Math Homework"}]
}
```
---
```RequestFriend(w http.ResponseWriter, request *http.Request) ```
Adds a friend request into the database.

Expected JSON information in request body:

```javascript
{ 
  "requester": "username1",
  "reciever": "username2"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Add friend request into the database, send error message if not possible.                           |

Possible outputs:
```
"Requester not found"
"Reciever not found"
"Reciever and Requester not found"
"request sent"
"this connection was already requested, status unchanged"
"already friends"
or
"the reciever has blocked this requester"

```
---
```AcceptFriend(w http.ResponseWriter, request *http.Request) ```
Accepts a friend request.

Expected JSON information in request body:

```javascript
{ 
  "requester": "username1",
  "reciever": "username2"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Update friend status in the database, send error message if not possible.                           |

Example of a possible output:
```
"Requester not found"
"Reciever not found"
"Reciever and Requester not found"
"no request to accept"
"request not found"
"status now updated"
"already accepted"
or
"requester has been blocked"
```
---
```BlockFriend(w http.ResponseWriter, request *http.Request) ```
Blocks a friend request.

Expected JSON information in request body:

```javascript
{ 
  "requester": "username1",
  "reciever": "username2"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Block a friend (from a friend request), or block a user that was once an accepted friend. In either case, the "reciever" does the blocking    |

Example of a possible output:
```
"Requester not found"
"Reciever not found"
"Reciever and Requester not found"
"no request to accept"
"request not found"
or
"requester has already been blocked"
```
```FriendStat(w http.ResponseWriter, request *http.Request) ```
Returns the friends, requests and blocked users associated with the given user.

Expected JSON information in request body:

```javascript
{ 
  "user": 1
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Return all blocked users, friends and pending requests.                                             |

Example of a possible output:
```
{
 "Blocked Users":["username1", "username2"],
 "Friends":["username3"],
 "Requests from":[]
}
```
---
```DeleteRequest(w http.ResponseWriter, request *http.Request) ```
Removes a friend request/status from the database. This function is only used in testing.

Expected JSON information in request body, should be sent as a POST:

```javascript
{ 
  "requester": "username1",
  "reciever": "username2"
}
```

| HTTP request type | Backend functionality                                                                               |
| -------------     |:-------------:                                                                                      |
| OPTIONS           | Handle the pre-flight request.                                                                      |
| POST              | Remove the pending request from the database.                                             |

Example of a possible output:
```
"request deleted"
or
"could not delete request"
```
---
# Testing

We have multiple testing files throughout the back end. Most of these are unit tests can can be run by opening a terminal, navigating into the desired folder (for example: HabbitRabbit\backend\users) and then run the following command: ```go test -v```


**Other resources we referenced:**

https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6
