package integration

import (
	"github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type WhatsAppTemplateCategory string
type WhatsAppComponentType string
type WhatsAppButtonType string
type WhatsAppTemplateStatus string
type WhatsAppHeaderFormat string

const (
	WhatsAppTemplateCategoryTransactional WhatsAppTemplateCategory = "TRANSACTIONAL"
	WhatsAppTemplateCategoryMarketing     WhatsAppTemplateCategory = "MARKETING"
	WhatsAppTemplateTypeOTP               WhatsAppComponentType    = "OTP"

	WhatsAppComponentBody    WhatsAppComponentType = "BODY"
	WhatsAppComponentHeader  WhatsAppComponentType = "HEADER"
	WhatsAppComponentFooter  WhatsAppComponentType = "FOOTER"
	WhatsAppComponentButtons WhatsAppComponentType = "BUTTONS"

	WhatsAppButtonQuickReply  WhatsAppButtonType = "QUICK_REPLY"
	WhatsAppButtonPhoneNumber WhatsAppButtonType = "PHONE_NUMBER"
	WhatsAppButtonUrl         WhatsAppButtonType = "URL"

	WhatsAppTemplateStatusNew             WhatsAppTemplateStatus = "NEW"
	WhatsAppTemplateStatusPending         WhatsAppTemplateStatus = "PENDING"
	WhatsAppTemplateStatusApproved        WhatsAppTemplateStatus = "APPROVED"
	WhatsAppTemplateStatusRejected        WhatsAppTemplateStatus = "REJECTED"
	WhatsAppTemplateStatusPendingDeletion WhatsAppTemplateStatus = "PENDING_DELETION"
	WhatsAppTemplateStatusDeleted         WhatsAppTemplateStatus = "DELETED"

	WhatsAppHeaderFormatText     WhatsAppHeaderFormat = "TEXT"
	WhatsAppHeaderFormatImage    WhatsAppHeaderFormat = "IMAGE"
	WhatsAppHeaderFormatDocument WhatsAppHeaderFormat = "DOCUMENT"
	WhatsAppHeaderFormatVideo    WhatsAppHeaderFormat = "VIDEO"
)

type WhatsAppTemplateComponent struct {
	Type    WhatsAppComponentType              `json:"type"`
	Format  WhatsAppHeaderFormat               `json:"format,omitempty"`
	Text    string                             `json:"text,omitempty"`
	Buttons []*WhatsAppTemplateComponentButton `json:"buttons,omitempty"`
	Example *WhatsAppTemplateExample           `json:"example,omitempty"`
}

type WhatsAppTemplateComponentButton struct {
	Type        string   `json:"type"`
	Text        string   `json:"text"`
	Url         string   `json:"url,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	Example     []string `json:"example,omitempty"`
}

type WhatsAppTemplateExample struct {
	HeaderText []string   `json:"header_text,omitempty"`
	BodyText   [][]string `json:"body_text,omitempty"`
	HeaderUrl  []string   `json:"header_url,omitempty"`
}

type WhatsAppTemplate struct {
	Id             string                       `json:"id"`
	Name           string                       `json:"name"`
	Language       string                       `json:"language"`
	Category       WhatsAppTemplateCategory     `json:"category"`
	Components     []*WhatsAppTemplateComponent `json:"components"`
	Status         WhatsAppTemplateStatus       `json:"status"`
	RejectedReason string                       `json:"rejectedReason,omitempty"`
	WabaId         string                       `json:"wabaId"`
	Namespace      string                       `json:"namespace"`
	CreatedAt      time.Time                    `json:"createdAt"`
	UpdatedAt      time.Time                    `json:"updatedAt"`
}

type CreateWhatsAppTemplateRequest struct {
	Name       string                       `json:"name"`
	Language   string                       `json:"language"`
	Components []*WhatsAppTemplateComponent `json:"components"`
	Category   WhatsAppTemplateCategory     `json:"category"`
	WabaId     string                       `json:"wabaId,omitempty"`
}

type CreateWhatsAppTemplateResponse struct {
	Name       string                       `json:"name"`
	Language   string                       `json:"language"`
	Category   WhatsAppTemplateCategory     `json:"category"`
	Components []*WhatsAppTemplateComponent `json:"components"`
	Status     WhatsAppTemplateStatus       `json:"status"`
	WabaId     string                       `json:"wabaId"`
	Namespace  string                       `json:"namespace"`
}

type ListWhatsAppTemplatesRequest struct {
	Limit     int    `url:"limit,omitempty"`
	Offset    int    `url:"offset,omitempty"`
	WabaId    string `url:"wabaId,omitempty"`
	ChannelId string `url:"channelId,omitempty"`
}

type ListWhatsAppTemplatesResponse struct {
	Offset     int                 `json:"offset"`
	Limit      int                 `json:"limit"`
	Count      int                 `json:"count"`
	TotalCount int                 `json:"totalCount"`
	Items      []*WhatsAppTemplate `json:"items"`
}

type DeleteWhatsAppTemplateRequest struct {
	Name     string
	Language string
}

func CreateWhatsAppTemplate(c messagebird.Client, request *CreateWhatsAppTemplateRequest) (*CreateWhatsAppTemplateResponse, error) {
	response := &CreateWhatsAppTemplateResponse{}
	err := c.Request(response, http.MethodPost, apiRoot+"/"+version+"/"+path+"/"+whatsAppTemplatePath, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ListWhatsAppTemplates(c messagebird.Client, request *ListWhatsAppTemplatesRequest) (*ListWhatsAppTemplatesResponse, error) {
	response := &ListWhatsAppTemplatesResponse{}
	apiURL := apiRoot + "/v3/" + path + "/" + whatsAppTemplatePath
	if request != nil {
		queryParams := url.Values{}
		if request.Limit != 0 {
			queryParams.Add("limit", strconv.Itoa(request.Limit))
		}
		if request.Offset != 0 {
			queryParams.Add("offset", strconv.Itoa(request.Offset))
		}
		if request.WabaId != "" {
			queryParams.Add("wabaId", request.WabaId)
		}
		if request.ChannelId != "" {
			queryParams.Add("channelId", request.ChannelId)
		}
		apiURL += "?" + queryParams.Encode()
	}

	// Encode query parameters and append them to the URL
	err := c.Request(response, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DeleteWhatsAppTemplate(c messagebird.Client, request *DeleteWhatsAppTemplateRequest) error {
	p := request.Name
	if request.Language != "" {
		p += "/" + request.Language
	}
	return c.Request(nil, http.MethodDelete, apiRoot+"/"+version+"/"+path+"/"+whatsAppTemplatePath+"/"+p, nil)
}
