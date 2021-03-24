// このファイルは `get` パッケージに所属させます．
package get

// 必要パッケージのインポートです．
// 一行ごとに
//   import "パッケージ名"
// と書くこともできますが，こちらの書き方 ( factored import ) が推奨されています．
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// レスポンスとして返す構造体の型定義です．
// string 型の `Message` というフィールドを持つ構造体です．
// また，各フィールドには `tag` をつけることができ，
// DB テーブルのフィールド名との紐付けや，
// JSON の key 名との紐付けが可能です．
// この例では， `Message` フィールドは，
//   * テーブルでは `message` フィールド
//   * JSON では message キー
// として扱われます．
type SampleResponse struct {
	Message string `db:"message" json:"message"`
}

// `GET` メソッドに対するハンドラとして使うことを想定したハンドラです．
// すべてのリクエストに対し，
//   `{ "message": "None" }`
// というボディを持ったレスポンスを返します．
func GetSample(writer http.ResponseWriter, request *http.Request) {
	// このハンドラが実行された旨を標準出力します．
	fmt.Println("Endpoint Hit: GetSample")
	// レスポンスヘッダにステータスコード 200(http.StatusOK) を追加します．
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer http.ResponseWriter) で Go の構造体を
	// json 形式に変換するエンコーダを生成し，Encode() で構造体をエンコードします．
	// エンコードする構造体は， `SampleResponse` 構造体です．
	json.NewEncoder(writer).Encode(SampleResponse{Message: "None"})
}

// すべてのリクエストに対し，
//   `{ "message": "`req` パラメタの値 || 空文字" }`
// というボディを持ったレスポンスを返します．
func GetQuerySample(writer http.ResponseWriter, request *http.Request) {
	// このハンドラが実行された旨を標準出力します．
	fmt.Println("Endpoint Hit: GetQuerySample")
	// URL の path param を取得します．
	// キー名を指定することで，対応する path param を取得します．
	requestId := request.URL.Query().Get("req")
	// レスポンスヘッダにステータスコード 200(http.StatusOK) を追加します．
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer http.ResponseWriter) で Go の構造体を
	// json 形式に変換するエンコーダを生成し．Encode() で構造体をエンコードします．
	// エンコードする構造体は， `SampleResponse` 構造体です．
	json.NewEncoder(writer).Encode(SampleResponse{Message: requestId})
}
