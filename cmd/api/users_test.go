package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/Damisicode/social/internal/store/cache"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

var testClaims = jwt.MapClaims{
	"aud": "test-aud",
	"iss": "test-aud",
	"sub": int64(42),
	"exp": time.Now().Add(time.Hour).Unix(),
}

func TestGetUser(t *testing.T) {
	withRedis := config{
		redisCfg: redisConfig{
			enabled: true,
		},
	}

	app := newTestApplication(t, withRedis)
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(testClaims)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(42)).Return(nil, nil).Twice()
		mockCacheStore.On("Set", mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/42", nil)

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("should hit the cache first and if not exists it sets the user on the cach", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Get", int64(42)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)

		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("should NOT hit the cache if it is not enabled", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisConfig{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()

		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNotCalled(t, "Get")

		mockCacheStore.Calls = nil // Reset mock expectations
	})
}
