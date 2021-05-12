package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("123123"))

func SetAlert(w http.ResponseWriter,r *http.Request,message string) error {
	session,err := store.Get(r,"alert-go")
	if err != nil {
		fmt.Println(err)
		return err
	}
	session.AddFlash(message)

	return session.Save(r,w)
}

func GetAlert(w http.ResponseWriter,r *http.Request) map[string]interface{} {
	session,err := store.Get(r,"alert-go")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := session.Flashes()

	if(len(flashes) > 0){
		data["is_alert"] = true
		data["message"] = flashes[0]
	}else{
		data["is_alert"] = false
		data["message"] = nil
	}
	fmt.Println(data)
	session.Save(r,w)

	return data
}