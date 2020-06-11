package middleware

import (
	"bytes"
	"fmt"
	Error "github.com/agitdevcenter/gopkg/error"
	"github.com/agitdevcenter/gopkg/json"
	Logger "github.com/agitdevcenter/gopkg/logger"
	Response "github.com/agitdevcenter/gopkg/response"
	Session "github.com/agitdevcenter/gopkg/session"
	Utils "github.com/agitdevcenter/gopkg/utils"
	ValueObject "github.com/agitdevcenter/gopkg/vo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

const (
	InternalServerErrorMessage = "Internal Server Error"
	Name                       = "LinkAja"
	Version                    = "1.0.0"
	RequestTime                = "RequestTime"
	RequestID                  = "RequestID"
	RequestError               = "RequestError"
	AlreadyLogged              = "AlreadyLogged"
	DebugURL                   = "/debug/pprof/*"
)

type Middleware struct {
	logger                     Logger.Logger
	debug                      bool
	port                       int
	name                       string
	version                    string
	recover                    bool
	cors                       bool
	gzip                       bool
	validator                  bool
	errorHandler               bool
	session                    bool
	acceptJSON                 bool
	internalServerErrorMessage string
	skipURLs                   []string
	health                     bool
	healthURL                  string
	profiling                  bool
	available                  bool
	availabilityEnabled        bool
	availabilityURLPrefix      string
	endpointAvailabilityURLs   []string
	endpointAvailabilityMap    map[string]bool
}

func New(opts []Option) *Middleware {
	m := &Middleware{
		port:                       80,
		name:                       Name,
		version:                    Version,
		internalServerErrorMessage: InternalServerErrorMessage,
		available:                  true,
		endpointAvailabilityMap:    make(map[string]bool),
	}

	for _, opt := range opts {
		opt(m)
	}

	if m.logger == nil {
		m.logger = Logger.Noop()
	}

	if m.profiling {
		m.skipURLs = append(m.skipURLs, DebugURL[0:len(DebugURL)-2])
	}

	if m.availabilityEnabled {
		m.skipURLs = append(m.skipURLs, m.availabilityURLPrefix)

		if len(m.endpointAvailabilityURLs) > 0 {
			for _, path := range m.endpointAvailabilityURLs {
				m.endpointAvailabilityMap[path] = true
			}
		}
	}

	return m
}

func (m *Middleware) SetPort(port int) {
	m.port = port
}

func (m *Middleware) SetLogger(logger Logger.Logger) {
	m.logger = logger
}

func (m *Middleware) SetDebug(enabled bool) {
	m.debug = enabled
}

func (m *Middleware) Setup(e *echo.Echo) {

	if m.profiling {
		e.GET(DebugURL, echo.WrapHandler(http.DefaultServeMux))
	}

	if m.health {
		e.GET(m.healthURL, func(c echo.Context) error {
			return m.healthHandler(c)
		})
	}

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Set(RequestTime, time.Now())

			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if len(reqId) == 0 {
				reqId = Utils.GenerateThreadId()
			}

			c.Set(RequestID, reqId)

			// check if server availability toggle is enabled
			if m.availabilityEnabled {
				// toggle availability for all server endpoint
				if c.Path() == m.availabilityURLPrefix {
					// negate the status
					m.available = !m.available

					if m.debug {
						m.logger.Info(fmt.Sprintf("availability status for HTTP server [%s %s] at [:%d] is %v", m.name, m.version, m.port, m.available))
					}

					response := make(map[string]bool)
					response["available"] = m.available

					return c.JSON(200, Response.DefaultResponse{
						Data: response,
						Response: Response.Response{
							Status:  Response.SuccessCode,
							Message: http.StatusText(http.StatusOK),
						},
					})
				}

				// check if we have configurations for specific url
				if len(m.endpointAvailabilityURLs) > 0 {
					// no need to check if the whole server is unavailable
					if m.available {
						// check only for url not in skipped list
						if !m.skip(c) {
							// go through the url list
							for _, path := range m.endpointAvailabilityURLs {
								// we go by prefix here
								if strings.HasPrefix(c.Path(), path) {
									// check if url is available
									var available bool
									var ok bool
									if available, ok = m.endpointAvailabilityMap[path]; !ok {
										available = false
									}
									// if not available return error
									if !available {
										return c.JSON(http.StatusServiceUnavailable, Response.DefaultResponse{
											Response: Response.Response{
												Status:  Response.GeneralError,
												Message: http.StatusText(http.StatusServiceUnavailable),
											},
										})
									}
								}
							}
						}
					}

					// toggle by url
					for _, path := range m.endpointAvailabilityURLs {
						// combine url toggle prefix with endpoint url
						if strings.HasPrefix(c.Path(), m.availabilityURLPrefix+path) {
							// check if available
							var available bool
							var ok bool
							if available, ok = m.endpointAvailabilityMap[path]; !ok {
								available = false
							}
							// negate the status
							m.endpointAvailabilityMap[path] = !available

							if m.debug {
								m.logger.Info(fmt.Sprintf("availability status for endpoint [%s] on HTTP server [%s %s] at [:%d] is %v", path, m.name, m.version, m.port, m.endpointAvailabilityMap[path]))
							}

							response := make(map[string]interface{})
							response["path"] = path
							response["available"] = m.endpointAvailabilityMap[path]

							return c.JSON(200, Response.DefaultResponse{
								Data: response,
								Response: Response.Response{
									Status:  Response.SuccessCode,
									Message: http.StatusText(http.StatusOK),
								},
							})
						}
					}
				}
			}

			// check if availability toggle is active and if it is unavailable and not in the skip list
			if m.availabilityEnabled && !m.available && !m.skip(c) {
				return c.JSON(http.StatusServiceUnavailable, Response.DefaultResponse{
					Response: Response.Response{
						Status:  Response.GeneralError,
						Message: http.StatusText(http.StatusServiceUnavailable),
					},
				})
			}

			if m.session {
				// - Set session to context
				request, errRequest := hookRequest(c)
				if errRequest != nil {
					c.Error(errRequest)
				}

				session := Session.New(m.logger).
					SetThreadID(reqId).
					SetAppName(m.name).
					SetAppVersion(m.version).
					SetPort(m.port).
					SetIP(c.Request().RemoteAddr).
					SetSrcIP(c.RealIP()).
					SetURL(c.Request().URL.String()).
					SetMethod(c.Request().Method).
					SetRequest(string(request)).
					SetHeader(formatHeader(c))

				if !m.skip(c) {
					session.T1("Incoming Request")
				}

				c.Set(ValueObject.AppSession, *session)
			}

			c.Response().Header().Set(echo.HeaderXRequestID, reqId)

			if m.acceptJSON {
				if !m.skip(c) {
					if c.Request().Header.Get(echo.HeaderContentType) != echo.MIMEApplicationJSON {
						return echo.NewHTTPError(http.StatusUnsupportedMediaType, http.StatusText(http.StatusUnsupportedMediaType))
					}
				}
			}

			return h(c)
		}
	})

	if m.gzip {
		e.Use(middleware.Gzip())
	}

	if m.recover {
		e.Use(middleware.Recover())
	}

	if m.cors {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.GET, echo.POST, echo.OPTIONS},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", "Pv", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: true,
		}))
	}

	if m.validator {
		e.Validator = &DataValidator{ValidatorData: validator.New()}
	}

	e.Use(middleware.BodyDump(func(c echo.Context, request []byte, response []byte) {
		var alreadyLogged bool
		var ok bool
		if alreadyLogged, ok = c.Get(AlreadyLogged).(bool); !ok {
			alreadyLogged = false
		}

		if !alreadyLogged {
			m.logRequest(c, request, response)
		}
	}))

	if m.errorHandler {
		e.HTTPErrorHandler = m.httpErrorHandler
	}

}

