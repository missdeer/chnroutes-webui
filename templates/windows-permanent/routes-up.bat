@echo off

for /F "tokens=3" %%* in ('route print ^| findstr "\<0.0.0.0\>"') do set "gw=%%*"

IF %gw%==%%* (
  echo 错误: 未能找到网关
  pause
  exit
)

ipconfig /flushdns

@echo on

{{range $i, $ip := .Ips}}route add -p {{$ip.Ip}} mask {{$ip.Mask}} %gw% metric 5
{{end}}
