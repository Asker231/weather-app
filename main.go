package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
    const api string = "https://jsonplaceholder.typicode.com/users"  
	func main(){
		users := []User{}
		resp,_ := http.Get(api)
		bytes,_ := io.ReadAll(resp.Body)
		json.Unmarshal(bytes,&users)
		CreateFile("users.json",&users)
		myUsers := Read("users.json")
		FindeUser(myUsers)	
		
	}
	func FindeUser(users []User){
		for i:=0; i < len(users); i++{
			if users[i].Id %2!=0{
				fmt.Println(users[i])
			}
		}
	}
	func Read(filename string) []User{
		users := []User{}
		data,_:= os.ReadFile(filename)
		json.Unmarshal(data,&users)  
		return users
	}

	func CreateFile(filename string,users *[]User){
		file,_:= os.Create(filename)
		jsonFile,_:=json.Marshal(*users)
		file.Write(jsonFile)	
		defer file.Close()
	}

	type User struct {
		Id int64 `json:"id"`
		Name string `json:"name"`
		Username string `json:"username"`
		Email string `json:"email"`
		Address AddressEntity `json:"address"`
		Phone string `json:"phone"`
		Website string `json:"website"`
		Company CompanyEntity `json:"company"`
	
	}
	
	type AddressEntity struct {
		Street string `json:"street"`
		Suite string `json:"suite"`
		City string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo GeoEntity `json:"geo"`
	
	}
	
	type GeoEntity struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	
	}
	
	type CompanyEntity struct {
		Name string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs string `json:"bs"`
	
	}
	
	
	

	