package session

import (
	"encoding/gob"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

const (
	secrekey    string = "xiaohangshu-secret-key"
	sessionName string = "xiaohangshu_session"
)

type (
	Session struct {
		store *sessions.CookieStore
		*zap.Logger
	}
)

// NewSession 创建一个 session 对象
//
// 参数:
//
//	log *zap.Logger: 日志对象
//
// 返回值:
//
//	*Session: session 对象
func NewSession(log *zap.Logger) *Session {

	// 注册 url.Values 类型
	gob.Register(url.Values{})

	store := sessions.NewCookieStore([]byte(secrekey))

	store.Options = &sessions.Options{
		Path: "/",
		// session 有效期
		// 单位秒
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
	}

	return &Session{
		store:  store,
		Logger: log,
	}
}

// Get 获取 session 中指定 key 的值
//
// 参数:
//
//	r *http.Request: 请求对象
//	name string: 要获取的 key 名称
//
// 返回值:
//
//	val interface{}: 获取到的 key 值
//	error: 错误信息
func (s *Session) Get(r *http.Request, name string) (val interface{}, err error) {

	// Get a session.
	session, err := s.store.Get(r, sessionName)
	if err != nil {
		return
	}

	val = session.Values[name]

	return
}

// Set 设置 session 中指定 key 的值
//
// 参数:
//
//	w http.ResponseWriter: 响应对象
//	r *http.Request: 请求对象
//	name string: 要设置的 key 名称
//	val interface{}: 要设置的 key 值
//
// 返回值:
//
//	error: 错误信息
func (s *Session) Set(w http.ResponseWriter, r *http.Request, name string, val interface{}) (err error) {
	// Get a session.
	session, err := s.store.Get(r, sessionName)
	if err != nil {
		return
	}

	session.Values[name] = val

	return session.Save(r, w)

}

// Delete 删除 session 中指定 key 的值
//
// 参数:
//
//	w http.ResponseWriter: 响应对象
//	r *http.Request: 请求对象
//	name string: 要删除的 key 名称
//
// 返回值:
//
//	error: 错误信息
func (s *Session) Delete(w http.ResponseWriter, r *http.Request, name string) (err error) {
	// Get a session.
	session, err := s.store.Get(r, sessionName)
	if err != nil {
		return
	}

	delete(session.Values, name)

	return session.Save(r, w)
}
