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
This function will be called when the user tries to create an account. When the router routes to this function, the incoming HTTP request should contain a JSON file with a username, email, and password. This function will call a helper function to validate the username, email and password from the JSON and adds this user to the database. If the username or password is invalid, the server will send a response indicating these issues in a string. We want to edit this function to check to see if the username, email, or password passed in the JSON exists in the database yet, so there are no multiple users with the same username, email or password.

```GoHome(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

```DisplayCalendar(w http.ResponseWriter, request *http.Request)```
This function still needs to be implemented with our front end. As a preliminary, we have this function set up to receive a request and return that the request was received.

**Testing**

We have multiple testing files throughout the back end. Most of these are unit tests can can be run by opening a terminal, navigating into the desired folder (for example: HabbitRabbit\backend\users) and then run the following command: ```go test -v```


**Other resources we referenced:**

https://medium.com/@anshap1719/getting-started-with-angular-and-go-setting-up-a-boilerplate-project-8c273b81aa6


