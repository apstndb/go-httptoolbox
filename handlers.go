package httptoolbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func Exec(w http.ResponseWriter, req *http.Request) {
	type request struct {
		File string
		Args []string
	}
	var r request
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	cmd := exec.Command(r.File, r.Args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err = cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func GetHeader(w http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	for k, v := range req.Header {
		fmt.Fprintf(buf, "%s=%v\n", k, v)
	}
	w.Header().Set("content-type", "plain/text")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func GetEnvs(w http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	for _, env := range os.Environ() {
		fmt.Fprintln(buf, env)
	}
	w.WriteHeader(http.StatusOK)
}

func ReadContent(w http.ResponseWriter, req *http.Request) {
	type request struct{ File string }
	var r request
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	bytes, err := ioutil.ReadFile(r.File)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
