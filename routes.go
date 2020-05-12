package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

// TheIP represent IP address
type TheIP struct {
	IP   string
	Cidr string
	Mask string
}

func (ip TheIP) String() string {
	return fmt.Sprintf("%s/%s", ip.IP, ip.Cidr)
}

// Generate result files from template.
func Generate(templatePath string, ips []TheIP, gateway string) []byte {
	type Data struct {
		Ips     []TheIP
		Gateway string
	}
	data := Data{
		Ips:     ips,
		Gateway: gateway,
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil {
		panic(err)
	}

	return b.Bytes()
}

// FetchIps local ips and remote ips.
func FetchIps() (ips []TheIP) {
	return FetchRemoteIps()
}

// FetchRemoteIps parse APNIC records
func FetchRemoteIps() (ips []TheIP) {
	chinaIPListFile := "china_ip_list.txt"
	inFile, err := os.Open(chinaIPListFile)
	if err != nil {
		fmt.Println("opening china_ip_list.txt failed", err)
		return
	}

	defer inFile.Close()

	re, _ := regexp.Compile(`^([\d\.]+)/(\d{1,2})$`)
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			ss := strings.Split(line, "/")

			ip := ss[0]
			numIP, _ := strconv.Atoi(string(ss[1]))
			cidr := 32 - int(math.Log2(float64(numIP)))
			cidrStr := strconv.Itoa(cidr)

			ips = append(ips, TheIP{ip, cidrStr, cidr2mask(cidr)})
		}
	}

	return ips
}

// Convert CIDR to Mask.
//
// Example:
//
//   cidr2mask(24)       #-> "255.255.255.0"
func cidr2mask(cidr int) string {
	mask := net.CIDRMask(cidr, 32)
	masks := []string{}
	for _, v := range mask {
		masks = append(masks, strconv.Itoa(int(v)))
	}

	return strings.Join(masks, ".")
}
