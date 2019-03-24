package wget

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/http2"
)

const (
	bufSize = 1024 * 8
)

// Wget : rite results of a GET request to file. If a fileName is given an empty string then the
// last chunk of the input url is used as a filename. Eg: http://foo/baz.jar => baz.jar
func Wget(url, fileName string) {
	resp := getResponse(url)
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	writeToFile(fileName, resp)
}

// Make the GET request to a url, return the response
func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	// configure transport for HTTPS and check for error
	errorChecker(http2.ConfigureTransport(tr))
	// create client using Transport
	client := &http.Client{Transport: tr}
	// get response and check for error
	resp, err := client.Get(url)
	errorChecker(err)
	return resp
}

// Write the response of the GET request to file
func writeToFile(fileName string, resp *http.Response) {
	// Credit for this implementation should go to github user billnapier
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	errorChecker(err)
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	errorChecker(err)
	_, err = io.Copy(bufferedWriter, resp.Body)
	errorChecker(err)
}

// Check if we received an error on our last function call
func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
