package consts

import "time"

type contextKey struct {
	name string
}

func (c *contextKey) String() string { return "context value " + c.name }

var (
	// AuthenticatedPrincipalContextKey is context key that contains the principal object
	AuthenticatedPrincipalContextKey = contextKey{name: "auth-principal"}
)

const (
	// const keys used to pass values from gRPC methods to http mux interface
	SessionIDGatewayKey     = "gateway-session-id"
	SessionDeleteGatewayKey = "gateway-session-delete"
	UserTokenGatewayKey     = "gateway-user-token"
	LocationGatewayKey      = "gateway-location"
	UserSecretGatewayKey    = "gateway-user-secret"

	// UserTokenRequestKey is returned from the application to client containing user details in
	// response headers
	UserTokenRequestKey = "x-user-token"

	// LocationRequestKey is used to set location response header for redirecting browser
	LocationRequestKey = "location"

	// SessionRequestKey is the key to store session value in browser
	SessionRequestKey = "sid"
	// SessionValidity defines the age of a session
	// TODO(kushsharma): should we expose this in config?
	SessionValidity = time.Hour * 24 * 30 // 30 days
)