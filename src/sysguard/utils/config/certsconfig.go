package config

import(
	"os"
	"sysguard/utils/print"
	"io/ioutil"
	"encoding/json"
	"crypto/tls"
	"fmt"
)

type Certs struct {
	Certs []Cert `json:"certs"`
}

type Cert struct {
	Certfile string `json:"certfile"`
	Keyfile string  `json:"keyfile"`
}

func Parse(path string) *tls.Config {
	tlsConf := &tls.Config{}
	jsonFile, err := os.Open(path)
	if err != nil {
		print.Critical("Can't read certs.json file")
		os.Exit(2)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var certs Certs
	
	json.Unmarshal(byteValue, &certs)

	tlsConf.Certificates = make([]tls.Certificate, len(certs.Certs))
	for i := 0; i < len(certs.Certs); i++ {
		tlsConf.Certificates[i], err = tls.LoadX509KeyPair(certs.Certs[i].Certfile, certs.Certs[i].Keyfile)
		if err != nil {
			print.Critical("Can't add certificate [ C: " + certs.Certs[i].Certfile + " | K: " + certs.Certs[i].Keyfile + " ] ")
			print.Critical("Exiting..")
			os.Exit(2)
		}
	}
	tlsConf.BuildNameToCertificate()
	fmt.Print(tlsConf.NameToCertificate)
	return tlsConf
}
