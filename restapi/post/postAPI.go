// このファイルは `post` パッケージに所属させます．
package post

// 必要パッケージのインポートです．
// 一行ごとに
//   import "パッケージ名"
// と書くこともできますが，こちらの書き方 ( factored import ) が推奨されています．
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
type SampleRequest struct {
	Message string `db:"message" json:"message"`
}

// すべてのリクエストに対し，リクエストボディの内容を `SampleRequest` 型に変換したものを
// レスポンスボディにし返します．
func PostSample(writer http.ResponseWriter, request *http.Request) {
	// このハンドラが実行された旨を標準出力します．
	fmt.Println("Endpoint Hit: PostSample")

	// リクエストボディを取得します．
	// このとき読み込まれるボディは `io.ReadCloser` 型であり，
	// http サーバ側では `Close` を呼び出し，必ず読み込みを終了させなければいけません．
	body := request.Body
	// `Close` を呼び出します．
	// `defer` とは，上位関数の `return` が実行されるまで実行を遅延させます．
	// ここでは， `PostSample` の実行後（`nil` が暗黙的に `return` された後）
	// `body.Close()` が実行されます．
	// `defer` は実行は遅延されますが，値の評価は遅延されません．
	// 仮に引数を渡した場合，引数の値はこの行に到達した時点での値で
	// 固定され，以降の処理の影響を受けることはありません．
	// 一方で， `defer` で遅延させた処理は， `return` した値を書き換えることが可能です．
	defer body.Close()

	// リクエストボディ読み取り用のバッファを定義します．
	buf := new(bytes.Buffer)
	// リクエストボディの内容をバッファに読み取ります．
	io.Copy(buf, body)

	// `SampleRequest` 型の変数を初期化します．
	bodyJson := SampleRequest{}
	// 読み込んだリクエストボディを `SampleRequest` 型に配置し直します．
	// リクエストボディに `message` キーが存在すればその値が，
	// なければ空文字が配置されます．
	// その他のキーは無視されます．
	json.Unmarshal(buf.Bytes(), &bodyJson)

	// json.NewEncoder(writer http.ResponseWriter) で Go の構造体を
	// json 形式に変換するエンコーダを生成し．Encode() で構造体をエンコードします．
	// エンコードする構造体は，先程要素の再配置を行った `SampleRequest`構造体（bodyJson）です
	json.NewEncoder(writer).Encode(bodyJson)
}
