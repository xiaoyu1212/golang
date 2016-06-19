package main

import S "strings"
import ("fmt"
		"os"  
		"io/ioutil"
		"strconv"
		"path/filepath"
)
type Person struct {
    name string
    age int
	}
func main() {
	var youngPersonSlice []Person
	var middlePersonSlice []Person
	var oldPersonSlice []Person
	const YOUNG = 18
	const MIDDLE = 60
	str:=getFilelistString("persons")
	
	strArray:=S.Split(str,",")
	for i :=0 ;i<len(strArray);i++{
		if strArray[i]!=""{
			name_age:=S.Split(strArray[i],":")
			name:=name_age[0]
			age,error :=strconv.Atoi(name_age[1])
			if error==nil{
				var p Person=Person{name,age}
				if p.age>0&&p.age<YOUNG+1{
					youngPersonSlice=append(youngPersonSlice,p)
				}else if p.age>YOUNG&&p.age<61{
					middlePersonSlice=append(middlePersonSlice,p)
				}else if p.age>MIDDLE{
					oldPersonSlice=append(oldPersonSlice,p)
				}
			}
		}
	}
	
	fmt.Printf("\n")
	//print youngPerson
	fmt.Printf("Young:0-18\n")
	fmt.Printf("============\n")
	for _,p:= range youngPersonSlice{
		fmt.Printf(p.name+":"+strconv.Itoa(p.age)+"\n")
	}
	fmt.Printf("\n")
	//print middlePerson
	fmt.Printf("Middle:19-60\n")
	fmt.Printf("============\n")
	for _,p:= range middlePersonSlice{
		fmt.Printf(p.name+":"+strconv.Itoa(p.age)+"\n")
	}
	fmt.Printf("\n")
	//print oldPerson
	fmt.Printf("Old:61-?\n")
	fmt.Printf("============\n")
	for _,p:= range oldPersonSlice{
		fmt.Printf(p.name+":"+strconv.Itoa(p.age)+"\n")
	}	

}

func getFilelistString(path string) string{
		var infoStr="";
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                tempStr:=read(path)
				if tempStr!=""{
					infoStr=infoStr+","+tempStr
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