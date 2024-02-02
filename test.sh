# 実行結果確認用関数
assert() {
    expected="$1"
    input="$2"

    if [ "$input" = "$expected" ]; then
        echo "$input => $expected, Good!"
    else
        echo "$expected expected, but got $input"
        exit 1
    fi
}

# curl wrapper

curl_no_basic() {
    endpoint="$1"
    curl -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Chrome 120.0.0.0 Windows 10.0" http://localhost:8080/$endpoint -o /dev/null -w '%{http_code}\n' -s
}

curl_basic_with_correct() {
    endpoint="$1"
    curl -u testUser:testPassword -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Chrome 120.0.0.0 Windows 10.0" http://localhost:8080/$endpoint -o /dev/null -w '%{http_code}\n' -s
}

curl_basic_with_incorrect() {
    endpoint="$1"
    id="$2"
    password="$3"
    curl -u $id:$password -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Chrome 120.0.0.0 Windows 10.0" http://localhost:8080/$endpoint -o /dev/null -w '%{http_code}\n' -s
}

# 正常系 1 : 対象とする API のみ Basic 認証がかかっているかどうかの確認
echo "正常系 1 : 対象とする API のみ Basic 認証がかかっているかどうか"
assert 200 $(curl_no_basic healthz)
assert 401 $(curl_no_basic todos)

# 正常系 2 : 正しい User ID, Password で Basic 認証をクリアしアクセスできるかどうか。
echo "正常系 2 : 正しい User ID, Password で Basic 認証をクリア出来るか"
assert 200 $(curl_basic_with_correct todos)

# 異常系 1 : 誤った User ID, Password で Basic 認証をクリアできないかどうか。
# 3 パターンのテスト
echo "異常系 1 : 誤った User ID, Password で Basic 認証をクリアできないか"
assert 401 $(curl_basic_with_incorrect todos testuser testPassword)
assert 401 $(curl_basic_with_incorrect todos testUser testpassword)
assert 401 $(curl_basic_with_incorrect todos testuser testpassword)

# 異常系 2 : 空の User ID, Password で Basic 認証をクリアできないかどうか。
echo "異常系 2 : 空の User ID, Password で Basic 認証をクリアできないか"
assert 401 $(curl_basic_with_incorrect todos "" "")

# 異常系 3 : User ID, Password がない状態で Basic 認証をクリアできないかどうか。
echo "異常系 3 : User ID, Password がない状態で Basic 認証をクリアできないか"
assert 401 $(curl_no_basic todos)

# curl http://localhost:8080/delay