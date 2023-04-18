# Detail work you've completed in Sprint 4

## Frontend:
Added friends page. Connected todo list and calendar to database. Made stylistic changes to the todo list.
## Backend:
We began the sprint by resolving an bug we had with our established register function. 

Later in this sprint, we created a table that would be able to store the calandar evenets of each user registered in our database. Along with this database, we created functions in our api to communicate with the front end and allow for users to add calandar items, delete calandar items, as well as fetch all calandar items for the requested user. 

Further, we created a friend table that could keep track of all user relations, including accepted friends, friend requests and blocked users. Along with this table, we created functions in our api to communicate with the front end and allow for the user to request to be friends with another user, allow a user to accept a firend request, and allow a user to block another. We also implemented a function that allows the client to request the friend status of a user, and recieve a list of all accepted friends, incoming friend requests and blocked friends. 

In order to create these tables and functions to communicate with the front end, we had to create new migration functions and new database functions to help us manage our database. Further, we added all of the tables we created for our backend database into our testing-only database so we would be able to run our unit tests. 

A good amount of time during this sprint was dedicated to connecting our front end with our backend and resolving issues involved in alligning these two parts of our project. Further, a lot of time went into creating meaningful unit tests for the new tables and functionality we added to our backend.

---
# List frontend unit and Cypress tests
---

# List backend unit tests

Unit Tests for Sprint4:
      -TestRegisterFunc(encapsulating LoginFunc):
         This test function was started in Sprint 3, however it is now fully operation with a few changes since the last sprint. Using a database used strictly for testing, a fake user input is loaded into the database. the resulting database configuration is then passed into the login function to see if the user is able to login. This result is then compared to a hardcoded user struct, and if the two login results produce the same string, then the test passed. After this is checked, the data is then deleted from the database to avoid issues with further testing. This allows us to test both the register function and the login function within the same test. For this reason, we did not find it necessart to continue on with the TestLoginFunc. The TestRegisterFunc currently runs and passes.
      -TestEditCal (encapsulating CalStatus and DeleteCal): 
        This test function not only tests the abilities of EditCal, but also tests the DeleteCal and CalStatus functions. For this reason, tests were not written for those functions. This function uses the same testing database as described in the previous test, and adds calendar items to the database. Adding duplicate events to the database is tested, along with deleting an event from the calender and database. After these attempted changes, the function calls CalStatus to see if the proper changes have been made. This result is then compared to a hardcoded string output of the calendar items, and if the two values produced are the same, then the test passes. After this is checked, the data is then deleted from the database to avoid issues with further testing. The TestEditCal function currently runs and passes.
      -TestDeleteToDo: 
         This test function uses the EditToDo function to add tasks to todo_items table in our testing database. After several tasks are added, the tests calls the delete function on an already added item, to check to see if it gets deleted. After the deletetion, the ToDoStatus function is called to check if the item was succesfully deleted. The string produced by that function call is compared to a hardcoded TodoItem struct, and if the two values produced are the same, then the test passes. After this is checked, the data is then deleted from the database to avoid issues with further testing. The TestDeleteToDo function currently runs and passes.
     
    
        

   Unit Tests Completed before Sprit4 (more details listed in Sprint2.md and Sprint3.md):
      -TestPanicHandler 
      -TestAddingToList
      -TestEditToDo: 
      -TestGoHome
      -TestDisplayCalender
      -TestUsernameValidation
      -TestPasswordValidation
      -TestEmailValidation
      -TestHashAndSalt
---
# Show updated documentation for your backend API
