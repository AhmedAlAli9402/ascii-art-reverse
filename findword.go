package asciiart

import "fmt"

func FindWord(sourcefile []byte, output []byte){
	var widthOfString int
	var widthOfletter int
	var count int
	// var endOfLetter int
	var oneletter []byte
	var allLetters [][]byte
	var count8 int
	fmt.Println(output)
	for i:=0; i<len(output);i++{
		if output[i] == 36 {
			widthOfString = i
			break
		}
	}
	for i:=0;i<len(output);i++{
		if output[i] == 32 {
			for j:=i;j<len(output);j++{
				if output[j] != 32 {
					count = 0
					break
				} else {
					count++
					j= j+widthOfString+1
					if count == 7 {
						widthOfletter = i
						fmt.Println(widthOfletter, "lett1")
						// endOfLetter = j
						count = 0
						break
					}
					continue
				}
			}
		}
		// fmt.Println(widthOfString, "str")
		fmt.Println(widthOfletter, "lett2")
		// fmt.Println(endOfLetter, "end")
		if output[i] == 32 {
		for k:=0;k<len(output);{
		if count8 == 7 {
			break
		}
		if k == widthOfletter {
			k = k + widthOfString - widthOfletter		
			// oneletter = append(oneletter, 10)
			count8++
			widthOfletter = widthOfletter + widthOfString
		}
		oneletter = append(oneletter, output[k])
		k++
	}
	count8 = 0
	if oneletter != nil {
	allLetters = append(allLetters, oneletter)
	oneletter = []byte{}
	i = i + widthOfString+2}
	}
	
}
	fmt.Println(len(allLetters))
	fmt.Println((allLetters[0]))
	// fmt.Println(string(allLetters[2]))
	for i:= 0;i<len(allLetters);i++{
		fmt.Print(string(allLetters[i]))
	}
}
