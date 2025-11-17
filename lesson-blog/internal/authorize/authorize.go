package authorize

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/henry-woo/golang-study/lesson-blog/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Tokens struct {
	AccessToken         string    `josn:"access_token"`
	RefreshToken        string    `josn:"refresh_token"`
	JTIAccessToken      string    `josn:"jit_access_token"`
	JTIRefreshToken     string    `josn:"jit_refresh_token"`
	ExpiredAccessToken  time.Time `josn:"expired_access_token"`
	ExpiredRefreshToken time.Time `josn:"expired_refresh_token"`
	UserID              string    `josn:"user_id"`
	Issuer              string    `josn:"issuer"`
	Audience            string    `josn:"audience"`
}

func IssueTokens(userID string) (*Tokens, error) {
	now := time.Now().UTC()
	t := &Tokens{
		UserID:              userID,
		JTIAccessToken:      uuid.NewString(),
		JTIRefreshToken:     uuid.NewString(),
		ExpiredAccessToken:  now.Add(15 * time.Minute),
		ExpiredRefreshToken: now.Add(7 * 24 * time.Hour),
		Issuer:              "jwt-blog-app",
		Audience:            "jwt-blog-client",
	}

	acc := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ID:        t.JTIAccessToken,
		Issuer:    t.Issuer,
		Audience:  jwt.ClaimStrings{t.Audience},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(t.ExpiredAccessToken),
	})

	ref := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ID:        t.JTIRefreshToken,
		Issuer:    t.Issuer,
		Audience:  jwt.ClaimStrings{t.Audience},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(t.ExpiredRefreshToken),
	})

	var err error
	t.AccessToken, err = acc.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	t.RefreshToken, err = ref.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return t, nil
}

func Persist(ctx context.Context, r *store.Redis, t *Tokens) error {
	if err := r.SetJTI(ctx, "access:"+t.JTIAccessToken, t.UserID, t.ExpiredAccessToken); err != nil {
		return err
	}

	if err := r.SetJTI(ctx, "refresh:"+t.JTIRefreshToken, t.UserID, t.ExpiredRefreshToken); err != nil {
		return err
	}

	return nil
}

func SetAuthCookies(c *gin.Context, t *Tokens) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", t.AccessToken, int(time.Until(t.ExpiredAccessToken).Seconds()), "/", "", true, true)
	c.SetCookie("refresh_token", t.RefreshToken, int(time.Until(t.ExpiredRefreshToken).Seconds()), "/", "", true, true)
}

func ClearAuthCookies(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
}

func ParseAccess(tokenStr string) (*jwt.RegisteredClaims, error) {
	secret := os.Getenv("ACCESS_SECRET")
	return ParseWithSecret(tokenStr, secret)
}

func ParseRefresh(tokenStr string) (*jwt.RegisteredClaims, error) {
	secret := os.Getenv("REFRESH_SECRET")
	return ParseWithSecret(tokenStr, secret)
}

func ParseWithSecret(tokenStr, secret string) (*jwt.RegisteredClaims, error) {
	if secret == "" {
		return nil, errors.New("jwt secret not configured")
	}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	token, err := parser.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Extra safety: ensure HMAC family
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
