# Octernship-Program
# Write a REST API for the input of calories in Golang

### Task Instructions
- API Users must be able to create an account and log in.
- All API calls must be authenticated.
- Implement at least three roles with different permission levels: a regular user would only be able to CRUD on their owned records, a user manager would be able to CRUD only users, and an admin would be able to CRUD all records and users.
- Each entry has a date, time, text, and number of calories.
- If the number of calories is not provided, the API should connect to a Calories API provider (for example, https://www.nutritionix.com) and try to get the number of calories for the entered meal.
- User setting – Expected number of calories per day.
- Each entry should have an extra boolean field set to true if the total for that day is less than the expected number of calories per day, otherwise should be false.
- The API must be able to return data in the JSON format.
- The API should provide filter capabilities for all endpoints that return a list of elements, as well should be able to support pagination.
- Write unit and e2e tests.
- Use any *Golang* web framework
- Use *Postgres* as the database

### Task Expectations
- API Design Best Practices
- Documentation of any assumptions or choices made and why
- Links as citation to any article / code referred to or used
- Unit tests covering the core calories logic
- Appropriate exception handling and error messages
- Code Quality - remove any unnecessary code, avoid large functions
- Good commit history - we won’t accept a repo with a single giant commit 🙅‍♀️

### Task submission
Using the [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow#following-github-flow) for assignment submission
1. Creating a new branch 
2. Raising a Pull Request for submission
3. Using GitHub Discussions to ask any relevant questions regarding the project
4. Final submission Checklist:
- [ ] SUBMISSION.md in the repository / PR, with:
  - [ ] commands to set up the repo (dependencies etc.)
  - [ ] commands to run the test suite
  - [ ] commands to run the API server

# Solution
## Enviroment Varible example
````
Database credentials
DB_HOST="localhost"
DB_DRIVER=postgres
DB_USER="<<DB_USER>>"
DB_PASSWORD="<<DB_PASSWORD>>"
DB_NAME="<<DB_NAME>>"
DB_PORT="3306"

# Default Admin User
ADMIN_USERNAME="<<ADMIN_USERNAME>>"
ADMIN_EMAIL="<<ADMIN_EMAIL>>"
ADMIN_PASSWORD="<<ADMIN_PASSWORD>>"

# Authentication credentials
TOKEN_TTL="1800"
JWT_PRIVATE_KEY="<<JWT_KEY>>"
````