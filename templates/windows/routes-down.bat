@echo off

net session >nul 2>&1
if not %errorlevel% == 0 (
  echo Please Run As Administrator
  pause
  exit 1
)

cd %~dp0
echo importing routing table...
rundll32.exe cmroute.dll,SetRoutes /STATIC_FILE_NAME routes-down.txt /DONT_REQUIRE_URL /IPHLPAPI_ACCESS_DENIED_OK
pause
