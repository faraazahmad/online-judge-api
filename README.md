# grpc-code-execution
a GRPC API and server to execute code remotely and return output/error.

## Currently supported languages

### Ruby

#### Request format

endpoint:   `/ruby`<br>
method:     `GET`<br>

The server requires three components from the client request. 
* `url` :   A URL where the code is stored in raw format. For example ,
            https://pastebin.com/raw/FLt4jxHJ
* `args`:   Arguments to be passed to the interpreter for running the code.
            For example, `-a`, `-c` etc.
* `stdin`:  The input to be provided to the code (including newlines).

body:       
```json
{
    "url": "",
    "args": "",
    "stdin": ""
}
```

### Response format
