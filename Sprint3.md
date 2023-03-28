# Detail work you've completed in Sprint 3
## Frontend:

## Backend:

At the begining of this sprint, we consolidated our database and added a table that could keep track of an unordered to-do list for each user. In addition to setting up a todo_items table, we wrot functions for our api that allow users to add items to their to do lists, update the status of an item (mark it complete), and delete an item from their to do list. We influded a functionality that calculates the percentage of overall tasks completed for each user. We hope that this number will provide users with an insight to their productivity metric. 

Further, we did a lot of research into how we would create unit tests for the back end moving forward, as many of our functions will be heavily reliant on a database. We decided to make a seperate database with the same tables as out main database, except this one would be used exclusivly for testing purpouses and be cleared at the end of each test. For example, at the begining of a test, relevant infromation will be added to the table, nessisary functions will then be called, results will be verified (test will pass or fail) and then whatever was added will be deleted. So far this has been the best way we found to run all unit tests in one testing database without conflicts.

We have also done a lot of research into creating a calandar/task scheduler in the backend. We were unsucessful in finding an applicable library/repository that could help us achieve our goals, so we are deciding to move forward by implementing our own functions to handle task scheduling for each user. We hope to find a library however that can store dates and times and organize objects based on these paramaters moving forward. 

Additionaly, we have identified how we would like to store and keep track of user's friends. We have decided to go with a single-valued column approach which will keep track of the requester user, recieveing user, friend status and spefifier (the last user to update the friendship status) instead of a multivalued column approach and hope to have this implemented soon. Here is the ideology we will be modeling our friend table after: https://dba.stackexchange.com/questions/135941/designing-a-friendships-database-structure-should-i-use-a-multivalued-column

# List frontend unit tests

# List backend unit tests

# Show updated documentation for your backend API 
