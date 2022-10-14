package auth

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kokoichi206/go-expert/web/todo/clock"
	"github.com/kokoichi206/go-expert/web/todo/entity"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	RoleKey     = "role"
	UserNameKey = "user_name"
)

// 現在地からの相対パスでよさそ
//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {

	j := &JWTer{Store: s}

	// jwk.Key の形式に変換
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}

	// jwk.Key の形式に変換
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}

	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c

	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {

	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (j *JWTer) GenerateToken(ctx context.Context, u entity.User) ([]byte, error) {
	token, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		// これでよき？
		Issuer(`github.com/kokoichi206/go-expert/web/todo`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(RoleKey, u.Role).
		Claim(UserNameKey, u.Name).
		Build()

	if err != nil {
		return nil, fmt.Errorf("GetToken: failed to build token: %w", err)
	}
	if err := j.Store.Save(ctx, token.JwtID(), u.ID); err != nil {
		return nil, err
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}

	return signed, nil
}
