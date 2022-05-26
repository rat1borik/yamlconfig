# yamlconfig

1. Make sure you have MinGW installed on Ubuntu:

   ```
   sudo apt-get install gcc-mingw-w64-i686
   ```

   and

   ```
   sudo apt-get install gcc-mingw-w64-x86-64
   ```

2. Compile using the following command:
   ```
   GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -ldflags="-s -w" -buildmode=c-shared -o yamlconfig.dll yamlconfig.go
   ```
