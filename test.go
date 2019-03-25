package main

import (
	wget "./wget"
)

func main() {
	wget.Wget("https://pastebin.com/raw/FLt4jxHJ", "./test_file")
}
