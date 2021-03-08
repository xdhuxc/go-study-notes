package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	f, err := os.Create("/Users/wanghuan/GolandProjects/GoPath/src/github.com/xdhuxc/go-study-notes/advanced/abc.txt")
	if err != nil {
		log.Errorln(err)
		return
	}

	_, _ = f.WriteString("abcd")
}

func reflectTypeName(sample interface{}) string {
	return reflect.TypeOf(sample).String()
}

func otherTest() {
	// 计算 2 的 5 次方
	fmt.Println(math.Pow(2, 5))

	percent := "23.456"
	// 使用 strconv 包的 ParseFloat 方法将字符串转换为 float64
	result, _ := strconv.ParseFloat(percent, 10)
	// 使用 fmt 包的 Sprintf 方法实现保留小数点后两位
	fmt.Println(fmt.Sprintf("%.2f", result))

	s := "Model"

	fmt.Println(strings.ToLower(string(s[0])))

	s1 := "sgt:group"
	fmt.Println(s1[4:])

	m := map[string]string{"": ""}
	fmt.Println(m)

	s = "/a/b/ skd/ dsgreg"
	fmt.Println(strings.Trim(s, "/"))

	dbname := "qiangbb-ss"

	sqlText := fmt.Sprintf("show tables from `%s`;", dbname)
	fmt.Println(sqlText)

	consum := 5288.75 + 4706.30 + 5215.16 + 5504.55 + 4807.83 + 5230.72 + 3980.79
	earn := 17317.46 + 25885.89 + 17359.07 + 39798.76 + 4078.02 + 17340.25 + 18031.53
	fmt.Println(consum)
	fmt.Println(earn)

	fmt.Println(reflectTypeName(consum))

	fun := runtime.FuncForPC(reflect.ValueOf(reflectTypeName).Pointer())
	fmt.Println(fun.Name())
	tokenized := strings.Split(fun.Name(), ".")

	fmt.Println(tokenized)
}
