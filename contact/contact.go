package contact

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	messagebird "github.com/messagebird/go-rest-api/v9"
)

// path represents the path to the Contacts resource.
const path = "https://contacts.messagebird.com/v2/contacts"

// Contact gets returned by the API.
type Contact struct {
	ID              string
	HRef            string
	DisplayName     string
	FirstName       string
	LastName        string
	Identifiers     []Identifier
	Languages       []string
	Timezone        string
	Country         string
	Avatar          string
	Gender          string
	Profiles        []Profile
	Attributes      map[string]string
	Status          string
	CreatedDatetime *time.Time
	UpdatedDatetime *time.Time
}

type Contacts struct {
	Limit, Offset     int
	Count, TotalCount int
	Items             []Contact
}

type Profile struct {
	ID         string `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Platform   string `json:"platform,omitempty"`
	ChannelId  string `json:"channelId,omitempty"`
}

// CreateRequest represents a contact for write operations, e.g. for creating a new
// contact or updating an existing one.
type CreateRequest struct {
	DisplayName string            `json:"displayName,omitempty"`
	FirstName   string            `json:"firstName,omitempty"`
	LastName    string            `json:"lastName,omitempty"`
	Avatar      string            `json:"avatar,omitempty"`
	Gender      string            `json:"gender,omitempty"`
	Country     string            `json:"country,omitempty"`
	Languages   []string          `json:"languages,omitempty"`
	Identifiers []Identifier      `json:"identifiers,omitempty"`
	Profiles    []Profile         `json:"profiles,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

// Filter can be applied to a list request to filter the results and paginate through them.
type Filter struct {
	IDs             []string
	ChannelId       string
	IdentifierExact string
	Status          string
	Offset          int
	Limit           int
}

func (f *Filter) QueryParams() string {
	if f == nil {
		return ""
	}

	query := url.Values{}
	if f.Limit > 0 {
		query.Set("limit", strconv.Itoa(f.Limit))
	}
	if f.Offset >= 0 {
		query.Set("offset", strconv.Itoa(f.Offset))
	}
	if f.IDs != nil && len(f.IDs) > 0 {
		query.Set("ids", strings.Join(f.IDs, ","))
	}
	if f.ChannelId != "" {
		query.Set("channelId", f.ChannelId)
	}
	if f.IdentifierExact != "" {
		query.Set("identifierExact", f.IdentifierExact)
	}
	if f.Status != "" {
		query.Set("status", f.Status)
	}

	return query.Encode()
}

func Create(c messagebird.Client, contactRequest *CreateRequest) (*Contact, error) {
	contact := &Contact{}
	if err := c.Request(contact, http.MethodPost, path, contactRequest); err != nil {
		return nil, err
	}

	return contact, nil
}

// Delete attempts deleting the contact with the provided ID. If nil is returned,
// the resource was deleted successfully.
func Delete(c messagebird.Client, id string) error {
	return c.Request(nil, http.MethodDelete, path+"/"+id, nil)
}

// List retrieves a paginated and filtered list of contacts, based on the options provided.
// It's worth noting DefaultListOptions.
func List(c messagebird.Client, options *Filter) (*Contacts, error) {
	contactList := &Contacts{}
	if err := c.Request(contactList, http.MethodGet, path+"?"+options.QueryParams(), nil); err != nil {
		return nil, err
	}

	return contactList, nil
}

// Read retrieves the information of an existing contact.
func Read(c messagebird.Client, id string) (*Contact, error) {
	contact := &Contact{}
	if err := c.Request(contact, http.MethodGet, path+"/"+id, nil); err != nil {
		return nil, err
	}

	return contact, nil
}

// Update updates the record referenced by id with any values set in contactRequest.
// Do not set any values that should not be updated.
func Update(c messagebird.Client, id string, contactRequest *CreateRequest) (*Contact, error) {
	contact := &Contact{}
	if err := c.Request(contact, http.MethodPatch, path+"/"+id, contactRequest); err != nil {
		return nil, err
	}

	return contact, nil
}
