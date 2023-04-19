# SWEProject
Group 40

Front-End Engineers: Sarah Perez and Alexa Melnychuk

Back-End Engineers: Sophie Ruetschi and Brooke Nelson

Project Name: Habbit Rabbit

Project Description: Habbit Rabbit is an all-in-one productivity dashboard targetted towards college students. Some of the included features will be a calendar, habit trackers, task lists, budget organizer, pomodoro timer, etc. The calendar feature will be able to auto schedule events based off of attendee's availibilities. Users will be able to make an account and connect with friends.

The web application will open to the user's dashboard and will have the specific features accessible on the left-hand side of the screen. 

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

# Project Set Up: Frontend Angular Application HabbitRabbit

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.1.3.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).