package utils
import (	
	"fmt"
	"encoding/base64"
	"golang.org/x/term"
	"os"
)
func RegisterUser(){
				fmt.Println("enter username")
				username := ""
				fmt.Scanln(&username)
				fmt.Println("enter password")
				bytepassword, err := term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil{
					fmt.Println("error reading user input")	
					return 
				}
				fmt.Println("confirm password")
				confirmBytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil{
					fmt.Println("error reading password")
					return
				}

				if string(bytepassword) != string(confirmBytePassword){
					fmt.Println("password are not the same")
					return
				}
				
				userString := fmt.Sprintf("%s:%s", username, string(bytepassword))
				authStringBase64 := base64.StdEncoding.EncodeToString([]byte(userString))
				fmt.Println(authStringBase64)
				decodedString, _ := base64.StdEncoding.DecodeString(authStringBase64)
				fmt.Println(string(decodedString))
}
