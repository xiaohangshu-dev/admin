package configs

import (
	"slices"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type OAuth2 struct {
	Issuer   string    `yaml:"issuer" mapstructure:"issuer"`
	LoginURL string    `yaml:"login_url" mapstructure:"login_url"`
	Manager  *Manager  `yaml:"manager" mapstructure:"manager"`
	Clients  []*Client `yaml:"clients" mapstructure:"clients"`
}

type Manager struct {
	AccessTokenExp  int    `yaml:"access_token_exp" mapstructure:"access_token_exp"`
	RefreshTokenExp int    `yaml:"refresh_token_exp" mapstructure:"refresh_token_exp"`
	TokenType       string `yaml:"token_type,omitempty" mapstructure:"token_type"`
	JWTPrivateKey   string `yaml:"jwt_private_key,omitempty" mapstructure:"jwt_private_key"` // JWT签名密钥
	JWTPublicKey    string `yaml:"jwt_public_key,omitempty" mapstructure:"jwt_public_key"`
	SigningMethod   string `yaml:"signing_method,omitempty" mapstructure:"signing_method"`
	Kid             string `yaml:"kid,omitempty" mapstructure:"kid"`
}

type Client struct {
	ID           string   `yaml:"id" mapstructure:"id"`
	Secret       string   `yaml:"secret" mapstructure:"secret"`
	RedirectURIs []string `yaml:"redirect_uris" mapstructure:"redirect_uris"`
	Scopes       []string `yaml:"scopes" mapstructure:"scopes"`
	GrantTypes   []string `yaml:"grant_types,omitempty" mapstructure:"grant_types"`
}

func NewOAuth2(cfgm *viper.Viper, log *zap.Logger) *OAuth2 {

	// 默认配置
	cfg := &OAuth2{
		Issuer:   "http://localhost:8080",
		LoginURL: "http://localhost:8081/login",
		Manager: &Manager{
			AccessTokenExp:  3600,
			RefreshTokenExp: 7200,
			TokenType:       "Bearer",
		},
		Clients: []*Client{},
	}

	// 读取配置文件中的 OAuth2 配置
	if cfgm != nil {
		if err := cfgm.UnmarshalKey("oauth2", cfg); err != nil {
			log.Error("Failed to unmarshal OAuth2 configuration", zap.Error(err))

		}
	}

	return cfg
}

// GetClient 根据客户端ID获取客户端配置
//
// 参数:
//
//	id: 客户端ID
//
// 返回值:
//
//	*Client: 客户端配置
//	error: 错误信息
//
// 错误信息:
//
//	ErrClientNotFound: 客户端ID错误
func (o *OAuth2) GetClient(id string) (*Client, error) {
	for _, c := range o.Clients {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrClientNotFound
}

// ContiansScope 判断客户端是否包含指定Scope
//
// 参数:
//
//	scope: 要判断的Scope
//
// 返回值:
//
//	bool: 包含返回true, 不包含返回false
func (c *Client) ContainsScope(scope string) bool {
	return slices.Contains(c.Scopes, scope)
}

// ContainsGrantType 判断客户端是否包含指定GrantType
//
// 参数:
//
//	grantType: 要判断的GrantType
//
// 返回值:
//
//	bool: 包含返回true, 不包含返回false
func (c *Client) ContainsGrantType(grantType string) bool {
	return slices.Contains(c.GrantTypes, grantType)
}
