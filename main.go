package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 1. จงหาเส้นทางที่มีค่ามากที่สุด

func maxPathSum(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func OpenJson(file_name string) int {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return 0
	}
	defer file.Close()

	var triangle [][]int
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&triangle)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0
	}
	return maxPathSum(triangle)
}

//2. จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ

func encrypt(text string) string {
	encrypttext := ""

	for i := 0; i < len(text)-1; i++ {
		currentText, _ := strconv.Atoi(string(text[i]))
		nextText, _ := strconv.Atoi(string(text[i+1]))

		if currentText > nextText {
			encrypttext += "L"
		} else if currentText < nextText {
			encrypttext += "R"
		} else {
			encrypttext += "="
		}
	}
	return encrypttext
}

func decrypt(ciphertext string) []int {
	arr := [6]int{0, 0, 0, 0, 0, 0}

	for i := range arr {
		for index, v := range ciphertext {
			if i == index {
				if string(v) == "L" {
					if arr[index+1] >= arr[index] {
						arr[index]++
					}
					if index > 0 && string(ciphertext[index-1]) == "L" {
						if arr[index-1] <= arr[index] {
							arr[index-1] += 1
						}
					}
				} else if string(v) == "R" {
					if arr[index+1] <= arr[index] {
						arr[index+1] = arr[index] + 1
					}
				} else {
					arr[index+1] = arr[index]
				}
			}
			if string(v) == "=" {
				if string(ciphertext[0]) == "=" {
					if arr[index+1] != arr[index] {
						arr[index] = arr[index+1]
					}
				}
			}
		}
	}
	return arr[:]
}

//3. พาย ไฟ ได - Pie Fire Dire

func GetRes(r *gin.Engine) {
	apiURL := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	articleText := string(body)

	targetWord := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}

	words := strings.Fields(articleText)

	var t_boneCount, fatbackCount, pastramiCount, porkCount, meatloafCount, jowlCount, enimCount, bresaolaCount int

	for _, word := range words {
		if strings.ToLower(word) == targetWord[0] {
			t_boneCount++
		} else if strings.ToLower(word) == targetWord[1] {
			fatbackCount++
		} else if strings.ToLower(word) == targetWord[2] {
			pastramiCount++
		} else if strings.ToLower(word) == targetWord[3] {
			porkCount++
		} else if strings.ToLower(word) == targetWord[4] {
			meatloafCount++
		} else if strings.ToLower(word) == targetWord[5] {
			jowlCount++
		} else if strings.ToLower(word) == targetWord[6] {
			enimCount++
		} else if strings.ToLower(word) == targetWord[7] {
			bresaolaCount++
		}
	}

	r.GET("/beef/summary", func(c *gin.Context) {
		beefCuts := map[string]interface{}{
			"beef": map[string]int{
				"t-bone":   t_boneCount,
				"fatback":  fatbackCount,
				"pastrami": pastramiCount,
				"pork":     porkCount,
				"meatloaf": meatloafCount,
				"jowl":     jowlCount,
				"enim":     enimCount,
				"bresaola": bresaolaCount,
			},
		}
		c.JSON(200, beefCuts)
	})
}

func main() {

	maxpathsun := OpenJson("hard.json")
	fmt.Println(maxpathsun)

	testCasesEncrypt := []string{"210122", "000210", "221012", "012001"}
	for _, testcase := range testCasesEncrypt {
		output := encrypt(testcase)
		fmt.Println(output)
	}

	testCasesDecrypt := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}
	for _, testcase := range testCasesDecrypt {
		output := decrypt(testcase)
		fmt.Println(output)
	}

	r := gin.Default()
	GetRes(r)
	r.Run(":3000")
}
