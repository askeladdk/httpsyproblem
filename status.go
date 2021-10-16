package httpsyproblem

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
//
// HTTP status codes can be served as errors:
//  httpsyproblem.Serve(w, r, httpsyproblem.StatusForbidden)
var (
	StatusContinue                      = Wrap(100, nil) // RFC 7231, 6.2.1
	StatusSwitchingProtocols            = Wrap(101, nil) // RFC 7231, 6.2.2
	StatusProcessing                    = Wrap(102, nil) // RFC 2518, 10.1
	StatusEarlyHints                    = Wrap(103, nil) // RFC 8297
	StatusOK                            = Wrap(200, nil) // RFC 7231, 6.3.1
	StatusCreated                       = Wrap(201, nil) // RFC 7231, 6.3.2
	StatusAccepted                      = Wrap(202, nil) // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo          = Wrap(203, nil) // RFC 7231, 6.3.4
	StatusNoContent                     = Wrap(204, nil) // RFC 7231, 6.3.5
	StatusResetContent                  = Wrap(205, nil) // RFC 7231, 6.3.6
	StatusPartialContent                = Wrap(206, nil) // RFC 7233, 4.1
	StatusMultiStatus                   = Wrap(207, nil) // RFC 4918, 11.1
	StatusAlreadyReported               = Wrap(208, nil) // RFC 5842, 7.1
	StatusIMUsed                        = Wrap(226, nil) // RFC 3229, 10.4.1
	StatusMultipleChoices               = Wrap(300, nil) // RFC 7231, 6.4.1
	StatusMovedPermanently              = Wrap(301, nil) // RFC 7231, 6.4.2
	StatusFound                         = Wrap(302, nil) // RFC 7231, 6.4.3
	StatusSeeOther                      = Wrap(303, nil) // RFC 7231, 6.4.4
	StatusNotModified                   = Wrap(304, nil) // RFC 7232, 4.1
	StatusUseProxy                      = Wrap(305, nil) // RFC 7231, 6.4.5
	StatusTemporaryRedirect             = Wrap(307, nil) // RFC 7231, 6.4.7
	StatusPermanentRedirect             = Wrap(308, nil) // RFC 7538, 3
	StatusBadRequest                    = Wrap(400, nil) // RFC 7231, 6.5.1
	StatusUnauthorized                  = Wrap(401, nil) // RFC 7235, 3.1
	StatusPaymentRequired               = Wrap(402, nil) // RFC 7231, 6.5.2
	StatusForbidden                     = Wrap(403, nil) // RFC 7231, 6.5.3
	StatusNotFound                      = Wrap(404, nil) // RFC 7231, 6.5.4
	StatusMethodNotAllowed              = Wrap(405, nil) // RFC 7231, 6.5.5
	StatusNotAcceptable                 = Wrap(406, nil) // RFC 7231, 6.5.6
	StatusProxyAuthRequired             = Wrap(407, nil) // RFC 7235, 3.2
	StatusRequestTimeout                = Wrap(408, nil) // RFC 7231, 6.5.7
	StatusConflict                      = Wrap(409, nil) // RFC 7231, 6.5.8
	StatusGone                          = Wrap(410, nil) // RFC 7231, 6.5.9
	StatusLengthRequired                = Wrap(411, nil) // RFC 7231, 6.5.10
	StatusPreconditionFailed            = Wrap(412, nil) // RFC 7232, 4.2
	StatusRequestEntityTooLarge         = Wrap(413, nil) // RFC 7231, 6.5.11
	StatusRequestURITooLong             = Wrap(414, nil) // RFC 7231, 6.5.12
	StatusUnsupportedMediaType          = Wrap(415, nil) // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable  = Wrap(416, nil) // RFC 7233, 4.4
	StatusExpectationFailed             = Wrap(417, nil) // RFC 7231, 6.5.14
	StatusTeapot                        = Wrap(418, nil) // RFC 7168, 2.3.3
	StatusMisdirectedRequest            = Wrap(421, nil) // RFC 7540, 9.1.2
	StatusUnprocessableEntity           = Wrap(422, nil) // RFC 4918, 11.2
	StatusLocked                        = Wrap(423, nil) // RFC 4918, 11.3
	StatusFailedDependency              = Wrap(424, nil) // RFC 4918, 11.4
	StatusTooEarly                      = Wrap(425, nil) // RFC 8470, 5.2.
	StatusUpgradeRequired               = Wrap(426, nil) // RFC 7231, 6.5.15
	StatusPreconditionRequired          = Wrap(428, nil) // RFC 6585, 3
	StatusTooManyRequests               = Wrap(429, nil) // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge   = Wrap(431, nil) // RFC 6585, 5
	StatusUnavailableForLegalReasons    = Wrap(451, nil) // RFC 7725, 3
	StatusInternalServerError           = Wrap(500, nil) // RFC 7231, 6.6.1
	StatusNotImplemented                = Wrap(501, nil) // RFC 7231, 6.6.2
	StatusBadGateway                    = Wrap(502, nil) // RFC 7231, 6.6.3
	StatusServiceUnavailable            = Wrap(503, nil) // RFC 7231, 6.6.4
	StatusGatewayTimeout                = Wrap(504, nil) // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = Wrap(505, nil) // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = Wrap(506, nil) // RFC 2295, 8.1
	StatusInsufficientStorage           = Wrap(507, nil) // RFC 4918, 11.5
	StatusLoopDetected                  = Wrap(508, nil) // RFC 5842, 7.2
	StatusNotExtended                   = Wrap(510, nil) // RFC 2774, 7
	StatusNetworkAuthenticationRequired = Wrap(511, nil) // RFC 6585, 6
)
