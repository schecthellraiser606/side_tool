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

