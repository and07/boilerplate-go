//https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h
package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/and07/boilerplate-go/pkg/token"
)

type AuthInterceptor struct {
	extractor       token.Extractor
	jwtManager      *token.JWTManager
	accessibleRoles map[string][]string
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		claim, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "claim", claim)

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		_, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (*token.AccessTokenCustomClaims, error) {

	accessToken, err := interceptor.extractor.ExtractGRPC(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := interceptor.jwtManager.ValidateAccessToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return claims, nil
	}

	for _, role := range accessibleRoles {

		if role == "*" {
			return claims, nil
		}

		if Contains(claims.Roles, func(e string) bool { return role == e }) {
			return claims, nil
		}
	}

	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")

}

func Contains[T comparable](s []T, match func(e T) bool) bool {
	for _, v := range s {
		if match(v) {
			return true
		}
	}
	return false
}

func NewAuthInterceptor(extractor token.Extractor, jwtManager *token.JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{extractor, jwtManager, accessibleRoles}
}
