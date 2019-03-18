# grpc-code-execution
a GRPC API and server to execute code remotely and return output/error.

## Currently supported languages

### Ruby

#### Request format

The server requires three key components from the client request. `code_url` 
is a URL from where the raw code can be downloaded into a file. Only that 
URL containing the raw code can be sent as a parameter in the request.
All other components are supposed to be sent in the request body.

endpoint:   `api/ruby/:code_url`<br>
method:     `GET`<br>
body:       
```json
{
    "args": [],     // array of strings
    "stdin": ""     // string
}
```
