// このファイルは `main` パッケージに所属させます．
package main

// 必要パッケージのインポートです．
// 一行ごとに
//   import "パッケージ名"
// と書くこともできますが，こちらの書き方 ( factored import ) が推奨されています．
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// プロジェクト作成時に `go mod init パッケージ名` を行っていれば，
	// プロジェクト配下のパッケージ名の名前解決ができます．
	"progmate-intern2021/restapi/get"
	"progmate-intern2021/restapi/post"
)

func main() {
	fmt.Println("Start listening...")
	// `http.HandleFunc` を使って，エンドポイントと各ハンドラを設定します．
	// 設定はエンドポイント名ごとに行い，`POST`, `GET` 等のメソッドの判定は
	// 第2引数の `request` 構造体の `Method` メンバで判定できます．
	//
	// ハンドラとして指定する関数の引数は
	//   writer    http.ResponseWriter  : レスポンスを作成するインタフェース
	//   request   *http.Request        : リクエスト（構造体）の参照
	// となっています．
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		// `http` パッケージには http リクエストやステータスコードが定数として
		// 定義されているので，そちらを利用します．
		// 値は `GET`, `POST` といった文字列型なので，
		// string 型変数との比較も可能です．
		case http.MethodPost:
			post.PostSample(writer, request)
		case http.MethodGet:
			get.GetSample(writer, request)
		default:
			json.NewEncoder(writer).Encode(http.StatusInternalServerError)
		}
	})
	// この例では `/query` エンドポイントにハンドラを登録しています．
	// ハンドラとして登録した関数の中身は
	//   * http メソッドが `GET` の場合は `GetQuerySample()` を実行
	//   * それ以外の場合はエラーを返す
	// 無名関数を指定しています．
	// 無名関数の書き方は他の言語とほとんど同じで，
	//   func(引数名 型名...) {任意の処理...}
	// の形式です．
	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			get.GetQuerySample(writer, request)
		default:
			json.NewEncoder(writer).Encode(http.StatusInternalServerError)
		}
	})

	// 第1引数に待ち受けるアドレスを指定し，listenを開始します．
	// 第2引数にはハンドラを指定しますが， `nil` を指定した場合
	// `DefaultServeMux` が使用されます．
	// `Fatal` レベルのログを標準出力に表示します．
	log.Fatal(http.ListenAndServe(":8081", nil))
}
