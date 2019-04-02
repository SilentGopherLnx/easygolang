package easygolang

import (
	"html"
	"io/ioutil" //ReadAll
	"net"
	"net/http"
	"net/url"
	"time"
)

func NetReadUrlText(url string) (string, bool) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return "", false
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", false
	} else {
		return string(data), true
	}
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
