# remove existing rules
/ip firewall address-list remove [/ip firewall address-list find list="rosroutes.domestic"]

{{range $i, $ip := .Ips}}/ip firewall address-list add address={{$ip.IP}}/{{$ip.Cidr}} list="rosroutes.domestic"
{{end}}
