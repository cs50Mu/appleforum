package main

import (
	"appleforum/models"
	"appleforum/views"
	"flag"
	"fmt"
	"log"
	"net/http"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func getCreds() (string, string) {
	fmt.Printf("Username: ")
	var userName string
	fmt.Scan(&userName)
	fmt.Printf("Passwd: ")
	passwd, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Printf("read passwd error: %v", err)
		return "", ""
	}
	fmt.Printf("\nPasswd(again): ")
	passwd2, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Printf("read passwd error: %v", err)
		return "", ""
	}
	fmt.Println()
	if string(passwd) != string(passwd2) {
		fmt.Println("passwd not match")
		return "", ""
	}
	return userName, string(passwd)
}

func main() {
	createUser := flag.Bool("createuser", false, "create normal user")
	migrate := flag.Bool("migrate", false, "migrate")

	flag.Parse()

	// migrate
	if *migrate {
		models.Migration()
		return
	}
	if *createUser {
		// create user
		name, pass := getCreds()
		if name != "" && pass != "" {
			err := models.CreateUser(name, pass, "i@wonder.com")
			if err != nil {
				fmt.Printf("[ERROR] create user failed\n")
			}
		}
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", views.IndexHandler)
	mux.HandleFunc("/login", views.LoginHandler)
	mux.HandleFunc("/logout", views.LogoutHandler)

	addr := ":1984"
	srv := http.Server{
		Handler: mux,
		Addr:    addr,
	}
	log.Printf("start appleforum at %s\n", addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicf("[ERROR]%s\n", err)
	}
}
