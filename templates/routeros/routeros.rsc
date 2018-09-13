# remove existing rules
/ip route rule remove [/ip route rule find table="rosroutes.domestic"]

{{range $i, $ip := .Ips}}/ip route rule add dst-address={{$ip.IP}}/{{$ip.Cidr}} action=lookup table="rosroutes.domestic"
{{end}}
