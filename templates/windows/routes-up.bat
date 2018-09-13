@echo off

net session >nul 2>&1
if not %errorlevel% == 0 (
  echo 请使用右键 "已管理员身份运行" 此脚本
  pause
  exit 1
)

cd /d %~dp0
echo 导入路由表...
rundll32.exe cmroute.dll,SetRoutes /STATIC_FILE_NAME add.txt /DONT_REQUIRE_URL /IPHLPAPI_ACCESS_DENIED_OK
pause
