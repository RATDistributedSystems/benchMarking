package main


import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

var totalCount float64
var totalTime float64

func main(){

	filename := []string{"buy.txt", "cancelBuy.txt", 
	"cancelBuyTrigger.txt", "cancelSell.txt", "cancelSellTrigger.txt", 
	"commitBuy.txt", "commitSell.txt", "displaySummary.txt", "dumpLog.txt", "quote.txt", 
	"quoteCache.txt", "sell.txt", "setBuyAmount.txt", "setBuyTrigger.txt", "setSellAmount.txt", "setSellTrigger.txt",
	"addCQL.txt"}

	for i := 0; i < len(filename); i++ {
		createFile("converted/" + filename[i] + ".MS")
		processFile(filename[i],"converted/" + filename[i] + ".MS")
		totalCount = 0
		totalTime = 0
	}

}

func processFile(filename string, output string){


    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	totalCount += 1.00
    	//fmt.Println(scanner.Text())
    	if(strings.Contains(scanner.Text(),"µs")){	//microseconds case
    		text := scanner.Text()
    		text = strings.Replace(text, "µs", "", -1)
    		textInt, _ := strconv.ParseFloat(text, 64)
    		textInt = textInt / 1000
    		text = strconv.FormatFloat(textInt, 'E', -1, 64)
    		appendToText(output, text)
    		totalTime += textInt

    	}else if(strings.Contains(scanner.Text(),"ms")){	//milliseconds case
    		text := scanner.Text()
    		text = strings.Replace(text, "ms", "", -1)
    		textInt, _ := strconv.ParseFloat(text, 64)
    		appendToText(output, text)
    		totalTime += textInt

    	}else if(strings.Contains(scanner.Text(),"s")){		//seconds case
    		//fmt.Println(scanner.Text())
    		text := scanner.Text()
    		text = strings.Replace(text, "s", "", -1)
    		textInt, _ := strconv.ParseFloat(text, 64)
    		textInt = textInt * 1000
    		text = strconv.FormatFloat(textInt, 'E', -1, 64)
    		appendToText(output, text)
    		totalTime += textInt

    	}else if(strings.Contains(scanner.Text(),"m")){		//minutes case "2m16.128419775s"
    		textArray := strings.Split(scanner.Text(),"m")
    		text := textArray[0]
    		text = strings.Replace(text, "m", "", -1)
    		textInt, _ := strconv.ParseFloat(text, 64)
    		textInt = textInt * 60 * 1000
    		textSeconds := strings.Replace(textArray[1], "s", "", -1)
    		textIntSeconds, _ := strconv.ParseFloat(textSeconds,64)
    		textIntSeconds = textIntSeconds * 1000
    		textInt += textInt + textIntSeconds

    		text = strconv.FormatFloat(textInt, 'E', -1, 64)
    		appendToText(output, text)
    		totalTime += textInt
    	}else{
    		text := scanner.Text()
    		appendToText(output, text)
    	}
    }


    if err := scanner.Err(); err != nil {
        panic(err)
    }

    average := strconv.FormatFloat(computeAverage(), 'E', -1, 64)
    fmt.Println("========> Completed Converting: " + filename + "\n")
    appendToText(output, "AVERAGE =" + average)

}


func appendToText(filename string, text string){

	//openfile
	text = text + "\n"
	//fmt.Println("Appending To File: " + filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
  	  panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
   	 panic(err)
	}

}

func createFile(filename string) {

	// detect if file exists
	var _, err = os.Stat(filename)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)

		if err!= nil { 
			return 
		}

		defer file.Close()
	}

	fmt.Println("==> Creating file", filename)
}

func computeAverage() float64{
	return (totalTime / totalCount);
}