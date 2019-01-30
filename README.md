# rest-hello

This aims at learning to create a rest api with go.

This repo has three go sub-projects that can help in absorbing multiple REST api development concepts.
Can think of it as step-by-step demonstration in builing production grade REST API.

1. basic_todo.go demos a very naive rest api with very little noise
2. model-handler-app demos how to modularize router and model to form an improved api
3. tdd-model-handler demos using test driven development to create a production grade REST API

Additional Notes:
1. model can be thought of as a mocked version of database. It uses slice to simulate a database for the CRUD operations.
2. The tests cover a lot of cases but are not complete. few negative cases are not covered; the intent being this is a demonstration - not attempting to write a full blown production code 