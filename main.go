package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

var filename string = "passwords.json"

func SetPassword() {
	fmt.Print("For which website do you want to set the password? ")
	var w string
	fmt.Scanln(&w)

	fmt.Print("For which username do you want to set the password? ")
	var u string
	fmt.Scanln(&u)

	b := make([]byte, 22)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error generating random password")
		return
	}
	p := base64.URLEncoding.EncodeToString(b)

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error while opening passwords.json:", err)
		return
	}
	defer f.Close()

	data := make(map[string]string)
	err = json.NewDecoder(f).Decode(&data)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("error decoding JSON:", err)
		return
	}

	k := fmt.Sprintf("%s|%s", w, u)
	data[k] = p

	f.Truncate(0)
	f.Seek(0, 0)
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error encoding JSON:", err)
		return
	}

	fmt.Println("Password set successfully!")
}

func GetPassword() {
	fmt.Print("For which website do you want to get the password? ")
	var w string
	fmt.Scanln(&w)

	fmt.Print("For which username do you want to get the password? ")
	var u string
	fmt.Scanln(&u)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error while opening passwords.json:", err)
		return
	}
	defer f.Close()

	data := make(map[string]string)
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}

	k := fmt.Sprintf("%s|%s", w, u)
	p, ok := data[k]
	if !ok {
		fmt.Println("No password found for that website and username.")
		return
	}

	fmt.Printf("Password for %s in %s: %s\n", u, w, p)
}

func main() {
	for {
		fmt.Print("What do you want to do? ")
		var s string
		fmt.Scanln(&s)

		switch s {
		case "set":
			SetPassword()

		case "get":
			GetPassword()

		default:
			fmt.Println("usage: ./main set | get")
		}
	}
}
