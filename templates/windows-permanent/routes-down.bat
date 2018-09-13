@echo on

{{range $i, $ip := .Ips}}route delete {{$ip.IP}}
{{end}}
