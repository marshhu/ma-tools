package simple_crawler

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Fetcher(url string, timeout int) (httpStatus int, body []byte, err error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return 0, nil, errors.New("url格式错误")
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	userAgent := "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"
	req.Header.Add("User-Agent", userAgent)

	//今日头条
	if strings.Contains(url, "www.toutiao.com") {
		cookieValue := "ttcid=a2067f315cdc49e09e94b1d0d81a3e9e38; SLARDAR_WEB_ID=d0ce9243-f932-4a9f-9ac9-60c42a250863; csrftoken=dc2f078f05af27daf0631e6184b2e4a8; WEATHER_CITY=%E5%8C%97%E4%BA%AC; sso_auth_status=1c7a287e4a7fcc377e31a39020a817e8; sso_uid_tt=e94c80692a76c0923a68c822dedf9c48; sso_uid_tt_ss=e94c80692a76c0923a68c822dedf9c48; toutiao_sso_user=efb6d8b362f877c79705f4a38ad9cce9; toutiao_sso_user_ss=efb6d8b362f877c79705f4a38ad9cce9; passport_auth_status=76bfa72c3d5af1c304263e6ae256c368%2C1f404406cf49209f866819dd3fe4f3f0; sid_guard=3041369e2d9012755f74021a5ed45b5b%7C1593145039%7C5184000%7CTue%2C+25-Aug-2020+04%3A17%3A19+GMT; uid_tt=d87c69248d57fa92f3a61f20351849fb; uid_tt_ss=d87c69248d57fa92f3a61f20351849fb; sid_tt=3041369e2d9012755f74021a5ed45b5b; sessionid=3041369e2d9012755f74021a5ed45b5b; sessionid_ss=3041369e2d9012755f74021a5ed45b5b; __utmz=24953151.1593162227.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); tt_webid=6843277474432992782; s_v_web_id=verify_kbyoeexd_hTsS4wcl_Pk5X_4PSn_9Qjn_vkSgGiStj39k; tt_webid=6843277474432992782; tt_scid=7m-U95sGkMlHSVaGOkeKywTxFj5jLpIn7EGJUamPRQEWq1C7kWsd-hfGuJx07G9g6b5f; __utmc=24953151; __utma=24953151.1722430384.1593162227.1593326605.1593332416.3; __utmb=24953151.3.10.1593332416;" +
			fmt.Sprintf(" __ac_nonce=%s; __ac_signature=%s;", "05ef85a610096fa21182b", "_02B4Z6wo00f01i1EgNwAAIBBNBpm6FfcxzItQIRAANW2oCm8GD5cR4IHTVcDfp.VuRl.Z3r5r3jdx5fXLm21e7PJ6yH5B11olUkMfITzsfjVwJcDM8gPv4t3hu0s3ZSojQCPBziiG9D-zWdm56")
		req.Header.Add("cookie", cookieValue)
	}
	resp, error := client.Do(req)
	if error != nil || resp == nil {
		return 0, nil, errors.New("获取失败")
	}

	bodyReader := bufio.NewReader(resp.Body)
	defer resp.Body.Close()
	e, _ := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	resBody, _ := ioutil.ReadAll(utf8Reader)
	return resp.StatusCode, resBody, nil
}

func determineEncoding(r *bufio.Reader) (encoding.Encoding, string) {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8, "utf8"
	}
	e, name, _ := charset.DetermineEncoding(bytes, "")
	if name == "windows-1252" {
		return simplifiedchinese.GBK, "gbk"
	}
	return e, name
}

func newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	userAgent := "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"
	req.Header.Add("User-Agent", userAgent)

	return req, err
}
