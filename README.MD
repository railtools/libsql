# libsql on railway

A minimal example using [Turso/LibSQL](https://github.com/tursodatabase/libsql) with JWT-based authentication in Go.  
This template helps you:

- Generate a JWT keypair
- Create a JWT for accessing your database
- Connect to LibSQL using the `authToken`
- Run a basic query in `main.go`
- Runs it on railway!

---

## 🔧 Prerequisites

- Go 1.21+ must be installed: \
  [https://go.dev/dl/](https://go.dev/dl/)
- LibSQL database deployed: \
  [![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/libsql-jwt-auth?referralCode=dPr4mc)

## 🚀 Quickstart

### 1. Install dependencies

```bash
go mod tidy
```
### 2. Generate a new keypair
```bash
go run jwtgen.go --gen
```

Example output:
```bash
🔐 PRIVATE KEY (keep this safe):
uNvD4CzPyA2VTzbO1O4BAWTcpgtNrqD3hlu31gvGw6CwOscCRmSj2XkG8WGEAqJXScnREpiGVfJ0olKw-854Bg

📣 PUBLIC KEY (Railway → SQLD_AUTH_JWT_KEY):
sDrHAkZko9l5BvFhhAKiV0nJ0RKYhlXydKJSsPvOeAY
```

Store the private key in a password manager.

### 3. Update environment variable on RAILWAY
On railway project -> Variables -> SQLD_AUTH_JWT_KEY
`SQLD_AUTH_JWT_KEY`=`sDrHAkZko9l5BvFhhAKiV0nJ0RKYhlXydKJSsPvOeAY`

Update the value with the public key you generated.

### 4. Generate a JWT
Fill in the `🔐 PRIVATE KEY` from the previous step. And generate your JWT token.
876000h = 100 years.

```bash
go run jwtgen.go \
  --priv uNvD4CzPyA2VTzbO1O4BAWTcpgtNrqD3hlu31gvGw6CwOscCRmSj2XkG8WGEAqJXScnREpiGVfJ0olKw-854Bg \
  --user user \
  --exp 876000h
```

Result:
```bash
📈 JWT:
eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MDYwMTg5NTQsInN1YiI6InVzZXIifQ.GaOet8Jv_WOUYEv71oSp-ps0XfJAzAkQY34W3zxAcE70xFZa8KRGpkt4ql-7IDXBHwVWpckHnynMgthsBpghDg
```
Use this as your authToken.

### 5. Update `main.go`

```go
const authToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9..." // JWT token from previous step
const appDomain = "libsql-production-8db1.up.railway.app" // railway domain

dsn := fmt.Sprintf("libsql://%s?authToken=%s", appDomain, authToken)
```

### 6. Run example

```bash
go run main.go
```
Expected outcome:

```bash

User 1: Alice
User 2: Bob
```

Invalid keypair:
```bash
2025/07/13 17:07:02 drop failed: failed to execute SQL:
error code 401: Unauthorized: `The JWT is invalid`
exit status 1
```

✅ Troubleshooting:

- Did you update the Railway environment with the correct public key?
- Did you use the correct JWT in main.go?