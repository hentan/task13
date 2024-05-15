package task13

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Patient struct {
	XMLName xml.Name `xml:"Patient"`
	Name    string   `xml:"Name" json:"name"`
	Age     int      `xml:"Age" json:"age"`
	Email   string   `xml:"Email" json:"email"`
}

type Patients struct {
	XMLName xml.Name  `xml:"Patients"`
	List    []Patient `xml:"Patient"`
}

func Do(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	defer file.Close()

	var patients Patients
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		var patient Patient
		err := json.Unmarshal(scan.Bytes(), &patient)
		if err != nil {
			return fmt.Errorf("parsing error: %v", err)
		}
		patients.List = append(patients.List, patient)
	}

	if err := scan.Err(); err != nil {
		return fmt.Errorf("ошибка сканера: %v", err)
	}

	res, err := xml.MarshalIndent(patients, "", "    ")
	if err != nil {
		return fmt.Errorf("ошибка создания xml: %v", err)
	}

	xmlHeader := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	res = append(xmlHeader, res...)

	err = os.WriteFile(outputPath, res, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %v", err)
	}

	return nil
}
