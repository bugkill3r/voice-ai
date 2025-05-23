package artifacts

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetToolProviders(t *testing.T) {
// 	providers, err := GetToolProviders()

// 	// Check if there's no error
// 	assert.NoError(t, err)

// 	// Check if providers slice is not empty
// 	assert.NotEmpty(t, providers)

// 	// Check if each provider has the expected fields
// 	for _, provider := range providers {
// 		assert.NotZero(t, provider.Id)
// 		assert.NotNil(t, provider.Provider)
// 		assert.NotEmpty(t, provider.Name)
// 		assert.NotEmpty(t, provider.Feature)
// 		assert.NotEmpty(t, provider.Actor)
// 		assert.NotEmpty(t, provider.ConnectConfiguration)
// 	}

// 	// Check for specific providers based on your YAML content
// 	// This is an example, adjust according to your actual data
// 	expectedProviders := map[string]struct {
// 		ID       uint64
// 		ToolID   uint64
// 		Provider struct {
// 			ID          uint64
// 			Name        string
// 			Description string
// 			Image       string
// 		}
// 		Name          string
// 		Feature       []string
// 		ConnectConfig map[string]string
// 		Actor         string
// 		Description   string
// 	}{
// 		"Confluence": {
// 			ID:     6707073499618514490,
// 			ToolID: 828163119872980643,
// 			Provider: struct {
// 				ID          uint64
// 				Name        string
// 				Description string
// 				Image       string
// 			}{
// 				ID:          828163119872980643,
// 				Name:        "Confluence",
// 				Description: "Collaborative workspace with templates, integration with Jira, and advanced search.",
// 				Image:       "https://cdn-01.rapida.ai/partners/tools/confluence-icon.svg",
// 			},
// 			Name:    "Confluence",
// 			Feature: []string{"connection", "data.knowledge"},
// 			ConnectConfig: map[string]string{
// 				"connect_url":  "/connect-knowledge/confluence",
// 				"connect_type": "oauth2",
// 			},
// 			Actor: "organization",
// 		},
// 		"Freshsales": {
// 			ID:     7675709218821439431,
// 			ToolID: 2625553142046091699,
// 			Provider: struct {
// 				ID          uint64
// 				Name        string
// 				Description string
// 				Image       string
// 			}{
// 				ID:          2625553142046091699,
// 				Name:        "Freshdesk",
// 				Description: "Integrated with ticketing system, AI-driven insights, customizable knowledge base, multilingual support.",
// 				Image:       "https://cdn-01.rapida.ai/partners/freshdesk.png",
// 			},
// 			Name:    "Freshsales",
// 			Feature: []string{"connection"},
// 			ConnectConfig: map[string]string{
// 				"form_input":   "",
// 				"connect_type": "form",
// 			},
// 			Actor:       "organization",
// 			Description: "Connection to manage all the request to freshsales",
// 		},
// 		"Zendesk": {
// 			ID:     7180679278431284956,
// 			ToolID: 6581214267333847816,
// 			Provider: struct {
// 				ID          uint64
// 				Name        string
// 				Description string
// 				Image       string
// 			}{
// 				ID:          6581214267333847816,
// 				Name:        "Zendesk",
// 				Description: "Self-service customer support platform enterprises.",
// 				Image:       "https://cdn-01.rapida.ai/partners/zendesk.jpg",
// 			},
// 			Name:    "Zendesk",
// 			Feature: []string{"connection"},
// 			ConnectConfig: map[string]string{
// 				"form_input":   "",
// 				"connect_type": "form",
// 			},
// 			Actor:       "organization",
// 			Description: "Connection to manage all the request to Zendesk",
// 		},
// 		"Twilio": {
// 			ID:     4413247467425994582,
// 			ToolID: 7835356314600149384,
// 			Provider: struct {
// 				ID          uint64
// 				Name        string
// 				Description string
// 				Image       string
// 			}{
// 				ID:          7835356314600149384,
// 				Name:        "Twilio",
// 				Description: "Cloud communications platform for building SMS, voice, and messaging applications",
// 				Image:       "https://cdn-01.rapida.ai/partners/tools/twilio-icon.svg",
// 			},
// 			Name:    "Twilio",
// 			Feature: []string{"connection", "action.call", "action.whatsapp", "action.sms"},
// 			ConnectConfig: map[string]string{
// 				"form_input": `[{"name": "messaging_service_sid", "type": "string", "label": "Messaging Service SID"},` +
// 					`{"name": "phone_number", "type": "string", "label": "Phone Number"}]`,
// 				"connect_type": "form",
// 			},
// 			Actor:       "organization",
// 			Description: "Setup and manage Twilio messaging service for SMS communication.",
// 		},
// 		"Microsoft Dynamics 365": {
// 			ID:     8408986260124700432,
// 			ToolID: 7646645603519189356,
// 			Provider: struct {
// 				ID          uint64
// 				Name        string
// 				Description string
// 				Image       string
// 			}{
// 				ID:          7646645603519189356,
// 				Name:        "Microsoft Dynamics 365",
// 				Description: "Integrated CRM and ERP platform for customer relationship management and business operations.",
// 				Image:       "https://cdn-01.rapida.ai/partners/tools/dynamics365-icon.webp",
// 			},
// 			Name:    "Microsoft Dynamics 365",
// 			Feature: []string{"connection"},
// 			ConnectConfig: map[string]string{
// 				"connect_url":  "/connect-crm/dynamics365",
// 				"connect_type": "oauth2",
// 			},
// 			Actor: "organization",
// 		},
// 	}

// 	for _, provider := range providers {
// 		expected, exists := expectedProviders[provider.Name]
// 		if exists {
// 			assert.Equal(t, expected.ID, provider.Id)
// 			assert.Equal(t, expected.Name, provider.Name)
// 			assert.Equal(t, expected.Actor, provider.Actor)

// 		}
// 	}
// }
