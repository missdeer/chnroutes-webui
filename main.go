package main

import (
	"bufio"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Platform string
	URL      string
	FileName string
}

func homePage(c *gin.Context) {
	gateway := c.Param("gateway")
	if gateway == "" {
		gateway = "auto"
	}
	platform := c.Param("platform")

	if gateway == "favicon.ico" && platform == "" {
		c.Data(404, "", []byte{})
		return
	}

	fmt.Println("gateway:", gateway, ", platform:", platform)
	var items []Item
	switch platform {
	case "":
		items = []Item{
			Item{Platform: "Linux", URL: "/" + gateway + "/linux/routes-up.sh", FileName: "routes-up.sh"},
			Item{Platform: "Linux", URL: "/" + gateway + "/linux/routes-down.sh", FileName: "routes-down.sh"},
			Item{Platform: "Linux", URL: "/" + gateway + "/linux/package.zip", FileName: "package.zip"},
			Item{Platform: "Android", URL: "/" + gateway + "/android/routes-up.sh", FileName: "routes-up.sh"},
			Item{Platform: "Android", URL: "/" + gateway + "/android/routes-down.sh", FileName: "routes-down.sh"},
			Item{Platform: "Android", URL: "/" + gateway + "/android/package.zip", FileName: "package.zip"},
			Item{Platform: "macOS", URL: "/" + gateway + "/mac/routes-up.sh", FileName: "routes-up.sh"},
			Item{Platform: "macOS", URL: "/" + gateway + "/mac/routes-down.sh", FileName: "routes-down.sh"},
			Item{Platform: "macOS", URL: "/" + gateway + "/mac/package.zip", FileName: "package.zip"},
			Item{Platform: "ChinaDNS", URL: "/" + gateway + "/chinadns/chnroute.txt", FileName: "chnroute.txt"},
			Item{Platform: "RouterOS", URL: "/" + gateway + "/routeros/routeros-address-list.rsc", FileName: "routeros-address-list.rsc"},
			Item{Platform: "RouterOS", URL: "/" + gateway + "/routeros/routeros.rsc", FileName: "routeros.rsc"},
			Item{Platform: "Windows", URL: "/" + gateway + "/windows/package.zip", FileName: "package.zip"},
		}
	case "routeros":
		items = []Item{
			Item{Platform: "RouterOS", URL: "/" + gateway + "/routeros/routeros-address-list.rsc", FileName: "routeros-address-list.rsc"},
			Item{Platform: "RouterOS", URL: "/" + gateway + "/routeros/routeros.rsc", FileName: "routeros.rsc"},
		}
	case "chinadns":
		items = []Item{
			Item{Platform: "ChinaDNS", URL: "/" + gateway + "/chinadns/chnroute.txt", FileName: "chnroute.txt"},
		}
	case "mac", "linux", "android":
		platformMap := map[string]string{
			"linux":   "Linux",
			"mac":     "macOS",
			"android": "Android",
		}
		items = []Item{
			Item{Platform: platformMap[platform], URL: "/" + gateway + "/" + platform + "/routes-up.sh", FileName: "routes-up.sh"},
			Item{Platform: platformMap[platform], URL: "/" + gateway + "/" + platform + "/routes-down.sh", FileName: "routes-down.sh"},
			Item{Platform: platformMap[platform], URL: "/" + gateway + "/" + platform + "/package.zip", FileName: "package.zip"},
		}
	case "windows":
		items = []Item{
			Item{Platform: "Windows", URL: "/" + gateway + "/windows/package.zip", FileName: "package.zip"},
		}
	}
	if gateway == "auto" {
		gateway = ""
	}
	if platform == "" {
		platform = "all"
	}
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"items":    items,
		"gateway":  gateway,
		"platform": platform,
	})
}

func getFile(c *gin.Context) {
	gateway := c.Param("gateway")
	if gateway == "" {
		c.String(http.StatusBadRequest, "bad gateway request")
		return
	}
	platform := c.Param("platform")
	if platform == "" {
		c.String(http.StatusBadRequest, "bad platform request")
		return
	}

}

func parseAPNIC() {
	apnicFile := "apnic.txt"
	inFile, err := os.Open(apnicFile)
	if err != nil {
		fmt.Println("opening apnic.txt failed", err)
	}

	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var records []string
	records = append(records, "# skip ip in China")
	template := "-A SS -d %s/%d -j RETURN"
	for scanner.Scan() {
		rec := scanner.Text()
		s := strings.Split(rec, "|")
		if len(s) == 7 && s[0] == "apnic" && s[1] == "CN" && s[2] == "ipv4" {
			v, err := strconv.ParseFloat(s[4], 64)
			if err != nil {
				fmt.Printf("converting string %s to float64 failed\n", s[4])
				continue
			}
			mask := 32 - math.Log2(v)
			records = append(records, fmt.Sprintf(template, s[3], int(mask)))
		}
	}
}

func main() {
	addr := ":8089"
	if bind, ok := os.LookupEnv("BIND"); ok {
		addr = bind
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tpl")
	r.GET("/", homePage)
	r.GET("/:gateway", homePage)
	r.GET("/:gateway/:platform", homePage)
	r.GET("/:gateway/:platform/:file", getFile)
	r.Run(addr)
}
