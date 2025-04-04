package gentleman

import (
	gocontext "context"
	"net/http"

	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/middleware"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/cookies"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

// NewContext is a convenient alias to context.New factory.
var NewContext = context.New

// NewHandler is a convenient alias to context.NewHandler factory.
var NewHandler = context.NewHandler

// NewMiddleware is a convenient alias to middleware.New factory.
var NewMiddleware = middleware.New

// Client represents a high-level HTTP client entity capable
// with a built-in middleware and context.
type Client struct {
	// Client entity can inherit behavior from a parent Client.
	Parent *Client

	// Each Client entity has it's own Context that will be inherited by requests or child clients.
	Context *context.Context

	// Client entity has its own Middleware layer to compose and inherit behavior.
	Middleware middleware.Middleware
}

// New creates a new high level client entity
// able to perform HTTP requests.
func New() *Client {
	return &Client{
		Context:    context.New(),
		Middleware: middleware.New(),
	}
}

// Request creates a new Request based on the current Client
func (c *Client) Request() *Request {
	req := NewRequest()
	req.SetClient(c)
	return req
}

// Options creates a new OPTIONS request.
func (c *Client) Options() *Request {
	req := c.Request()
	req.Method("OPTIONS")
	return req
}

// Get creates a new GET request.
func (c *Client) Get() *Request {
	req := c.Request()
	req.Method("GET")
	return req
}

// Post creates a new POST request.
func (c *Client) Post() *Request {
	req := c.Request()
	req.Method("POST")
	return req
}

// Put creates a new PUT request.
func (c *Client) Put() *Request {
	req := c.Request()
	req.Method("PUT")
	return req
}

// Delete creates a new DELETE request.
func (c *Client) Delete() *Request {
	req := c.Request()
	req.Method("DELETE")
	return req
}

// Patch creates a new PATCH request.
func (c *Client) Patch() *Request {
	req := c.Request()
	req.Method("PATCH")
	return req
}

// Head creates a new HEAD request.
func (c *Client) Head() *Request {
	req := c.Request()
	req.Method("HEAD")
	return req
}

// Method defines a the default HTTP method used by outgoing client requests.
//
// ⚠️ Method employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.Method()` instead.
func (c *Client) Method(name string) *Client {
	c.Middleware.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Method = name
		h.Next(ctx)
	})
	return c
}

// URL defines the URL for client requests.
// Useful to define at client level the base URL and base path used by child requests.
//
// ⚠️ URL employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.URL()` instead.
func (c *Client) URL(uri string) *Client {
	c.Use(url.URL(uri))
	return c
}

// BaseURL defines the URL schema and host for client requests.
// Useful to define at client level the base URL used by client child requests.
//
// ⚠️ BaseURL employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.BaseURL()` instead.
func (c *Client) BaseURL(uri string) *Client {
	c.Use(url.BaseURL(uri))
	return c
}

// Path defines the URL base path for client requests.
//
// ⚠️ Path employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.Path()` instead.
func (c *Client) Path(path string) *Client {
	c.Use(url.Path(path))
	return c
}

// AddPath concatenates a path slice to the existent path in at client level.
//
// ⚠️ AddPath employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.AddPath()` instead.
func (c *Client) AddPath(path string) *Client {
	c.Use(url.AddPath(path))
	return c
}

// Param replaces a path param based on the given param name and value.
//
// ⚠️ Param employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.Param()` instead.
func (c *Client) Param(name, value string) *Client {
	c.Use(url.Param(name, value))
	return c
}

// Params replaces path params based on the given params key-value map.
//
// ⚠️ Params employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.Params()` instead.
func (c *Client) Params(params map[string]string) *Client {
	c.Use(url.Params(params))
	return c
}

// SetHeader sets a new header field by name and value.
// If another header exists with the same key, it will be overwritten.
//
// ⚠️ SetHeader employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.SetHeader()` instead.
func (c *Client) SetHeader(key, value string) *Client {
	c.Use(headers.Set(key, value))
	return c
}

