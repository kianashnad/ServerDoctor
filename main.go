package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	//Loading environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		// panic and exit if the .env doesn't load
		panic("Failed to load the environment variables.")
	}
	//get the list of URLS to health-check from envs
	urls := strings.Split(os.Getenv("SERVER_DOCTOR_URLS"), ",")
	//create a wait group so the main go routine won't exit
	var wg sync.WaitGroup
	//only 1 is enough, since we won't have any wg.Done()
	wg.Add(1)

	//creating one go routine per URL
	for _, url := range urls {
		go addHealthCheck(url)
	}

	//waiting forever :) so the program won't exit
	wg.Wait()
}

//this function will be used when we need to alert the admin
func sendAlarm(url string, httpError string) {
	//Loading envs for telegram
	botToken := os.Getenv("SERVER_DOCTOR_TG_TOKEN")
	// this is the ID of the user which we will send the alert to
	adminTelegramUUID := os.Getenv("SERVER_DOCTOR_TG_ADMIN_UUID")
	//the message we will send
	message := "⛑ Warning!%0AThere is an issue for this URL:%0A<code>" + url + "</code>%0A%0AThe error:%0A<code>" + httpError + "</code>"
	//the http request to the telegram server
	_, err := http.Get("https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + adminTelegramUUID + "&text=" + message + "&parse_mode=HTML")
	if err != nil {
		//logging the error, if any happens. Which is unlikely.
		log.Println(err)
	}
}

//This function is where the real business happens.
func addHealthCheck(url string) {
	// an infinite loop that executes after a specific duration of time (or sleeps?!)
	for true {
		// here we check the url
		_, err := http.Get(url)
		if err != nil {
			//if the http request fails, send the alarm
			sendAlarm(url, err.Error())
		}
		// the loop sleeps for a duration of time (by default, 10 minutes), until the next execution.
		time.Sleep(10 * time.Minute)
	}
}
