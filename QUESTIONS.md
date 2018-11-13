# Questions

For this assignment you also have to answer a couple of questions.
There is no correct answer and none is mandatory, if you don't know just skip it.

 - **What do you think of the initial project structure ?**  
  The project structure look nice, as I research about how to make a clean architecture (in short period). It make clear to understand which folder is responsible for, but it require some knowledge about software pattern design (such as adapter pattern, repository pattern).

 - **What you will improve from your solution ?**
1. configurations (db configs, server port, ...) should be in config file and load to project use viper lib.
2. error handling look messy, need to make it looks more clean.
3. validation query or params should be done in middleware before pass to handler.
4. only repository testing need to test with real DB, the others should use mock.
5. docker setup still need to improve.
6. should assert database table (knights) when start the service.
7. routes file should split by subdomain (for easy to read, otherwise it'll grow too long).
8. still confuse on how to use postgres in docker for development.

 - **For you, what are the boundaries of a service inside a micro-service architecture ?**  
  a service must have high cohesion (enough information to nearly run standalone), low coupling (less dependency not involve too much services), and must do only one responsibility.

 - **For you, what are the most relevant usage for SQL, NoSQL, key-value and document store ?**  
  if need predefined schema, complex query, has many relation between entities => choose SQL.  
  if the data won't need to change much, flexible structure, data very big, need high performance => choose NoSQL.  
  if need to cached, simple message broker, or store something that must able to access incredibly fast (it's in RAM), but maybe not impotant data (can be lost if some error occur) => choose in-memory DB.

