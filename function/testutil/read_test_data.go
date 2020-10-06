package testutil

import "io/ioutil"

// ReadTestData read file and return content
func ReadTestData(filename string) string {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(buf)
}
