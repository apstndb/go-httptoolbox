package httptoolbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"io"
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
		return
	}
	cmd := exec.Command(r.File, r.Args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err = cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func WhoAmI(w http.ResponseWriter, r *http.Request) {
	tokenSource, err := google.DefaultTokenSource(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	token, err := tokenSource.Token()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=" + token.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func ExecDmesg(w http.ResponseWriter, _ *http.Request) {
	cmd := exec.Command("/bin/dmesg")
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func GetHeaders(w http.ResponseWriter, req *http.Request) {
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
		return
	}
	bytes, err := ioutil.ReadFile(r.File)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
