# remove existing rules
/ip firewall address-list remove [/ip firewall address-list find list="freedomroutes.domestic"]

{{range $i, $ip := .Ips}}/ip firewall address-list add address={{$ip.IP}}/{{$ip.Cidr}} list="freedomroutes.domestic"
{{end}}
