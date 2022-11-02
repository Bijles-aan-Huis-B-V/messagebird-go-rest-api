package contact

import (
	messagebird "github.com/messagebird/go-rest-api/v9"
	"net/http"
	"time"
)

// pathIdentifier represents the path to the Identifier resource of a Contact.
const pathIdentifier = "/identifiers"

type Identifier struct {
	ID        string    `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	Value     string    `json:"value,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type Identifiers struct {
	Count, TotalCount int
	Offset, Limit     int
	Items             []Identifier
}

func ListIdentifiers(c messagebird.Client, contact string) (*Identifiers, error) {
	identifiers := &Identifiers{}
	if err := c.Request(identifiers, http.MethodGet, path+"/"+contact+pathIdentifier, nil); err != nil {
		return nil, err
	}
	return identifiers, nil
}

func ReadIdentifier(c messagebird.Client, contact, id string) (*Identifier, error) {
	identifier := &Identifier{}
	if err := c.Request(identifier, http.MethodGet, path+"/"+contact+pathIdentifier+"/"+id, nil); err != nil {
		return nil, err
	}
	return identifier, nil
}

//func CreateIdentifier(c messagebird.Client, contact string, identifier *Identifier) (*Identifier, error) {
//
//}
