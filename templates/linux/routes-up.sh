#!/bin/sh

export PATH="/bin:/sbin:/usr/sbin:/usr/bin"
gateway=$(ip route show 0/0 | grep via | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+')

ip -batch - <<EOF
  {{range $i, $ip := .Ips}}route add {{$ip.Ip}}/{{$ip.Cidr}} via $gateway
  {{end}}
EOF
