@echo off

net session >nul 2>&1
if not %errorlevel% == 0 (
  echo Please Run As Administrator
  pause
  exit 1
)

{{if (ne .Gateway "$gateway")}}
cd /d %~dp0
for /F "tokens=3" %%i in ('route print ^| findstr "\<0.0.0.0\>"') do  ( set gw=%%i && goto :break )

:break
echo Current default gateway %gw%
route change 0.0.0.0 mask 0.0.0.0 {{ .Gateway }}
{{end}}

echo importing routing table...
rundll32.exe cmroute.dll,SetRoutes /STATIC_FILE_NAME routes-up.txt /DONT_REQUIRE_URL /IPHLPAPI_ACCESS_DENIED_OK

{{if (ne .Gateway "$gateway")}}
echo Recover default gateway to %gw%
route change 0.0.0.0 mask 0.0.0.0 %gw%
route add 0.0.0.0 mask 0.0.0.0 {{ .Gateway }}
{{end}}

pause
