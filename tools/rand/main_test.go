package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"testing"

	compressecpoint "github.com/wodadehencou/compress-ec-point"
)

func Test_HH2I(t *testing.T) {
	curve := elliptic.P256()
	priv, _, _, _ := elliptic.GenerateKey(curve, rand.Reader)
	hx, hy := curve.ScalarBaseMult(priv)
	ix, iy := curve.ScalarMult(hx, hy, priv)
	compressedH := compressecpoint.Compress(curve, hx, hy)
	compressedI := compressecpoint.Compress(curve, ix, iy)
	base64H := base64.StdEncoding.EncodeToString(compressedH)
	base64I := base64.StdEncoding.EncodeToString(compressedI)
	t.Log(base64H)
	t.Log(base64I)
}