func (m *Middleware) logRequest(c echo.Context, request []byte, response []byte) {
	if m.health && strings.HasPrefix(c.Path(), m.healthURL) || m.skip(c) {
		return
	}

	if m.session {
		session := c.Get(ValueObject.AppSession).(Session.Session)

		var requestError error
		var ok bool
		if requestError, ok = c.Get(RequestError).(error); ok {
			session.SetErrorMessage(requestError.Error())
		}

		var resp map[string]interface{}
		json.Unmarshal(response, &resp)

		session.T4(resp)
	}

	if !m.session {
		var requestID string
		var ok bool
		if requestID, ok = c.Get(RequestID).(string); !ok {
			requestID = Utils.GenerateThreadId()
		}

		var requestTime time.Time
		if requestTime, ok = c.Get(RequestTime).(time.Time); !ok {
			requestTime = time.Now()
		}

		tdrModel := Logger.LogTdrModel{
			AppName:    m.name,
			AppVersion: m.version,
			IP:         c.Request().RemoteAddr,
			Port:       m.port,
			SrcIP:      c.RealIP(),
			RespTime:   time.Now().Sub(requestTime).Nanoseconds() / 1000000,
			Path:       c.Path(),
			Header:     c.Request().Header,
			Request:    string(request),
			Response:   string(response),
			ThreadID:   requestID,
		}

		var requestError error
		if requestError, ok = c.Get(RequestError).(error); ok {
			tdrModel.Error = requestError.Error()
		}

		m.logger.TDR(tdrModel)
	}
}

