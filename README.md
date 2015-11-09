# jwx tool

提供命令列 jwt 加解密與驗證

### 語法

產生token
```sh
./jwx-tool <secret> < [FILE]

# or

echo <json string> | ./jwx-tool <secret>

```

example

```sh
echo '{"num":123}' | ./jwx-tool xyz
# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJudW0iOjEyM30.2U4ZKC-yACzysDm_CUI5ONuylTgcQN-wURjqexcqxJY

```


解析與驗證

```sh
./jwx-tool -d <secret> < [FILE]

# or

echo <token string> | ./jwx-tool -d <secret>
```

example
```sh
echo 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJudW0iOjEyM30.2U4ZKC-yACzysDm_CUI5ONuylTgcQN-wURjqexcqxJY' | ./jwx-tool -d aaa
# signature is invalid

echo 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJudW0iOjEyM30.2U4ZKC-yACzysDm_CUI5ONuylTgcQN-wURjqexcqxJY' | ./jwx-tool -d xyz
# {"num":123}
```
