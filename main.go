package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	DeploymentType uint8 = iota
	StatefulSetType
)

const (
	// PathParameterKind = indicator of Request parameter type "path"
	PathParameterKind = iota

	// QueryParameterKind = indicator of Request parameter type "query"
	QueryParameterKind

	// BodyParameterKind = indicator of Request parameter type "body"
	BodyParameterKind

	// HeaderParameterKind = indicator of Request parameter type "header"
	HeaderParameterKind

	// FormParameterKind = indicator of Request parameter type "form"
	FormParameterKind
)

func panic1() {
	fmt.Println("panic1")
	panic("panic1 -> panic1-----")
}

func panic2() {
	panic1()

	panic("painc2")
}

func panic3() {
	panic("painc2")

	panic1()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandString 函数生成指定长度的字符串
func RandString(n int) string {
	b := make([]byte, n)
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}

	return string(b)
}

func main() {
	fmt.Println(BodyParameterKind)

	var i int64
	i = 12
	j := uintptr(i)
	fmt.Printf("the value is %#v \n", j)

	ptr := unsafe.Pointer(&i)
	fmt.Printf("the ptr is %#v", ptr)

	/*
		s1 := "hello world"
		s2 := s1
		s3 := s2[:]

		fmt.Println(&s1, &s2, &s3)

		s3 = "lalalala"
		fmt.Println(&s3)

	*/

	fmt.Println(strings.ToLower(RandString(6)))
	r := fmt.Sprintf("\"%s\"", "x")
	fmt.Println(r)

	TimestampFormat := "20060102150405"
	fmt.Println(time.Now().Format(TimestampFormat))

	fmt.Println("\n")
	s1 := "1234567890"
	a, err := convert(s1)
	if err != nil {
		return
	}
	fmt.Println(a)

	fmt.Println(convert2(s1))
}

func convert(s string) ([]int, error) {
	target := make([]int, 0, len(s))
	var n int
	var err error
	for _, c := range s {
		fmt.Println(c - 48)
		n, err = strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		target = append(target, n)
	}
	return target, nil
}

func convert2(s string) []int {
	target := make([]int, 0, len(s))
	for _, c := range s {
		target = append(target, int(c-48))
	}

	return target
}

func main12() {
	s := "cbs-flink-sg/main-sg2-test/"
	r := strings.SplitN(s, "/", 2)
	fmt.Println(r)

	defer func() {
		fmt.Println("first panic")
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main panic", r)
		}
	}()
	defer func() {
		fmt.Println("third panic")
	}()

	u, p, err := decodeAuth("YWRtaW46ak9FMjByIXZxMyV+cDFTKw==")
	fmt.Println(err)
	fmt.Println(u)
	fmt.Println(p)

	panic2()
}

func syntax() {
	nicknames := []string{"你算哪块小饼干", "我笑起来真好看", "一条有梦想的咸鱼", "进击的小可爱", "来自星星的我"}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(nicknames[rand.Intn(len(nicknames))])

	fmt.Println(len(nicknames) / 2)

	s := uuid.New().String()
	fmt.Println(s)
	fmt.Println(strings.ReplaceAll(s, "-", ""))

	fmt.Println(strings.Replace(s, "-", "", 2))
}

func decodeAuth(authStr string) (string, string, error) {
	if authStr == "" {
		return "", "", nil
	}

	decLen := base64.StdEncoding.DecodedLen(len(authStr))
	decoded := make([]byte, decLen)
	authByte := []byte(authStr)
	n, err := base64.StdEncoding.Decode(decoded, authByte)
	if err != nil {
		return "", "", err
	}
	if n > decLen {
		return "", "", errors.Errorf("Something went wrong decoding auth config")
	}
	arr := strings.SplitN(string(decoded), ":", 2)
	if len(arr) != 2 {
		return "", "", errors.Errorf("Invalid auth configuration file")
	}
	password := strings.Trim(arr[1], "\x00")
	return arr[0], password, nil
}

func test2() {
	fmt.Println(time.Now())

	// 毫秒
	fmt.Println(time.Now().UnixNano() / 1e6)

	run("nihao", "aa", "bb", "cc")

	fmt.Println(DeploymentType, StatefulSetType)
}

func run(address string, names ...string) {
	fmt.Println("------", address)

	for _, name := range names {
		fmt.Println(name)
	}
}

func test1() {
	message := `用户名不一致：{"shareid": %s, "nameFromDingTalk": %s, "nameFromMySQL": %s}`
	result1 := fmt.Sprintf(message, "wanghuan", "王欢", "wang11")
	fmt.Println(result1)

	params := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
		"d": "d",
	}

	var query string
	length := len(params)
	count := 1
	for key, value := range params {
		query = query + key + "=" + value
		if count < length {
			query = query + "&"
		}

		count = count + 1
	}
	fmt.Println(query)

	u := uuid.New().String()
	fmt.Println(u)

	layout := "2006-01-02 15:04:05"
	lo, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now().In(lo).Format(layout))

	a := "a,d,b"
	fmt.Println(string(a[len(a)-1]))

	as := []string{"a", "b", "c"}
	fmt.Println(as[0 : len(as)-1])

	fmt.Println(time.Now())

	percent := "23.456"

	// 使用 strconv 包的 ParseFloat 方法将字符串转换为 float64
	result, _ := strconv.ParseFloat(percent, 32)
	// 使用 fmt 包的 Sprintf 方法实现保留小数点后两位
	fmt.Println(fmt.Sprintf("%.2f", result))

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int63n(1000))

	t := time.Now()
	h := md5.New()
	io.WriteString(h, t.String())
	fmt.Println(fmt.Sprintf("%x", h.Sum(nil)))

	ss := []string{"a", "b", "c", "d"}
	s := strings.Join(ss, "-")
	fmt.Println(s)
}

func encrypt(shareid string, password string, ou string) (string, error) {
	iv := []byte("This is ushareit")
	data := []byte(fmt.Sprintf("%vqzkc666%v", shareid, ou))
	pass := []byte(password)

	aesKey := []byte(base64.StdEncoding.EncodeToString(data)[:16])
	encodeBytes := []byte(pass)
	//generate encrypted password
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	//PKCS5 padding
	padding := 16 - len(encodeBytes)%16
	padtext := bytes.Repeat([]byte{0}, padding)
	encodeBytes = append(encodeBytes, padtext...)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return hex.EncodeToString(crypted), nil
}