func (m *Middleware) httpErrorHandler(err error, c echo.Context) {
	c.Set(RequestError, err)

	var (
		code    = http.StatusInternalServerError
		message = http.StatusText(code)
	)

	response := Response.DefaultResponse{
		Response: Response.Response{
			Status:  Response.GeneralError,
			Message: message,
		},
		Data: struct{}{},
	}

	if he, ok := err.(*Error.ApplicationError); ok {
		response.Status = he.ErrorCode
		response.Message = he.Message
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		response.Message = he.Message.(string)
	} else {
		response.Message = err.Error()
	}

	if !c.Response().Committed {
		var responseError error
		if c.Request().Method == http.MethodHead { // Issue #608
			responseError = c.NoContent(code)
		} else {
			message := m.internalServerErrorMessage
			switch response.Message {
			case http.StatusText(http.StatusUnsupportedMediaType):
				message = response.Message
				code = http.StatusUnsupportedMediaType
				break
			case http.StatusText(http.StatusNotFound):
				message = response.Message
				code = http.StatusNotFound
				break
			case http.StatusText(http.StatusMethodNotAllowed):
				message = response.Message
				code = http.StatusMethodNotAllowed
				break
			}
			response.Message = message
			responseError = c.JSON(code, response)
		}

		var alreadyLogged bool
		var ok bool
		if alreadyLogged, ok = c.Get(AlreadyLogged).(bool); !ok {
			alreadyLogged = false
		}

		if !alreadyLogged {
			c.Set(AlreadyLogged, true)

			responseByte, _ := json.Marshal(response)
			requestByte, _ := hookRequest(c)

			m.logRequest(c, requestByte, responseByte)
		}

		err = responseError
	}
}

func (m *Middleware) healthHandler(c echo.Context) error {
	content := c.QueryParam("content")
	from := c.Get("RequestTime").(time.Time)
	data := make(map[string]interface{})
	data["name"] = m.name
	data["version"] = m.version
	data["elapsed"] = time.Now().Sub(from).Nanoseconds() / 1000000
	response := Response.CreateResponse(Response.SuccessCode, "healthy", data)
	if content == "true" {
		return c.JSON(http.StatusOK, response)
	}
	return c.NoContent(http.StatusNoContent)
}

func (m *Middleware) skip(c echo.Context) (skip bool) {
	for _, url := range m.skipURLs {
		if strings.HasPrefix(strings.ToLower(c.Request().URL.String()), url) {
			skip = true
			return
		}
	}
	return
}

func hookRequest(c echo.Context) (body []byte, err error) {
	if c.Request().Body != nil { // Read
		body, err = ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return body, err
		}
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, err
}

func formatHeader(c echo.Context) (result map[string]interface{}) {
	result = make(map[string]interface{})
	if headers := c.Request().Header; headers != nil {
		for k, v := range headers {
			result[k] = v
		}
		return result
	}
	return
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}
