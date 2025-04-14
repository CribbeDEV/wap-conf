package file_handling

import (
	"bufio"
	"go_wap/types"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HandleOutput(file string, data string) {
	output, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	writer.WriteString(data)
	writer.Flush()
}

func ReplaceLine(line string, config types.Configuration) string {
	new_line := strings.Replace(line, "<original-name>", config.OriginalName, -1)
	new_line = strings.Replace(new_line, "<new-name>", config.NewName, -1)
	new_line = strings.Replace(new_line, "<location>", config.Location, -1)
	new_line = strings.Replace(new_line, "<country-code>", config.CountryCode, -1)
	new_line = strings.Replace(new_line, "<wlc-main>", config.Wlc.GetMainWLC(), -1)
	sec, ok := config.Wlc.GetSecondaryWLC()
	if ok {
		new_line = strings.Replace(new_line, "<wlc-secondary>", sec, -1)
	} else {
		return "\b"
	}
	new_line = strings.Replace(new_line, "<site>", config.Site, -1)
	new_line = strings.Replace(new_line, "<mac>", config.OriginalName[2:], -1)
	return new_line
}

func FillTemplate(temp string, config types.Configuration) string {
	var filled string
	var scanner *bufio.Scanner

	if strings.Contains(temp, "http") {
		resp, err := http.Get(temp)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		scanner = bufio.NewScanner(strings.NewReader(string(body)))
	} else {
		template, err := os.Open(temp)
		if err != nil {
			panic(err)
		}
		defer template.Close()

		scanner = bufio.NewScanner(template)
	}

	for scanner.Scan() {
		line := scanner.Text()
		new_line := ReplaceLine(line, config)
		filled += new_line + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return filled + "!\n"
}
