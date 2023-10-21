package myrouter

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// RouterのGETメソッドのテストコードを実装する
func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  http.Handler
	}{
		{
			name:     "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/helloのエンドポイントにハンドラを追加する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hoge/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/:userのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/:user",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/:userのエンドポイントにハンドラを追加する",
			endpoint: "/:user/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// panicが起きないことを確認する
			r.GET(testcase.endpoint, testcase.handler)
		})
	}
}

// RouterのSearchメソッドのテストコードを実装する
func TestRouter_Search(t *testing.T) {
	testcases := []struct {
		name           string
		setendpoint    string
		accessendpoint string
		handler        http.Handler
	}{
		{
			name:           "/のエンドポイントのハンドラを取得する",
			setendpoint:    "/",
			accessendpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/hoge/fugaのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge/fuga",
			accessendpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/helloのエンドポイントのハンドラを取得する",
			setendpoint:    "/hello",
			accessendpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge",
			accessendpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/piyoのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo",
			accessendpoint: "/piyo",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/piyo/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo/hoge",
			accessendpoint: "/piyo/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/piyo/fugaのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo/fuga",
			accessendpoint: "/piyo/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/hoge/:profileのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge/:profile",
			accessendpoint: "/hoge/www",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/:foo/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/:foo/hoge",
			accessendpoint: "/www/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/:foo/hoge/:barのエンドポイントのハンドラを取得する",
			setendpoint:    "/:foo/hoge/:bar",
			accessendpoint: "/fwww/hoge/bwww",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/:foo/:profileのエンドポイントのハンドラを取得する",
			setendpoint:    "/:foo/:profile",
			accessendpoint: "/fwww/pwww",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:           "/:foo/:profile/:barのエンドポイントのハンドラを取得する",
			setendpoint:    "/:foo/:profile/:bar",
			accessendpoint: "/fwww/pwww/bwww",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		r.GET(testcase.setendpoint, testcase.handler)
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			handler, params := r.Search(http.MethodGet, testcase.accessendpoint)

			for _, p := range params {
				fmt.Printf("key: %s, value: %s\n", p.key, p.value)
			}

			// 関数のポインタを比較する
			if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(testcase.handler).Pointer() {
				t.Errorf("ハンドラが異なります\nwant: %v\ngot: %v", testcase.handler, handler)
			}
		})
	}
}
