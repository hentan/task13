package task13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(inputPath, outputPath string) error {
	data, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("ошибка чтения файла")
		return err
	}

	var patients []Patient
	var patient Patient
	scan := bufio.NewScanner(data)

	for scan.Scan() {
		err := json.Unmarshal(scan.Bytes(), &patient)
		if err != nil {
			fmt.Println("Parsing error", err)
		}
		fmt.Println(patient)
		patients = append(patients, patient)
	}
	fmt.Println(patients)

	jsonData, err := json.MarshalIndent(patients, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	return err
}
