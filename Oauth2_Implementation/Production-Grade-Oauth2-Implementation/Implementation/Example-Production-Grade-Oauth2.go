/*
Here's a complete, simplified working example of a Golang OAuth 2.0 Authorization Server with JWT token generation, PKCE (Proof Key for Code Exchange), and Redis for token storage.

This implementation is a minimal version to get you started with OAuth 2.0 in a highly scalable environment. It includes PKCE validation, token generation, and storing tokens in Redis.

Note: This is a basic implementation to demonstrate the flow. You need to configure proper error handling, security practices, and production-grade optimizations before going live.

Step-by-Step Code
1. Install Dependencies
Install the necessary dependencies:

bash
Copy code
go get golang.org/x/oauth2
go get github.com/golang-jwt/jwt/v4
go get github.com/go-redis/redis/v8
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
2. Environment Setup
Create a .env file in your project directory to store sensitive information like client ID, secret, etc.

bash
Copy code
CLIENT_ID=your-client-id
CLIENT_SECRET=your-client-secret
REDIRECT_URL=http://localhost:8080/callback
OAUTH2_PROVIDER=http://localhost:8080
REDIS_ADDR=localhost:6379
3. Main OAuth Server Code (main.go)
Here's a working OAuth 2.0 Authorization Server with PKCE, JWT, and Redis.

go
Copy code
*/
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/go-redis/redis/v8"
// 	"github.com/joho/godotenv"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// )

// var (
// 	oauth2Config oauth2.Config
// 	redisClient   *redis.Client
// 	ctx           context.Context
// )

// func init() {
// 	// Load environment variables
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Initialize OAuth2 Config
// 	oauth2Config = oauth2.Config{
// 		ClientID:     "1095205211624-cf6oi6q9b3abdcfuib94skdhm6jopsu4.apps.googleusercontent.com",          // Replace with your actual OAuth Client ID
// 		ClientSecret: "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",      // Replace with your actual OAuth Client Secret
// 		RedirectURL:  "http://localhost:8080/callback",
// 		Scopes:       []string{"profile", "email"},
// 		Endpoint:     google.Endpoint,
// 	}

// 	// Initialize Redis client
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379", // Redis server address
// 	})

// 	// Initialize context
// 	ctx = context.Background()
// }

// func main() {
// 	http.HandleFunc("/", handleMain)
// 	http.HandleFunc("/login", handleLogin)
// 	http.HandleFunc("/callback", handleCallback)
// 	http.HandleFunc("/token", handleToken)
// 	http.HandleFunc("/logout", handleLogout)
// 	fmt.Println("Server running on port:8080",http.ListenAndServe(":8080", nil))
// }

// func handleMain(w http.ResponseWriter, r *http.Request){
// 	var htmlIndex=`<html>
// 	<body>
// 	<a href="/login"> Google Log in</a>
// 	<body>
// 	</html>`
// 	fmt.Fprintf(w, htmlIndex)
// }

// // handleLogin starts the OAuth flow (PKCE)
// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	// Generate a random code_verifier and code_challenge
// 	codeVerifier, codeChallenge := generatePKCE()

// 	// Store code_verifier in Redis for later validation
// 	redisClient.Set(ctx, codeVerifier, "true", 0)

// 	// Generate the authorization URL
// 	authURL := oauth2Config.AuthCodeURL("", oauth2.SetAuthURLParam("code_challenge", codeChallenge), oauth2.SetAuthURLParam("code_challenge_method", "S256"))
// 	http.Redirect(w, r, authURL, http.StatusFound)
// }

// // handleCallback is where the OAuth2 callback will be received with the authorization code
// func handleCallback(w http.ResponseWriter, r *http.Request) {
// 	code := r.URL.Query().Get("code")
// 	codeVerifier := r.URL.Query().Get("code_verifier")

// 	// Validate the PKCE code_verifier
// 	_, err := redisClient.Get(ctx, codeVerifier).Result()
// 	if err != nil {
// 		http.Error(w, "Invalid PKCE code_verifier", http.StatusUnauthorized)
// 		return
// 	}

// 	// Exchange the authorization code for an access token
// 	token, err := oauth2Config.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
// 	if err != nil {
// 		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Generate a JWT token using the access token
// 	jwtToken, err := generateJWT("my-secret", token.AccessToken)
// 	if err != nil {
// 		http.Error(w, "Failed to create JWT token", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the JWT token to the user
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"access_token": jwtToken,
// 	})
// }

// // handleToken is where clients can request access tokens by using refresh tokens (not implemented in this example).
// func handleToken(w http.ResponseWriter, r *http.Request) {
// 	// In a real system, you would check the refresh token here and return a new access token
// 	http.Error(w, "Token refresh not implemented", http.StatusNotImplemented)
// }

// // handleLogout invalidates the token by removing it from Redis
// func handleLogout(w http.ResponseWriter, r *http.Request) {
// 	// Assuming the client has a valid token, you would revoke the access token here
// 	token := r.URL.Query().Get("token")
// 	redisClient.Del(ctx, token)
// 	w.Write([]byte("Logged out successfully"))
// }

// // generateJWT generates a JWT token containing the access token as a claim
// func generateJWT(secret string, accessToken string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"access_token": accessToken,
// 		"exp":          time.Now().Add(time.Hour * 1).Unix(), // Expire in 1 hour
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secret))
// }

