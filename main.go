package main

import (
	cr "crypto/rand"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"shortURL/models"
	"shortURL/mysql"
	"strings"
	"time"
)

//	func updateCount(count *int64, db *gorm.DB) {
//		for true {
//			db.Model(&models.All76CarOwnerInformation{}).Count(count)
//			println("now", time.Now().UnixNano(), *count)
//			time.Sleep(5 * time.Minute)
//		}
//	}
const S = "1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-" +
	"1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-" +
	"1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-" +
	"1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-" +
	"1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-" +
	"1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$_-"

type JsonRequest struct {
	// Define the structure of your JSON request
	// For example, if your JSON looks like {"data": "example"}
	Url string `json:"url"`
}

func main() {

	db := mysql.DB
	//go updateCount(&count, db)

	r := gin.Default()
	r.LoadHTMLGlob("016html/*")
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/:short", func(c *gin.Context) {

		short := c.Request.URL.String()[1:]
		//println(short)
		if isGood(short) {

			su := &models.SU{}
			db.Where("short_url = ?", short).First(&su)

			if len(su.Url) == 0 {
				c.JSON(200, gin.H{
					"code": 404,
					"msg":  "not found",
				})
			} else {
				//println(su.Url)
				c.HTML(200, "t.html", gin.H{
					"url": su.Url,
				})
			}

		} else {
			c.JSON(200, gin.H{
				"code": 300,
				"msg":  "链接不正确",
			})
		}

	})

	r.GET("/1/*shortUrl", func(c *gin.Context) {

		// Access data from JSON request
		u := c.Request.URL.String()[3:]
		//println(u)
		if !isURL(u) {
			c.JSON(200, gin.H{
				"code": 300,
				"msg":  "url格式不正确",
			})
			return
		}
		str := generateRandomString(4)
		for true {
			su := &models.SU{}
			db.Where("short_url = ?", str).First(&su)
			if len(su.ShortUrl) == 0 {
				break
			}
			str = generateRandomString(4)
		}
		su := models.SU{ShortUrl: str, Url: u}
		db.Create(&su)

		c.String(200, "%s", "https://t.016.wiki/"+str)

		// Your business logic here
	})

	r.POST("/", func(c *gin.Context) {
		// Parse JSON request body
		jsonRequest := JsonRequest{}
		if err := c.ShouldBindJSON(&jsonRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Access data from JSON request
		u := jsonRequest.Url
		//println(u)
		if !isURL(u) {
			c.JSON(200, gin.H{
				"code": 300,
				"msg":  "url格式不正确",
			})
			return
		}
		str := generateRandomString(4)
		for true {
			su := &models.SU{}
			db.Where("short_url = ?", str).First(&su)
			if len(su.ShortUrl) == 0 {
				break
			}
			str = generateRandomString(4)
		}
		su := models.SU{ShortUrl: str, Url: u}
		db.Create(&su)

		c.JSON(200, gin.H{
			"code": 200,
			"url":  "https://t.016.wiki/" + str,
		})

		// Your business logic here
	})

	err := r.Run(":6668")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
func isGood(str string) bool {
	if strings.TrimSpace(str) == "" {
		return false
	}
	//println(len(str))
	if len(str) != 4 && len(str) != 6 {
		return false
	}
	//println(1111111111111)
	if !allCharactersInStringBExistInStringA(S, str) {
		return false
	}

	return true

}

func allCharactersInStringBExistInStringA(strA, strB string) bool {
	// 创建一个 map 用于存储字符串 A 中的字符
	charMap := make(map[rune]bool)

	// 将字符串 A 中的字符添加到 map 中
	for _, charA := range strA {
		charMap[charA] = true
	}

	// 遍历字符串 B 中的每个字符，如果在 map 中不存在，则返回 false
	for _, charB := range strB {
		if _, found := charMap[charB]; !found {
			return false
		}
	}

	return true
}

func isURL(s string) bool {
	return govalidator.IsURL(s)
}

var count = 0

func generateRandomString(length int) string {
	/*// 生成随机字符串的字符集
	charSet := S

	// 种子是当前时间的纳秒部分
	rand.Seed(time.Now().UnixNano())

	// 生成随机字符串
	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(result)*/

	return generateRandomString1(length)

}

func generateRandomString1(length int) string {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间戳作为随机种子
	characters := S
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func generateRandomString2(length int) string {
	bytes := make([]byte, length)
	_, err := cr.Read(bytes)
	if err != nil {
		panic(err)
		return generateRandomString1(length)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
