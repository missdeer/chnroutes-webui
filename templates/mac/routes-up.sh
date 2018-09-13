#!/bin/sh

export PATH="/bin:/sbin:/usr/sbin:/usr/bin"
gateway=`route -n get default | grep gateway | awk '{print $2}'`

dscacheutil -flushcache

{{range $i, $ip := .Ips}}route add {{$ip.Ip}}/{{$ip.Cidr}} "${gateway}"
{{end}}
