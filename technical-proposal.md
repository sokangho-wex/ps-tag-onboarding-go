## What
Third party libraries that are being used for this project:

### Gin - https://github.com/gin-gonic/gin
Gin is a HTTP web framework that has the following features:
- Fast
- API endpoint routing
- Middleware support
- JSON validation
- Error management

## Why

### Gin
For this particular project, Gin can help with the below tasks. 
Note that the below tasks can be done using the std lib, but using Gin will simplify the code.
- Routing API endpoints to their respective handlers
- Getting the request body and query parameters
- Middleware support for returning the correct http code and error message for UserNotFoundError and UserValidationError
- Converting JSON payload to and from structs

Another reason for using Gin for this project is to get familiarity with this framework since it's been used by a couple of TAG services for similar tasks.


