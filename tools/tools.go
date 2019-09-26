package main

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"html/template"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

// convert interface array to string array
func Convert2Strings(source []interface{}) []string {
	var target []string
	for _, item := range source {
		switch v := item.(type) {
		case string:
			target = append(target, v)
		case int:
			target = append(target, strconv.FormatInt(int64(v), 10))
		default:
			// maybe occur an error
			target = append(target, item.(string))
		}
	}
	return target
}

type RenderResult struct {
	output string
}

func (r *RenderResult) Write(b []byte) (n int, err error) {
	r.output += string(b)
	return len(b), nil
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToByte(v string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&v))))
}

// render data into email template
func RenderEamilFromFile(fullTmplPath string, result interface{}) (string, error) {
	// fullTmplPath := fmt.Sprintf("%s/report.html", config.GetConfig().EmailConfig.TmplPath)
	t, err := template.ParseFiles(fullTmplPath)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	resultWriter := &RenderResult{}
	if err := t.Execute(resultWriter, result); err != nil {
		log.Errorln(err)
		return "", err
	}

	return resultWriter.output, nil
}

// render data into email string template
func renderEamilFromString(name string, templateContent string, result interface{}) (string, error) {
	t := template.New(name)

	_, err := t.Parse(templateContent)
	if err != nil {
		return "", err
	}

	resultWriter := &RenderResult{}
	if err := t.Execute(resultWriter, result); err != nil {
		log.Errorln(err)
		return "", err
	}

	return resultWriter.output, nil
}

// GET request
func DoGetWithRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("the response is not ok, the status code is %s", strconv.Itoa(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Errorf("the response body is %s", string(body))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return body, nil
}

// GET request
func DoGetWithURL(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("the response is not ok, the status code is %s", strconv.Itoa(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Errorf("the response body is %s", string(body))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return body, nil
}

// send email
func SendEmail(from string, to string, subject string, cc string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetHeader("Cc", cc)
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "noreply@xdhuxc.com", "xdhuxc")
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
