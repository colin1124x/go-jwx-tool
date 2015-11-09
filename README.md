# jwx tool

提供命令列 jwt 加解密與驗證

### 語法

產生token
```sh
./jwx-tool <secret> < [FILE]

# or

echo <json string> | ./jwx-tool <secret>

```

解析與驗證

```sh
./jwx-tool -d <secret> < [FILE]

# or

echo <token string> | ./jwx-tool -d <secret>
```


