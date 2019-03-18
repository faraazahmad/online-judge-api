# grpc-code-execution
a grpc API and server to execute code remotely and return output/error.

## Currently supported languages

### Ruby

endpoint:   `api/ruby/:code_url`<br>
method:     `GET`<br>
body:       
```json
{
    "args": [],     // array of strings
    "stdin": ""     // string
}
```
