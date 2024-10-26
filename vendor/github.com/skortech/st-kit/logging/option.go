package logging

import "github.com/skortech/st-kit/service"

type Option func(*option)

// Options holds the DB configuration
type option struct {
	reference   map[string]interface{}
	identity    string
	service     service.Service
	errors      []error
	appVersion  int
	prettyPrint bool
	stackTrace  bool
	referenceID string
	deviceID    string
	ip          string
	skorTransID string
	platform    string // ANDROID, IOS
	skipLog     bool
}

func WithReference(ref map[string]interface{}) func(*option) {
	return func(h *option) {
		h.reference = ref
	}
}

func WithIdentity(userID string) func(*option) {
	return func(h *option) {
		h.identity = userID
	}
}

func WithService(service service.Service) func(*option) {
	return func(h *option) {
		h.service = service
	}
}

// WithTraceError keeps the original error for back tracing
func WithTraceError(err error) func(*option) {
	return func(o *option) {
		if err != nil {
			o.errors = append(o.errors, err)
		}
	}
}

func WithStackTrace(enable bool) func(*option) {
	return func(o *option) {
		o.stackTrace = enable
	}
}

func WithAppVersion(code int) func(*option) {
	return func(o *option) {
		o.appVersion = code
	}
}

func WithPrettyPrint(enable bool) func(*option) {
	return func(o *option) {
		o.prettyPrint = enable
	}
}

func WithReferenceID(referenceID string) func(*option) {
	return func(h *option) {
		h.referenceID = referenceID
	}
}

func WithDeviceID(id string) func(*option) {
	return func(h *option) {
		h.deviceID = id
	}
}

func WithPlatform(platform string) func(*option) {
	return func(h *option) {
		h.platform = platform
	}
}

func WithIP(ip string) func(*option) {
	return func(h *option) {
		h.ip = ip
	}
}

func WithSkorTransID(id string) func(*option) {
	return func(h *option) {
		h.skorTransID = id
	}
}

func WithSkip(skip bool) func(*option) {
	return func(h *option) {
		h.skipLog = skip
	}
}
