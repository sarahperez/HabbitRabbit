# **Work You've Completed in Sprint 2**

**Frontend Team:**

2/13/2023
Sarah: I worked on setting up Cypress. Downloaded Cypress to the correct directory and created our first simple test. The test checks that the calendar page opens and has a certain URL. Next I will work on component testing using Cypress.
Alexa: Worked on creating login component to link front and back ends

2/15/2023
Sarah: Worked on creating component unit tests using Angular Testing. I believe that I will be unit testing with Angular Testing then E2E testing using Cypress. Today I created a unit test to test the functionality of our sideNav toggling.

Created more unit tests for the full calendar component.

**Backend Team:**

Combined our server with helpful packages and functions from the Backend Banking App tutorial we followed, allowing us to edit and utilize their implementation with our backend functionality. We set up a SQLite database to store the necessary data for our program and famalizarized ourselves with sqlit3 commands to follow our program's interaction with the database. We edited the validation functions from the banking app tutorial to be more streamlined by breaking a singular validation function into three separate functions to handle usernames, passwords, and email addresses separately. We edited the requirements set for valid usernames, passwords, and email addresses. We integrated a Go password package we found online to make these edits. Further, we edited the login and registration functions so that when an invalid parameter is passed in, it is specified in the output what parameter is invalid. We also added a feature to the registration function that checks the passes in username and email to make sure that they are not already associated with a user in our database (this will prevent multiple registered users with the same username or password). We began to incorperate router functions to handle the "OPTIONS" method preflight requests and overall better familarized ourselves with the HTTP request protocol. We resolved issues with our CORS handler (through further reading the Gorilla Mux doccumentation). We also implemented unit testing on most of the functions in our program.

# **Unit Tests and Cypress Tests for Frontend**
#Cypress Tests
First simple Cypress test: Clicks calendar button on side nav and checks if the user is redirected to the calendar page. This unit test was passing then we implemented the login page. Since the test does not account for the login page, this test is not surrently passing.

#Unit Tests
Sidenav: Toggles sidenav button to check if the first click opens the sidenav bar and the second click closes the sidenav bar.
Full-Calendar Component: 
    Weekend Toggler: Checks if the weekend toggle shows weekends when button is selected and removed weekends when the button is not selected.
    Clicking creates new event: Checks to see if date selects function is called. This function adds a new event. 
    Clicking event deletes event; Checks to see if date click function is called. This function deleted the event thats clicked.
Login?? May not get to this
    Check if we can input text into text input
    We did not get to this because our login was the most recently implemented. This will be tested in a later sprint.

# **Unit tests for Backend:**
For unit testing in the backend, we decided to use the testing package that Golang already offers. 

Although many tests have been written, we found it important to have functions for testing GoHome and DisplayCalendar for this Sprint since we were integrating the frontend and the backend. These unit testing functions allowed us to test the handler response by making sure the response recorder returned the response that was expected. Both these functions are implemented and are passing. 

Some other important unit testing functions that were implemented were the validation functions for the email, username, and password. These functions checked to make sure invalid passwords, usernames, and emails were caught, and to make sure every requirement for these were met. All of these functions are implemented and are passing. We are continuing to work on implementing more testing functions and getting them to run, however we found it more important to focus on those most relevant to this sprint, as we wanted to highlight those.

Summary:
    Unit Tests that are implemented, running, and passing:
        -TestGoHome
        -TestDisplayCalender
        -TestUsernameValidation
        -TestPasswordValidation
        -TestEmailValidation
        -TestHashAndSalt

    Unit tests that are in the process of being implemented:
     -TestLoginFunc
     -TestGetAccount 
     -TestUpdateAccount
     -TestPrepareToken


# **Documentation For Your Backend API- (The following summary can also be found in the readme.md file inside the backend folder)**

**Setting Up Go**

Start by installing Go: https://go.dev
Since we are running our code in Visual Studio Code, we also installed this extension: https://code.visualstudio.com/docs/languages/go#:~:text=Go%20in%20Visual%20Studio%20Code%20Using%20the%20Go,the%20Go%20extension%20from%20the%20VS%20Code%20Marketplace.



**Preparing to Access the Database**

Our app implements a SQLite database. Before running the server, download and install the https://jmeubank.github.io/tdm-gcc/download/ GCC compiler. We opted for the "Minimal online installer." 

We also downloaded the precompiled binary for our system from the SQLite website: https://www.sqlite.org/download.html (we downloaded the bundle option to check updates to the database using sqlite3 https://www.sqlite.org/cli.html). Make sure to unzip that download into a folder and add the path to the folder to the Path environment variables on your system. Close VS code and re-open it if you had the app running during the installation.

**After Downloading + Installing Go and SQLite**

Pull the repository and run the necessary "go get" commands to install the necessary go packages.

**Starting the Server**

Open a terminal (in VS Code). Navigate into the backend folder. Then run: ```go run main.go```

**Basics of Our Server**

Our server relies on the github.com/gorilla/mux package, which helped us to create a router to handle HTTP requests.

**Current Handling Functions in our API**

```LoginFunc(w http.ResponseWriter, request *http.Request)```
This function is called when the user tries to sign in on the angular app. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username and password. This function will call a helper function to validate the username and password from the JSON, check to see if the user exists in the database and return to the server the user account information and an okay message. If the username or password is invalid or the user was not found in the database, the server will respond with a string indicating these issues.

```RegisterFunc(w http.ResponseWriter, r *http.Request)```
This function will be called when the user tries to create an account. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username, email, and password. This function will call a helper function to validate the username, email and password from the JSON. The username and email will then be checked against the database to see if they are already associated with an account. If the paramaters pass these checks, the user is added to the database. If the username, email or password is invalid, or already associated with an account the server will send a response indicating these issues in a string.

```GoHome(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

```DisplayCalendar(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

**Testing**

We have multiple testing files throughout the back end. Most of these are unit tests can can be run by opening a terminal, navigating into the desired folder (for example: HabbitRabbit\backend\users) and then run the following command: ```go test -v```
