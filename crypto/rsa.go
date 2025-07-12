package crypto

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"

	"go-blog/errs"
	"go-blog/logger"

	"github.com/pkg/errors"
)

func RsaSign(ctx context.Context, prvkey []byte, hash crypto.Hash, data []byte) ([]byte, error) {
	block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.Errorf("decode private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logger.Errorf(ctx, "failed to parse pkc private key, data=%s", string(data))
		return nil, err
	}

	// MD5 and SHA1 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	rawSign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
	if err != nil {
		return nil, errs.EnsureStack(err)
	}

	//return []byte(base64.StdEncoding.EncodeToString(rawSign)), nil
	return rawSign, nil
}

func RsaVerifySign(ctx context.Context, pubkey []byte, hash crypto.Hash, data, sigArg []byte) error {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return errors.Errorf("decode public key error, pubkey=%s, sig=%s", string(pubkey), string(sigArg))
	}
	//pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		logger.Errorf(ctx, "failed to parse pkcs1 block, pubkey=%s, sig=%s", string(pubkey), string(sigArg))
		return errs.EnsureStack(err)
	}

	// SHA1 and MD5 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	//return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), hash, hashed, sig)
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sigArg)
}