// AddHeader adds a new header field by name and value
// without overwriting any existent header.
//
// ⚠️ AddHeader employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.AddHeader()` instead.
func (c *Client) AddHeader(name, value string) *Client {
	c.Use(headers.Add(name, value))
	return c
}

// SetHeaders adds new header fields based on the given map.
//
// ⚠️ SetHeaders employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.SetHeaders()` instead.
func (c *Client) SetHeaders(fields map[string]string) *Client {
	c.Use(headers.SetMap(fields))
	return c
}

// AddCookie sets a new cookie field based on the given http.Cookie struct
// without overwriting any existent cookie.
//
// ⚠️ AddCookie employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.AddCookie()` instead.
func (c *Client) AddCookie(cookie *http.Cookie) *Client {
	c.Use(cookies.Add(cookie))
	return c
}

// AddCookies sets a new cookie field based on a list of http.Cookie
// without overwriting any existent cookie.
//
// ⚠️ AddCookies employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.AddCookies()` instead.
func (c *Client) AddCookies(data []*http.Cookie) *Client {
	c.Use(cookies.AddMultiple(data))
	return c
}

// CookieJar creates a cookie jar to store HTTP cookies when they are sent down.
//
// ⚠️ CookieJar employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.CookieJar()` instead.
func (c *Client) CookieJar() *Client {
	c.Use(cookies.Jar())
	return c
}

// UseContext adds a cancelation context to the client to enable the use of early cancelation. This is useful for
// server outgoing calls where we can attach the context from the incoming client. This will allow the downstream
// calls to be canceled early on the case of a tcp close or http2 cancellation.
// e.g.   client.Get().URL("someUrl").UseContext( incomingServerRequest.Context() ).Send()
func (c *Client) UseContext(cancelContext gocontext.Context) *Client {
	c.UseRequest(func(c *context.Context, handler context.Handler) {
		handler.Next(c.SetCancelContext(cancelContext))
	})
	return c
}

// Use uses a new plugin to the middleware stack.
//
// ⚠️ Use employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
//
// Use `Request.Use()` instead.
//
// Example:
// httpClient.Use(middleware).Post() // Middleware applies to ALL requests (⚠️ use with caution!)
// httpClient.Post().Use(middleware) // Middleware applies ONLY to this request (✅ recommended)
func (c *Client) Use(p plugin.Plugin) *Client {
	c.Middleware.Use(p)
	return c
}

// UseRequest uses a new middleware function for request phase.
//
// ⚠️ UseRequest employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.UseRequest()` instead.
func (c *Client) UseRequest(fn context.HandlerFunc) *Client {
	c.Middleware.UseRequest(fn)
	return c
}

// UseResponse uses a new middleware function for response phase.
//
// ⚠️ UseResponse employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.UseResponse()` instead.
func (c *Client) UseResponse(fn context.HandlerFunc) *Client {
	c.Middleware.UseResponse(fn)
	return c
}

// UseError uses a new middleware function for error phase.
//
// ⚠️ UseError employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.UseError()` instead.
func (c *Client) UseError(fn context.HandlerFunc) *Client {
	c.Middleware.UseError(fn)
	return c
}

// UseHandler uses a new middleware function for the given phase.
//
// ⚠️ UseHandler employs a new plugin within the middleware stack.
// Exercise caution when utilising this method. Considering its applicability to all requests, it may yield unforeseen consequences.
// Behaviours, such as mutex locks, may lead to complications if misused. Should you require middleware for a single request only?
// use `Request.UseHandler()` instead.
func (c *Client) UseHandler(phase string, fn context.HandlerFunc) *Client {
	c.Middleware.UseHandler(phase, fn)
	return c
}

// UseParent uses another Client as parent
// inheriting its middleware stack and configuration.
func (c *Client) UseParent(parent *Client) *Client {
	c.Parent = parent
	c.Context.UseParent(parent.Context)
	c.Middleware.UseParent(parent.Middleware)
	return c
}
