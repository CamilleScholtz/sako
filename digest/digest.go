package digest

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// digestHeaders tracks the state of authentication.
type digestHeaders struct {
	Realm     string
	Qop       string
	Method    string
	Nonce     string
	Algorithm string
	HA1       string
	HA2       string
	Cnonce    string
	Path      string
	Nc        int16
	Username  string
	Password  string
}

// ApplyAuth authenticates against a given URI for the request.
func ApplyAuth(c *http.Client, username, password string,
	req *http.Request) error {
	nreq, err := http.NewRequest(req.Method, req.URL.String(), nil)
	if err != nil {
		return err
	}
	res, err := c.Do(nreq)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	io.Copy(ioutil.Discard, res.Body)

	if res.StatusCode == http.StatusUnauthorized {
		authorization := digestAuthParams(res)

		d := &digestHeaders{
			Realm:     authorization["realm"],
			Qop:       authorization["qop"],
			Nonce:     authorization["nonce"],
			Algorithm: authorization["algorithm"],
			Path:      "/json_rpc",
			Nc:        0x0,
			Username:  username,
			Password:  password,
		}
		d.applyAuth(req)

		return nil
	}

	return fmt.Errorf("response status code should have been 401, it was %v",
		res.StatusCode)
}

func (d *digestHeaders) digestChecksum() {
	// A1
	h := md5.New()
	A1 := fmt.Sprintf("%s:%s:%s", d.Username, d.Realm, d.Password)
	io.WriteString(h, A1)
	d.HA1 = hex.EncodeToString(h.Sum(nil))

	// A2
	h = md5.New()
	A2 := fmt.Sprintf("%s:%s", d.Method, d.Path)
	io.WriteString(h, A2)
	d.HA2 = hex.EncodeToString(h.Sum(nil))
}

// ApplyAuth adds proper auth header to the passed request.
func (d *digestHeaders) applyAuth(req *http.Request) {
	d.Nc += 0x1
	d.Cnonce = randomKey()
	d.Method = req.Method
	d.Path = req.URL.RequestURI()
	d.digestChecksum()

	response := doMD5(strings.Join([]string{d.HA1, d.Nonce, fmt.Sprintf("%08x",
		d.Nc), d.Cnonce, d.Qop, d.HA2}, ":"))
	AuthHeader := fmt.Sprintf(
		`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop=%s, nc=%08x, cnonce="%s"`,
		d.Username, d.Realm, d.Nonce, d.Path, d.Algorithm, response, d.Qop,
		d.Nc, d.Cnonce)

	req.Header.Set("Authorization", AuthHeader)
}

// digestAuthParams parses Authorization header from the http.Request. Returns
// a map of auth parameters or nil if the header is not a valid parsable Digest
// auth header.
func digestAuthParams(r *http.Response) map[string]string {
	s := strings.SplitN(r.Header.Get("Www-Authenticate"), " ", 2)
	if len(s) != 2 || s[0] != "Digest" {
		return nil
	}

	result := map[string]string{}
	for _, kv := range strings.Split(s[1], ",") {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			continue
		}

		result[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}

	return result
}

func randomKey() string {
	k := make([]byte, 8)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}

		bytes += n
	}

	return base64.StdEncoding.EncodeToString(k)
}

// H function for MD5 algorithm (returns a lower-case hex MD5 digest)
func doMD5(data string) string {
	digest := md5.New()
	digest.Write([]byte(data))

	return hex.EncodeToString(digest.Sum(nil))
}
