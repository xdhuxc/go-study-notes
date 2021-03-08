package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strings"

	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

// errorHandler handle error and return a bool value
// return false will break conext, return true will handle next context
type errorHandler func(resp *restful.Response, err error) bool

// TokenExtractor extract jwt token
type tokenExtractor func(name string, req *restful.Request) (string, error)

// customValidator jwt standard validator
type customValidator func(config *Config, req *restful.Request, resp *restful.Response) error

// Provider JWT validator handler
type Provider struct {
	Config *Config
}

var (
	DefaultUserKey  string = "x-xsp-user"
	DefaultRoleKey  string = "x-xsp-role"
	DefaultGroupKey string = "x-xsp-group"

	CustomHeaderPrefix string = "X"
)

//Config JWT Middleware config
type Config struct {
	// Name token name, default: Authorization
	Name string
	// Signing key to validate token
	SigningKey string
	// ErrorHandler validate error handler, default: defaultOnError
	ErrorHandler errorHandler
	// Extractor extract jwt token, default extract from header: defaultExtractorFromHeader
	Extractor tokenExtractor
	// EnableAuthOnOptions http option method validate switch
	EnableAuthOnOptions bool
	// SigningMethod sign method, default: HS256
	SigningMethod gojwt.SigningMethod
	// ExcludeURL exclude url will skip jwt validator
	ExcludeURL []string
	// ExcludePrefix exclude url prefix will skip jwt validator
	ExcludePrefix []string
	// ContextKey Context key to store user information from the token into context.
	// Optional. Default value "".
	ContextKey string

	// Default value "x-xsp-user".
	UserKey string

	// Default value "x-xsp-role".
	RoleKey string
	// Default value "x-xsp-group".
	GroupKey string

	// CustomValidator custom validator suggestion flow：
	// 1. check exlude url, and exclude url prefix
	// 2. extract token string
	// 3. check token sign
	// 4. check token ttl
	// 5. save custom value to conext after check passed
	// 6. handle addon validator
	CustomValidator customValidator
	// validationKeyGetter
	validationKeyGetter gojwt.Keyfunc
}

// New create a JWT provider
func New(config *Config) (*Provider, error) {
	if config.Name == "" {
		config.Name = "Authorization"
	}
	if config.ContextKey == "" {
		config.ContextKey = config.Name
	}
	if config.UserKey == "" {
		config.UserKey = DefaultUserKey
	}
	if config.RoleKey == "" {
		config.RoleKey = DefaultUserKey
	}
	if config.GroupKey == "" {
		config.GroupKey = DefaultGroupKey
	}
	if config.ErrorHandler == nil {
		config.ErrorHandler = defaultOnError
	}
	if config.Extractor == nil {
		config.Extractor = defaultExtractor
	}
	if config.SigningMethod == nil {
		config.SigningMethod = gojwt.SigningMethodHS256
	}

	if config.SigningKey == "" {
		return nil, fmt.Errorf("JWT 需要 SigningKey")
	} else {
		config.validationKeyGetter = func(token *gojwt.Token) (interface{}, error) {
			return []byte(config.SigningKey), nil
		}
	}
	if config.CustomValidator == nil {
		config.CustomValidator = defaultCheckJWT
	}

	return &Provider{config}, nil
}

// defaultOnError default error handler
// return false will break conext, return true will handle next context
func defaultOnError(resp *restful.Response, err error) bool {
	resp.WriteHeader(http.StatusUnauthorized)
	resp.Write([]byte(err.Error()))
	log.Errorln(err)
	return false
}

// defaultExtractorFromHeader extract token from header
func defaultExtractorFromHeader(name string, req *restful.Request) (string, error) {
	// 对于 X 开头的，直接取值即可
	if strings.HasPrefix(name, CustomHeaderPrefix) {
		return req.HeaderParameter(name), nil
	}
	header := req.HeaderParameter(name)
	if govalidator.IsNull(header) {
		return "", nil // No error, just no token
	}
	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(header, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("请求头 authorization 的格式为：Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// defaultExtractorFromQuery extract token from query
func defaultExtractorFromQuery(name string, req *restful.Request) (string, error) {
	return req.QueryParameter(name), nil
}

// defaultExtractor
func defaultExtractor(name string, req *restful.Request) (string, error) {
	var token string
	var err error

	token, err = defaultExtractorFromHeader(name, req)
	if err != nil {
		return "", err
	}
	if !govalidator.IsNull(token) {
		return token, nil
	}

	return defaultExtractorFromQuery(name, req)
}

func (p *Provider) Auth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if err := p.Config.CustomValidator(p.Config, req, resp); err != nil {
		log.Errorln(err)
		if r := p.Config.ErrorHandler(resp, err); r == false {
			return
		}
	}

	chain.ProcessFilter(req, resp)
}

// defaultCheckJWT execlude token check flow, or returns error
// 1. check exlude url, and exclude url prefix
// 2. extract token string
// 3. check token sign
// 4. check token ttl
// 5. save custom value to conext after check passed
func defaultCheckJWT(config *Config, req *restful.Request, resp *restful.Response) error {
	// extract token
	token, err := config.Extractor(config.Name, req)

	if err != nil {
		return fmt.Errorf("解析 token 时发生错误，错误为：%s", err)
	}
	if token == "" {
		// no token
		return fmt.Errorf("没有找到认证 token")
	}

	// parse token value
	parsedToken, err := gojwt.Parse(token, config.validationKeyGetter)
	if err != nil {
		return fmt.Errorf("解析 token 时发生错误，token 为：%s，错误为：%s", token, err)
	}

	if config.SigningMethod != nil && config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		return fmt.Errorf("error validating token algorithm: %s", message)
	}

	if !parsedToken.Valid {
		return fmt.Errorf("token is invalid")
	}

	claims := parsedToken.Claims.(gojwt.MapClaims)
	if _, ok := claims[config.ContextKey]; !ok {
		return fmt.Errorf("token not has useful info")
	}

	infos := strings.Split(claims[config.ContextKey].(string), "|")

	if len(infos) < 3 {
		return fmt.Errorf("token not has enough useful info")
	}

	resp.AddHeader(config.UserKey, infos[0])
	resp.AddHeader(config.GroupKey, infos[1])
	resp.AddHeader(config.RoleKey, infos[2])

	// save custom value to context
	req.SetAttribute(config.UserKey, infos[0])
	req.SetAttribute(config.GroupKey, infos[1])
	req.SetAttribute(config.RoleKey, infos[2])

	return nil
}
