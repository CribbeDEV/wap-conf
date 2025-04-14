package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go_wap/file_handling"
	"go_wap/types"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var user_cfg types.UserVariables

func main() {
	fmt.Print(`
       __________  __  ___   ______  __________  _____    _       _____    ____ 
      / ____/ __ \/ / / / | / / __ \/ ____/ __ \/ ___/   | |     / /   |  / __ \
     / / __/ /_/ / / / /  |/ / / / / /_  / / / /\__ \    | | /| / / /| | / /_/ /
    / /_/ / _, _/ /_/ / /|  / /_/ / __/ / /_/ /___/ /    | |/ |/ / ___ |/ ____/ 
    \____/_/ |_|\____/_/ |_/_____/_/    \____//____/     |__/|__/_/  |_/_/

    `)

	cfile, err := os.Open("conf.json")
	if err == nil {
		decoder := json.NewDecoder(cfile)
		err = decoder.Decode(&user_cfg)
		if err != nil {
			fmt.Println("Error decoding conf.json:", err)
			return
		}
	} else {
		user_cfg = types.UserVariables{
			types.TemplateLocations{
				Main:  "templates/main_template.txt",
				Site:  "templates/site_template.txt",
				Reset: "templates/ap_reset.txt",
			},
			"./output",
		}
	}
	defer cfile.Close()

	if user_cfg.Templates.Main == "" || user_cfg.Templates.Site == "" || user_cfg.Templates.Reset == "" {
		fmt.Println("Error: Template paths are empty in conf.json")
		return
	}

	if user_cfg.OutputDirectory != "" {
		if err := os.MkdirAll(user_cfg.OutputDirectory, 0755); err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			return
		}
	}

	file_location, config := GetUserInput()
	if file_location == "" {
		fmt.Println("Error: No file location provided")
		return
	}

	file, err := os.Open(file_location)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Undefined, is variable
	reader.Comma = 0x3B
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("Error: CSV file is empty")
		return
	}

	var wg sync.WaitGroup
	var datas [3]string
	var mu sync.Mutex

	for _, row := range data {
		if len(row) < 2 {
			fmt.Printf("Warning: Skipping row with insufficient data: %v\n", row)
			continue
		}

		wg.Add(1)
		go func(row []string) {
			defer wg.Done()

			new_conf := types.Configuration{
				OriginalName:      row[1],
				NewName:           row[0],
				UserConfiguration: config,
			}

			mu.Lock()
			datas[0] += file_handling.FillTemplate(user_cfg.Templates.GetMain(), new_conf)
			if config.Site != "" {
				datas[1] += file_handling.FillTemplate(user_cfg.Templates.GetSite(), new_conf)
				datas[2] += file_handling.FillTemplate(user_cfg.Templates.GetReset(), new_conf)
			}
			mu.Unlock()
		}(row)
	}

	wg.Wait()

	if datas[0] != "" {
		outputPath := filepath.Join(user_cfg.OutputDirectory, "main_output.txt")
		file_handling.HandleOutput(outputPath, datas[0])
		fmt.Printf("Created main output file: %s\n", outputPath)
	}

	if datas[1] != "" {
		outputPath := filepath.Join(user_cfg.OutputDirectory, "site_output.txt")
		file_handling.HandleOutput(outputPath, datas[1])
		fmt.Printf("Created site output file: %s\n", outputPath)
	}

	if datas[2] != "" {
		outputPath := filepath.Join(user_cfg.OutputDirectory, "reset_output.txt")
		file_handling.HandleOutput(outputPath, datas[2])
		fmt.Printf("Created reset output file: %s\n", outputPath)
	}

	fmt.Println("Processing completed successfully!")
	time.Sleep(5 * time.Second)
}
