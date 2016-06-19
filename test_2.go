package main

import (
    "path/filepath"
    "os"
	"io/ioutil"
    "fmt"
)

func main() {

fmt.Printf(getFilelist("persons"))

}

func getFilelistString(path string) string{
		var infoStr="";
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                tempStr:=read(path)
				if tempStr!=""{
					infoStr=infoStr+"\n"+tempStr
				}
                return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
        }
		
		return infoStr
}
func read(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    fd,err := ioutil.ReadAll(fi)   
    return string(fd)  
} 