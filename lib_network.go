package easygolang

import (
	"bytes"
	//	"fmt"
	"html"
	"io/ioutil" //ReadAll
	"net"
	"net/http"
	"net/url"
	"time"
)

func NetReadUrlText(url string) (string, error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return "", err1
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	} else {
		return string(data), nil
	}
}

func NetPostUrl(url string) (string, error) {
	client := &http.Client{}

	host_ := StringSplit(url, "?")
	siteHost := host_[0]
	if len(host_) != 2 {
		return "", ErrorWithText("url not splitted by ?")
	}
	data, err := UrlQueryParse(host_[1])
	if err != nil {
		return "", err
	}
	// data := url.Values{}
	// data.Set("client_id", `Lazy Test`)
	// data.Add("client_secret", clientSecret)
	// data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", siteHost, bytes.NewBufferString(data.Encode())) // fmt.Sprintf("%s/token", siteHost)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")    // This makes it work
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func NetPortIsFree(port int, timeout_milliseconds int) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", I2S(port)), time.Millisecond*time.Duration(timeout_milliseconds))
	if conn != nil {
		conn.Close()
		return false
	} else {
		return true
	}
}

func NetLocalAddr() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		Prln(">>" + addr.Network() + " || " + addr.String())
	}
}

func HtmlEscape(str string) string {
	return html.EscapeString(str)
}

func UrlQueryParse(urlstr string) (url.Values, error) {
	return url.ParseQuery(urlstr)
}

func UrlQueryUnescape(urlstr string) string {
	res, err := url.QueryUnescape(urlstr)
	if err != nil {
		return ""
	}
	return res
}

func UrlQueryEscape(urlstr string) string {
	return url.QueryEscape(urlstr)
}
