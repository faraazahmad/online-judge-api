package wget

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// Wget : rite results of a GET request to file. If a fileName is given an empty string then the
// last chunk of the input url is used as a filename. Eg: http://foo/baz.jar => baz.jar
func Wget(url, filePath string) {
	if filePath == "" {
		urlSplit := strings.Split(url, "/")
		filePath = urlSplit[len(urlSplit)-1]
	}

	// Get the data
	resp, err := http.Get(url)
	errorChecker(err)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	errorChecker(err)
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	errorChecker(err)
}

// Check if we received an error on our last function call
func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
