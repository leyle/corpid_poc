package poc

type DataType string

// fabric method names

var (
	MethodCreate  = "Create"
	MethodGetById = "GetById"
	MethodGetAll  = "GetAll"
)

var (
	CredentialDataType    DataType = "credentialData"
	AuthenticationDIDType DataType = "authenticationDID"
	CredentialDIDType     DataType = "credentialDID"
)

type LedgerState struct {
	Id         string   `json:"id"`
	Type       DataType `json:"type"`
	DataString string   `json:"data"`
}

type CredentialDataSchema struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Object    string    `json:"object"`
	CreatedAt string    `json:"createdAt"`
	Schemas   []*Schema `json:"schema"`
}

type Schema struct {
	Key        string `json:"key"`
	DataType   string `json:"datatype"`
	IsCredAttr bool   `json:"isCredAttr"`
}

type AuthenticationDID struct {
	Id        string `json:"id"`
	AtContext string `json:"@context"`
	Type      string `json:"type"`
	Created   string `json:"created"`
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

type CredentialDID struct {
	Id               string   `json:"id"`
	AtContext        string   `json:"@context"`
	Type             []string `json:"type"`
	Created          string   `json:"created"`
	CredentialStatus string   `json:"credentialStatus"`
	Hash             string   `json:"hash"`
	HolderSignature  string   `json:"holderSignature"`
	IssuerSignature  string   `json:"issuerSignature"`
}
