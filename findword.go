package asciiart

func FindWord(sourcefile []byte, output []byte){
	var widthOfString int
	var widthOfWord int
	var count int
	var endOfLetter int
	var oneletter []byte
	var allletters [][]byte
	for i:=0; i<len(output);i++{
		if output[i] == 10 {
			lengthofstring = i
			break
		}
	}
	for i:=0;i<len(output);i++{
		if output[i] == 32 {
			for j:=i;j<len(output);j+lengthofstring{
				if output[j] != 32 {
					count = 0
					break
				} else {
					count++
					if count == 7 {
						allletters = append(allletters, oneletter)
						oneletter = []byte{}
						widthofword = i
						endOfLetter = j
						count = 0
						break
					}
					continue
				}
			}
		}
	for k:=0;k<endOfLetter;{
		if k == widthOfWord {
			k = k + lengthofstring - widthOfWord		
			oneletter = append(oneletter, 10)
			widthOfWord = widthOfWord + lengthofstring
		}
		oneletter = append(oneletter, output[k])
		k++
	}
	}
}
}