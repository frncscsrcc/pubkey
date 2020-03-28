package operator

import (
	"fmt"
	"log"
	"net/http"
)

func (o *Operator) healtCheck(w http.ResponseWriter, r *http.Request) {
	o.log.Trace.Println("Received healt check request")
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

func (o *Operator) initializeWebServer() {
	http.HandleFunc("/healtcheck", o.healtCheck)

	o.log.Info.Println("WebServer is initialized")
	log.Fatal(http.ListenAndServe(":"+o.config.WebServerPort, nil))

}
