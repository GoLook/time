package main

import (
	"os"
	"fmt"
	"strconv"
	"time"
	"regexp"
	"strings"
)

var p = fmt.Println
func main() {
	  args := os.Args
	 lenArg :=args[1]
	  if args == nil || len(args) !=2 || strings.EqualFold(lenArg,"-") || strings.EqualFold(lenArg,"+") {
		 p("Parameter Error !")
		 return
	  }

	//验证字符串是否为数字(包含正负数)
     	isTrue,_ := regexp.MatchString("^[-\\+]?[\\d]*$", lenArg)
	   if isTrue {
		  yday :=yesday(lenArg)
		  fmt.Println(yday)
	}else {
		  fmt.Println("Err:验证字符串是否为数字(包含正负数) ")
	  }

}

func yesday(lenArg string) (string){
	lenArg1, _ := strconv.Atoi(lenArg)
	//isInt :=typeChck(lenArg1,"int")
	//isNum :=  isNumeric(lenArg1)
	//if !isInt {
     //   fmt.Println("typeChckInt:",isInt)
	//	return "2018.9.25 Ver:0.0.1"
	//}else {
			nTime := time.Now()
			yesTime := nTime.AddDate(0, 0, lenArg1)
			logDay := yesTime.Format("20060102")
 		return logDay
	//}
}
