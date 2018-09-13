@echo on

{{range $i, $ip := .Ips}}route delete {{$ip.Ip}}
{{end}}
