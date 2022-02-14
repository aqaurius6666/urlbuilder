package urlbuilder

import (
	"fmt"
	"strings"
)

var (
	DEFAULT_HOST       = "localhost"
	DEFAULT_SCHEME     = "http"
	DEFAULT_PORT       = 80
	DEFAULT_QUERY      = make(map[string]string)
	DEFAULT_PATH_PARAM = make(map[string]string)
	DEFAULT_PATH       = ""
	DEFAULT_BASE_PATH  = ""
)

type UrlBuilder struct {
	Host      string
	Scheme    string
	Port      int
	Query     map[string]string
	Path      string
	PathParam map[string]string
	BasePath  string
}

func (s UrlBuilder) Build() string {
	url := s.buildBase()
	if s.Path != "" {
		url = fmt.Sprintf("%s/%s", url, s.buildPath())
	}
	if len(s.Query) != 0 {
		url = fmt.Sprintf("%s?%s", url, s.buildQuery())
	}
	return url
}

func (s UrlBuilder) buildBase() string {
	if s.Port == 80 && s.Scheme == "http" {
		return fmt.Sprintf("%s://%s", s.Scheme, s.Host)
	}
	if s.Port == 443 && s.Scheme == "https" {
		return fmt.Sprintf("%s://%s", s.Scheme, s.Host)
	}
	return fmt.Sprintf("%s://%s:%d%s", s.Scheme, s.Host, s.Port, s.BasePath)
}

func (s UrlBuilder) buildPath() string {
	path := s.Path
	for k, v := range s.PathParam {
		path = strings.Replace(path, fmt.Sprintf(":%s", k), v, 1)
	}
	for string(path[0]) == "/" {
		path = path[1:]
	}
	return path
}

func (s UrlBuilder) buildQuery() string {
	queries := make([]string, 0)
	for k, v := range s.Query {
		queries = append(queries, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(queries, "&")
}

func (s *UrlBuilder) WithPort(port int) *UrlBuilder {
	s.Port = port
	return s
}
func (s *UrlBuilder) WithPath(path string) *UrlBuilder {
	s.Path = path
	return s
}

func (s *UrlBuilder) WithScheme(scheme string) *UrlBuilder {
	s.Scheme = scheme
	return s
}

func (s *UrlBuilder) WithHost(host string) *UrlBuilder {
	s.Host = host
	return s
}
func (s *UrlBuilder) WithQueryParam(key string, value string) *UrlBuilder {
	s.Query[key] = value
	return s
}
func (s *UrlBuilder) WithPathParam(key string, value string) *UrlBuilder {
	s.PathParam[key] = value
	return s
}
func (s *UrlBuilder) WithBasePath(base string) *UrlBuilder {
	s.BasePath = base
	return s
}
func (s *UrlBuilder) WithPathParams(pathParam map[string]string) *UrlBuilder {
	for k, v := range pathParam {
		s.PathParam[k] = v
	}
	return s
}
func (s *UrlBuilder) WithQueryParams(queryParam map[string]string) *UrlBuilder {
	for k, v := range queryParam {
		s.Query[k] = v
	}
	return s
}

func copyMap(src map[string]string) map[string]string {
	dest := make(map[string]string)
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
func From(builder UrlBuilder) *UrlBuilder {

	return &UrlBuilder{
		Host:      builder.Host,
		Scheme:    builder.Scheme,
		Port:      builder.Port,
		Query:     copyMap(builder.Query),
		Path:      builder.Path,
		PathParam: copyMap(builder.PathParam),
		BasePath:  builder.BasePath,
	}
}

func Init() *UrlBuilder {
	return &UrlBuilder{
		Host:      DEFAULT_HOST,
		Scheme:    DEFAULT_SCHEME,
		Port:      DEFAULT_PORT,
		Query:     DEFAULT_QUERY,
		Path:      DEFAULT_PATH,
		PathParam: DEFAULT_PATH_PARAM,
		BasePath:  DEFAULT_BASE_PATH,
	}
}
