package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetRegion(ipAddr string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://whois.pconline.com.cn/ip.jsp?ip=%v", ipAddr))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
	body, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	str := strings.Replace(string(body), "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	log.Print(str)

	return str, nil
}
