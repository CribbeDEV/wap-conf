package main

import (
	"bufio"
	"fmt"
	"go_wap/types"
	"os"
	"strings"
)

func GetUserInput() (string, types.UserConfiguration) {

	scanner := bufio.NewScanner(os.Stdin)

	// CSV file
	fmt.Print("\rCSV File: ")
	scanner.Scan()
	csv_location := scanner.Text()

	// Location
	fmt.Print("\rLocation: ")
	scanner.Scan()
	location := scanner.Text()

	// Country Code
	fmt.Print("\rCountry code: ")
	scanner.Scan()
	country := scanner.Text()

	// Site Tag
	fmt.Print("\rSite-tag: ")
	scanner.Scan()
	site := scanner.Text()

	// Primary WLC
	fmt.Print("\rPrimary WLC IP: ")
	scanner.Scan()
	primary := scanner.Text()

	// Secondary WLC
	fmt.Print("\rSecondary WLC IP: ")
	scanner.Scan()
	secondary := scanner.Text()

	fmt.Print("If a output folder already exists, it will be deleted with all its contents. Confirm? (y/N): ")
	scanner.Scan()
	confirmation := scanner.Text()
	if strings.ToLower(confirmation) != "y" {
		fmt.Println("Exiting...")
		os.Exit(3)
	} else {
		err := os.RemoveAll(user_cfg.OutputDirectory)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		os.Mkdir(user_cfg.OutputDirectory, os.ModePerm)
	}

	config := types.UserConfiguration{
		Location:    location,
		CountryCode: country,
		Site:        site,
		Wlc: types.WlcConf{
			Main:      primary,
			Secondary: secondary,
		},
	}
	return csv_location, config
}
