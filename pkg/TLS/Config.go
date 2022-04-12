package TLS

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func GetConfig() *tls.Config {
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("certs/ca/ca.crt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	certPool.AppendCertsFromPEM(ca)

	certificateKeyPair, certReadingErr := tls.LoadX509KeyPair("certs/client/client.crt", "certs/client/client.key")

	if certReadingErr != nil {
		panic(certReadingErr)
	}

	return &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		Certificates:       []tls.Certificate{certificateKeyPair},
	}
}
