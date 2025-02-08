package ascci

import (

	"bufio"
	"fmt"
	"os"
	"strings"

)


func Makeandreturn(word, detect string) [][]string {


		file := "./assci/" + detect
		
		alphabetascciart , err := os.Open(file)
		
		if err != nil {
			fmt.Println("not open")
		}
		contianalphabet  := bufio.NewScanner(alphabetascciart)
		m := make(map[string]string)

		curent := ""
		count := 32
		for contianalphabet.Scan() {
			line := contianalphabet.Text() 
			if line == "" {
				if curent != "" && count <= 126 { 
				 
				m[string(rune(count))] = curent
				curent = ""
				count++
				}
			}else {
			curent += line + "\n"
			}
		}
	defer alphabetascciart.Close()
	
	if curent != "" && count <= 126 {
		m[string(rune(count))] = curent
	} 
	l := [][]string{}
	//adddemilateur := []string{""}
	output := make([]string, 8)
	i := 0
	for i < len(word) {
		charStr := string(word[i])
		if word[i] == 10 {
			// her i print new line
			j := 0
			for j < 8 {
				output[j] += "\n"
				j++
			} 
			l = append(l, output)
			output = make([]string, 8)
			i += 1
			continue
		}
	
		lines := strings.Split(m[charStr], "\n")
	
		for j, line := range lines {
			if j < 8 {
				if i == len(word)-1 {
					output[j] += line + "\n"
				} else {
					output[j] += line 
				}
			}
		}	
		i++
	}
	
	if len(output[0]) > 0 {
		l = append(l, output)
	}
	
	return l
}	