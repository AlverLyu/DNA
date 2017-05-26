package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func ParseCert(certData []byte) (*x509.Certificate, error) {
	p, _ := pem.Decode(certData)
	if p != nil { // the certificate is PEM format
		return x509.ParseCertificate(p.Bytes)
	} else { // the certificate is DER format
		return x509.ParseCertificate(certData)
	}
}

// Open the certificate file which path is @certPath
// The file can be either PEM or DER format
func OpenCert(certPath string) (*x509.Certificate, error) {
	buf, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	return ParseCert(buf)
}

// Verify whether the signature is signed with the certificate
// @certPath:  path of the certificate file
// @alg:       algorithm used for signing
// @signed:    the signed data
// @signature: signature to be verified
//
// return nil if verification succeeded, otherwise error
func VerifySignature(certPath string, alg x509.SignatureAlgorithm, signed, signature []byte) error {
	cert, err := OpenCert(certPath)
	if err != nil {
		return err
	}

	return cert.CheckSignature(alg, signed, signature)
}
