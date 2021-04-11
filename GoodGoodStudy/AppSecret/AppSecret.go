package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sort"
	"strconv"
	"strings"
)

var SERVER_NAME = "zonst_by"
var chars []string = []string{
	"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z"}

func getAppId() string {
	buffer := &strings.Builder{}
	uuid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	for i := 0; i < 8; i++ {
		str := uuid[i*4 : i*4+4]
		x, _ := strconv.ParseInt(str, 16, 32)
		buffer.WriteString(chars[x%0x3E])
	}
	return buffer.String()
}

func getAppSecret(appID string) string {
	array := []string{appID, SERVER_NAME}
	//字符串排序
	sort.Strings(array)
	buffer := &strings.Builder{}
	for _, str := range array {
		buffer.WriteString(str)
	}

	newStr := buffer.String()
	h := sha1.New()
	h.Write([]byte(newStr))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func main() {
	appID := getAppId()
	fmt.Println("appID:" + appID)
	fmt.Println("appSecret:" + getAppSecret(appID))
}
