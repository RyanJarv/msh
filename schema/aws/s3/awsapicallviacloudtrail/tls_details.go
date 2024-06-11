package awsapicallviacloudtrail

type TlsDetails struct {
    CipherSuite string `json:"cipherSuite"`
    ClientProvidedHostHeader string `json:"clientProvidedHostHeader"`
    TlsVersion string `json:"tlsVersion"`
}

func (t *TlsDetails) SetCipherSuite(cipherSuite string) {
    t.CipherSuite = cipherSuite
}

func (t *TlsDetails) SetClientProvidedHostHeader(clientProvidedHostHeader string) {
    t.ClientProvidedHostHeader = clientProvidedHostHeader
}

func (t *TlsDetails) SetTlsVersion(tlsVersion string) {
    t.TlsVersion = tlsVersion
}
