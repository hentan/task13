package task13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type ByAge []Patient

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

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

	sort.Sort(ByAge(patients))

	jsonData, err := json.MarshalIndent(patients, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	return err
}
