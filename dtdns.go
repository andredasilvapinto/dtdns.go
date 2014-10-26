package main

import (
	"net/http"
	"os"
    "fmt"
    "io/ioutil"
)

const filename string = "ip.dat"
const hostname string = "HOSTNAME"
const password string = "PASSWORD"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getAndRead(url string) string {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(body)
}

func update(ip string) {
	body := getAndRead("https://www.dtdns.com/api/autodns.cfm?id=" + hostname + "&pw=" + password)
    
    if body == ("Host " + hostname + " now points to " + ip + ".\n") {
	    err := ioutil.WriteFile(filename, []byte(ip), 0644)
	    check(err)
    }

    fmt.Println(body)
}

func main() {
	ip := getAndRead("http://echoip.com")

    if _, err := os.Stat(filename); os.IsNotExist(err) {
		update(ip)
	} else {
		storedIp, err := ioutil.ReadFile(filename)
		check(err)

		fmt.Printf("%s == %s? %t\n", ip, storedIp, string(ip) == string(storedIp))

		if string(ip) != string(storedIp) {
			update(ip)
		}
	}
}