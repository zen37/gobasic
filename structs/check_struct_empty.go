//https://play.golang.org/p/_lfWfOUvXP9

package main

import (
	"fmt"
	//"time"
)

type Profile struct {
	UserID      string `json:"userid,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Bio         string `json:"bio,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

//var profiledata = Profile{}

func UpdateProfile() (p Profile, err error) {

	p = Profile{UserID: ""}
	return p, nil

}

func main() {
	var profiledata = Profile{}
	//profiledata := Profile{}

	profiledata, _ = UpdateProfile()

	if (profiledata == Profile{}) {
		fmt.Println("no user data provided")
	} else {
		fmt.Println(profiledata)
	}
}
