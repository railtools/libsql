package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	user := flag.String("user", "anonymous", "Subject (user) in the JWT")
	expStr := flag.String("exp", "10y", "Expiration duration (e.g. 10m, 2h)")
	genKey := flag.Bool("gen", false, "Generate and show new key pair")
	privKey := flag.String("priv", "", "Private key (base64url) to sign JWT")
	flag.Parse()

	var priv ed25519.PrivateKey
	var pub ed25519.PublicKey

	if *genKey {
		pub, priv, _ = ed25519.GenerateKey(nil)
		fmt.Println("🔐 PRIVATE KEY (keep this safe):")
		fmt.Println(base64.RawURLEncoding.EncodeToString(priv))
		fmt.Println("📣 PUBLIC KEY (Railway → SQLD_AUTH_JWT_KEY):")
		fmt.Println(base64.RawURLEncoding.EncodeToString(pub))
		return
	}

	if *privKey == "" {
		fmt.Fprintln(
			os.Stderr, "Missing --priv (private key in base64url). Use --gen to generate one.",
		)
		os.Exit(1)
	}

	privBytes, err := base64.RawURLEncoding.DecodeString(*privKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid private key:", err)
		os.Exit(1)
	}
	priv = ed25519.PrivateKey(privBytes)

	expDur, err := time.ParseDuration(*expStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid expiration duration:", err)
		os.Exit(1)
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodEdDSA, jwt.MapClaims{
			"sub": *user,
			"exp": time.Now().Add(expDur).Unix(),
		},
	)

	signed, err := token.SignedString(priv)
	if err != nil {
		fmt.Fprintln(os.Stderr, "JWT signing failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ JWT:")
	fmt.Println(signed)
}
