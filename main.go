package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

const (
	port = 8080
)

func main() {
	router := httprouter.New()

	router.POST("/webhook", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		defer r.Body.Close()
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Errorf("Bad request: %+v", err)
			return
		}

		log.Infof("Request received: %s", string(s))
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	handler := http.Handler(router)

	addr := fmt.Sprintf(":%d", port)
	log.Infof("%s started on 0.0.0.0%s", "webhook_testr", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("%+v", errors.Wrapf(err, "error starting webhook_testr"))
	}

}
