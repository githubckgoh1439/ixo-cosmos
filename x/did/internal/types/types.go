package types

import (
	"errors"
	
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

var _ ixo.DidDoc = (*BaseDidDoc)(nil)

type BaseDidDoc struct {
	Did         ixo.Did         `json:"did"`
	PubKey      string          `json:"pubKey"`
	Credentials []DidCredential `json:"credentials"`
}

type DidCredential struct {
	CredType []string `json:"type"`
	Issuer   ixo.Did  `json:"issuer"`
	Issued   string   `json:"issued"`
	Claim    Claim    `json:"claim"`
}

type Claim struct {
	Id           ixo.Did `json:"id"`
	KYCValidated bool    `json:"KYCValidated"`
}

type Credential struct{}

func (dd BaseDidDoc) GetDid() ixo.Did                 { return dd.Did }
func (dd BaseDidDoc) GetPubKey() string               { return dd.PubKey }
func (dd BaseDidDoc) GetCredentials() []DidCredential { return dd.Credentials }

func InitDidDoc(did ixo.Did, pubKey string) BaseDidDoc {
	return BaseDidDoc{
		did,
		pubKey,
		make([]DidCredential, 0),
	}
}

func (dd BaseDidDoc) SetDid(did ixo.Did) error {
	if len(dd.Did) != 0 {
		return errors.New("cannot override BaseDidDoc did")
	}
	
	dd.Did = did
	
	return nil
}

func (dd BaseDidDoc) SetPubKey(pubKey string) error {
	if len(dd.PubKey) != 0 {
		return errors.New("cannot override BaseDidDoc pubKey")
	}
	
	dd.PubKey = pubKey
	
	return nil
}

func (dd *BaseDidDoc) AddCredential(cred DidCredential) {
	if dd.Credentials == nil {
		dd.Credentials = make([]DidCredential, 0)
	}
	
	dd.Credentials = append(dd.Credentials, cred)
}

type DidMsg interface {
	IsNewDid() bool
}