// // generatePKCE generates a code_verifier and code_challenge for PKCE
// func generatePKCE() (string, string) {
// 	// In production, use a secure random string generator for code_verifier
// 	codeVerifier := "random_code_verifier" // Should be securely generated
// 	codeChallenge := generateCodeChallenge(codeVerifier)
// 	return codeVerifier, codeChallenge
// }

// // generateCodeChallenge generates a base64url-encoded SHA256 hash of the code_verifier
// func generateCodeChallenge(codeVerifier string) string {
// 	// In production, implement proper base64url encoding of SHA256 hash of code_verifier
// 	return codeVerifier
// }

package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauth2Config oauth2.Config
	redisClient  *redis.Client
	ctx          context.Context
	jwtSecret    string
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load JWT secret
	jwtSecret = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is missing in .env")
	}

	// Initialize OAuth2 Config
	oauth2Config = oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes:       strings.Split(os.Getenv("OAUTH_SCOPES"), " "),
		Endpoint:     google.Endpoint,
	}

	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	ctx = context.Background()
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/logout", handleLogout)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	htmlIndex := `<html>
	<body>
	<a href="/login">Google Log in</a>
	</body>
	</html>`
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := randomString(32)
	codeVerifier := randomString(64)
	codeChallenge := generateCodeChallenge(codeVerifier)

	// Store code_verifier in Redis keyed by state
	redisClient.Set(ctx, state, codeVerifier, time.Minute*5)

	authURL := oauth2Config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
	http.Redirect(w, r, authURL, http.StatusFound)
}
func handleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")

	// Retrieve code_verifier from Redis
	codeVerifier, err := redisClient.Get(ctx, state).Result()
	if err != nil {
		http.Error(w, "Invalid or expired state", http.StatusUnauthorized)
		return
	}

	// Exchange authorization code for tokens
	token, err := oauth2Config.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch user info from Google
	client := oauth2Config.Client(ctx, token)
	resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
		return
	}

	// Generate JWT with user info inside
	claims := jwt.MapClaims{
		"access_token": token.AccessToken,
		"exp":          time.Now().Add(time.Hour).Unix(),
		"user":         userInfo, // Include user details
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJWT, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Failed to create JWT token", http.StatusInternalServerError)
		return
	}

	// Return JWT + user info in response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"jwt_token": signedJWT,
		"user":      userInfo,
	})
}

// func handleCallback(w http.ResponseWriter, r *http.Request) {
// 	state := r.URL.Query().Get("state")
// 	code := r.URL.Query().Get("code")

// 	// Retrieve code_verifier from Redis
// 	codeVerifier, err := redisClient.Get(ctx, state).Result()
// 	if err != nil {
// 		http.Error(w, "Invalid or expired state", http.StatusUnauthorized)
// 		return
// 	}

// 	// Exchange authorization code for tokens
// 	token, err := oauth2Config.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
// 	if err != nil {
// 		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Generate JWT
// 	jwtToken, err := generateJWT(jwtSecret, token.AccessToken)
// 	if err != nil {
// 		http.Error(w, "Failed to create JWT token", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"jwt_token": jwtToken,
// 	})
// }

func handleLogout(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != "" {
		redisClient.Del(ctx, token)
	}
	w.Write([]byte("Logged out successfully"))
}

func generateJWT(secret string, accessToken string) (string, error) {
	claims := jwt.MapClaims{
		"access_token": accessToken,
		"exp":          time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func randomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)[:length]
}

func generateCodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

/*
Breakdown of the Code:
OAuth2 Authorization Server:

The /login endpoint initiates the OAuth flow and redirects the user to the authorization provider (e.g., Google) with the PKCE challenge.

The /callback endpoint receives the authorization code from the OAuth provider and exchanges it for an access token.

A JWT token is generated and returned to the client.

PKCE:

The PKCE (Proof Key for Code Exchange) mechanism is implemented to improve security in public clients like mobile apps and single-page apps.

The code_verifier is stored in Redis for validation during token exchange.

Redis:

Redis is used to store and validate the PKCE code verifier and simulate token revocation.

JWT:

A JWT token is generated using the access token and a secret key. The JWT is returned to the client for subsequent requests.

4. Running the Server
To run the server:

Ensure Redis is running locally on port 6379 or adjust the Redis connection details.

Replace "your-client-id" and "your-client-secret" with your actual OAuth client credentials.

Run the server using:

bash
Copy code
go run main.go
Now, your server will be running on http://localhost:8080.

5. Testing the Flow
Login: Visit http://localhost:8080/login to start the OAuth 2.0 login flow.

Callback: After logging in, the callback will exchange the authorization code for an access token.

JWT Token: The server will respond with a JWT token that can be used for subsequent requests.

Logout: To invalidate the token, access the /logout endpoint and pass the token to be revoked.

6. Next Steps:
This is a minimal working example. In a production environment, you should:

Enhance security by generating a real, cryptographically secure code_verifier and code_challenge.

Store refresh tokens and support token renewal with refresh_token flow.

Implement error handling, logging, and monitoring.

Scale the system by deploying with Kubernetes and using auto-scaling with load balancing.
*/