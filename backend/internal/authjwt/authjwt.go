// Package authjwt verifies the JWTs issued by services/auth-service's
// BetterAuth jwt() plugin against its published JWKS. BetterAuth signs with
// EdDSA by default and embeds the full user record — including the role
// field added by the admin() plugin — in the payload, so a verified token's
// claims are enough to authorize an admin request without a DB round trip.
package authjwt

import (
	"context"
	"fmt"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type Verifier struct {
	keyfunc keyfunc.Keyfunc
	issuer  string
}

// NewVerifier fetches the JWKS at jwksURL and keeps it refreshed for the
// lifetime of ctx. issuer is BetterAuth's BASE_URL — by default it signs
// every JWT with iss and aud both set to that URL, so pinning both here
// stops a token minted for some other purpose (BetterAuth also issues JWTs
// for its oauth-provider/MCP plugins) from being accepted just because it
// carries a valid signature from the same key.
func NewVerifier(ctx context.Context, jwksURL, issuer string) (*Verifier, error) {
	k, err := keyfunc.NewDefaultCtx(ctx, []string{jwksURL})
	if err != nil {
		return nil, fmt.Errorf("authjwt: fetch jwks: %w", err)
	}
	return &Verifier{keyfunc: k, issuer: issuer}, nil
}

// Verify checks the token's signature, standard claims, and that its alg is
// EdDSA — restricting the accepted algorithm prevents an alg-confusion
// attack where a caller supplies a token signed with a different method.
func (v *Verifier) Verify(tokenString string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, v.keyfunc.Keyfunc,
		jwt.WithValidMethods([]string{"EdDSA"}),
		jwt.WithIssuer(v.issuer),
		jwt.WithAudience(v.issuer),
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("authjwt: invalid token")
	}
	return &claims, nil
}
