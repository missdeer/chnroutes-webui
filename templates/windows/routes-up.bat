@echo off

net session >nul 2>&1
if not %errorlevel% == 0 (
  echo ��ʹ���Ҽ� "�ѹ���Ա�������" �˽ű�
  pause
  exit 1
)

cd /d %~dp0
echo ����·�ɱ�...
rundll32.exe cmroute.dll,SetRoutes /STATIC_FILE_NAME add.txt /DONT_REQUIRE_URL /IPHLPAPI_ACCESS_DENIED_OK
pause
