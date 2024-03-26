## 编译
**Linux下编译**
- Linux
```bash
go build -ldflags="-s -w " -trimpath main.go
upx -9 main
```
- Winodws
```bash
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -ldflags="-s -w " -trimpath main.go
upx -9 main.exe
```
## 使用
```bash
./main -f file(or dir) -p [google TOTP验证码]
```
