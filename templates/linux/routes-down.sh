#!/bin/sh

export PATH="/bin:/sbin:/usr/sbin:/usr/bin"

ip -batch - <<EOF
  {{range $i, $ip := .Ips}}route del {{$ip.IP}}/{{$ip.Cidr}}
  {{end}}
EOF
