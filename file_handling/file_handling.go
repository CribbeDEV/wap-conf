package file_handling

import (
	"bufio"
	"go_wap/types"
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
	new_line = strings.Replace(new_line, "<wlc-main>", types.WLC_HOSTS[config.Wlc.Main]+" "+config.Wlc.Main, -1)
	new_line = strings.Replace(new_line, "<wlc-secondary>", types.WLC_HOSTS[config.Wlc.Secondary]+" "+config.Wlc.Secondary, -1)
	new_line = strings.Replace(new_line, "<site>", config.Site, -1)
	new_line = strings.Replace(new_line, "<mac>", config.OriginalName[2:], -1)
	return new_line
}

func FillTemplate(temp string, config types.Configuration) string {
	var filled string

	template, err := os.Open(temp)
	if err != nil {
		panic(err)
	}
	defer template.Close()

	scanner := bufio.NewScanner(template)
	for scanner.Scan() {
		line := scanner.Text()
		new_line := ReplaceLine(line, config)
		filled += new_line + "\n"
	}

	return filled + "!\n"
}
