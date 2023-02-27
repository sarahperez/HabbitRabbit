**Detail work you've completed in Sprint 2**
2/13/2023
#Front-end Team
Sarah: I worked on setting up Cypress. Downloaded Cypress to the correct directory and created our first simple test. The test checks that the calendar page opens and has a certain URL. Next I will work on component testing using Cypress.
Alexa: Worked on creating login component to link front and back ends

2/15/2023
Sarah: Worked on creating component unit tests using Angular Testing. I believe that I will be unit testing with Angular Testing then E2E testing using Cypress. Today I created a unit test to test the functionality of our sideNav toggling.

#Backend Team
Combined our server with helpful packages and functions from the Backend Banking App tutorial we followed, allowing us to edit and utilize their implementation with our backend functionality. We set up a SQLite database to store the necessary data for our program and famalizarized ourselves with sqlit3 commands to follow our program's interaction with the database. We edited the validation functions from the banking app tutorial to be more streamlined by breaking a singular validation function into three separate functions to handle usernames, passwords, and email addresses separately. We edited the requirements set for valid usernames, passwords, and email addresses. We integrated a Go password package we found online to make these edits. Further, we edited the login and registration functions so that when an invalid parameter is passed in, it is specified in the output what parameter is invalid. We also added a feature to the registration function that checks the passes in username and email to make sure that they are not already associated with a user in our database (this will prevent multiple registered users with the same username or password). We began to incorperate router functions to handle the "OPTIONS" method preflight requests and overall better familarized ourselves with the HTTP request protocol. We resolved issues with our CORS handler (through further reading the Gorilla Mux doccumentation). We also implemented unit testing on most of the functions in our program.

**List unit tests and Cypress tests for frontend**
#Cypress Tests
First simple Cypress test: Clicks calendar button on side nav and checks if the user is redirected to the calendar page. 

#Unit Tests
Sidenav: Toggles sidenav button to check if the first click opens the sidenav bar and the second click closes the sidenav bar.
Full-Calendar Component: 
    Weekend Toggler: Checks if the weekend toggle shows weekends when button is selected and removed weekends when the button is not selected.
    Clicking creates new event
    Clicking event deletes event
Login?? May not get to this
    Check if we can input text into text input
    We did not get to this because we implemented our login was he most recently implemented

**List unit tests for backend**

**Add documentation for your backend API**
