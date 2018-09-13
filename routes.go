package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	apnicFile := "apnic.txt"
	inFile, err := os.Open(apnicFile)
	if err != nil {
		fmt.Println("opening apnic.txt failed", err)
		return
	}

	defer inFile.Close()
	body, err := ioutil.ReadAll(inFile)
	if err != nil {
		fmt.Println("reading apnic.txt failed", err)
		return
	}

	re, _ := regexp.Compile(`apnic\|CN\|ipv4\|([\d\.]+)\|(\d+)\|`)
	rows := re.FindAllSubmatch(body, -1)
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		ip := string(row[1])
		numIP, _ := strconv.Atoi(string(row[2]))
		cidr := 32 - int(math.Log2(float64(numIP)))
		cidrStr := strconv.Itoa(cidr)

		ips = append(ips, TheIP{ip, cidrStr, cidr2mask(cidr)})
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
