#!/bin/sh

alias route='/system/xbin/busybox route'

{{range $i, $ip := .Ips}}route del -net {{$ip.IP}} netmask {{$ip.Mask}}
{{end}}
