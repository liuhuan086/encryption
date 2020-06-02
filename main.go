package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NumStr    = "0123456789"
	CharStr   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SymbolStr = "+=-@#~,.[]()!%^*$"
)

var (
	tag    string
	length int
	f      *os.File
	//pwd []byte
)

func getTag() {
	ok := false
	fmt.Println("请输入标记（例如：微信的密码）：")
	for !ok {
		_, _ = fmt.Scanln(&tag)
		if len(tag) >= 6 {
			ok = true
		} else {
			fmt.Println("不能少于6个字母或两个汉字")
		}
	}
}

func getLength() {
	ok := false
	fmt.Println("请输入需要的密码长度（例如：16）：")
	for !ok {
		inputString := bufio.NewReader(os.Stdin)
		inputRes, err := inputString.ReadString('\n')
		if err != nil {
			fmt.Println("获取输入时错误")
		}
		//unix是"\n", windows是"\r\n"
		trimStr := strings.Trim(inputRes, "\r\n")
		inputInt, err2 := strconv.Atoi(trimStr)
		if err2 != nil {
			fmt.Println("请输入大于0的整数")
		} else {
			length = inputInt
			ok = true
		}
	}
}

func generatePwd() string {
	var pwd []byte = make([]byte, length, length)
	pwdCom := NumStr + CharStr + SymbolStr
	for i := 0; i < length; i++ {
		index := rand.Intn(len(pwdCom))
		pwd[i] = pwdCom[index]
	}
	return string(pwd)
}
func saveFile() {
	finalPwd := generatePwd()
	pwdAndTag := tag + "    " + finalPwd
	f, _ := os.OpenFile("pwd.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	writeStr := strings.Join([]string{pwdAndTag, "\n"}, "")
	buf := []byte(writeStr)
	_, _ = f.Write(buf)
	f.Close()
}


func main() {
	rand.Seed(time.Now().UnixNano())
	getTag()
	getLength()
	saveFile()
}
