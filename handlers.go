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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cmd := exec.Command(r.File, r.Args...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err = cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func metadataImpl(path string) ([]byte, error) {
	hreq, err := http.NewRequest(http.MethodGet, "http://metadata.google.internal/" + path, nil)
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Metadata-Flavor", "Google")

	resp, err := http.DefaultClient.Do(hreq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response is not 200: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Metadata(w http.ResponseWriter, r *http.Request) {
	b, err := metadataImpl(r.URL.Query().Get("path"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Email(w http.ResponseWriter, r *http.Request) {
	b, err := metadataImpl("/computeMetadata/v1/instance/service-accounts/default/email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func TokenInfo(w http.ResponseWriter, r *http.Request) {
	tokenSource, err := google.DefaultTokenSource(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := tokenSource.Token()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	w.Write(buf.Bytes())
}

func ReadContent(w http.ResponseWriter, req *http.Request) {
	type request struct{ File string }
	var r request
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := ioutil.ReadFile(r.File)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	req.Write(w)
}
