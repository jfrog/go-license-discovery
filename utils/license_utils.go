package utils

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"
)

func LicenseToSha256(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", (h.Sum(nil)))
}

func ReadPomComments(pomData string) string {
	reader := strings.NewReader(string(pomData))
	closer := ioutil.NopCloser(reader)
	scanner := bufio.NewScanner(closer)
	var buffer bytes.Buffer
	var line string
	var startBuffer bool
	var stopBuffer bool
	var comment string
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "<!--") {
			startBuffer = true
			line = strings.Replace(line, "<!--", "", -1)
		}
		if strings.HasPrefix(strings.TrimSpace(line), "-->") {
			stopBuffer = true
			line = strings.Replace(line, "-->", "", -1)
		}
		if startBuffer && !stopBuffer {
			if len(strings.TrimSpace(line)) > 0 {
				buffer.WriteString(line)
			}
		}
		if stopBuffer {
			comment = buffer.String()
			break
		}
	}
	closer.Close()
	return comment
}
