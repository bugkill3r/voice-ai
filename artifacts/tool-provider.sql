CREATE TABLE tool_providers (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    image TEXT,
    feature JSONB,
    connect_configuration JSONB
);

INSERT INTO tool_providers (id, name, description, image, feature, connect_configuration) VALUES
(828163119872980643, 'Confluence', 'Collaborative workspace with templates, integration with Jira, and advanced search.', 'https://cdn-01.rapida.ai/partners/tools/confluence-icon.svg', 
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/confluence", "connect_type": "oauth2"}'::jsonb),

(8962189241232892279, 'Notion', 'Flexible workspace, database functionality, rich media support, collaboration tools.', 'https://cdn-01.rapida.ai/partners/tools/notion-icon.svg',
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/notion", "connect_type": "oauth2"}'::jsonb),

(1152680016786286407, 'SharePoint', 'Document management, intranet capabilities, integration with Office 365, collaboration tools.', 'https://cdn-01.rapida.ai/partners/tools/sharepoint-icon.svg',
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/sharepoint", "connect_type": "oauth2"}'::jsonb),

(7126162197231234774, 'Google Drive', 'Cloud storage service with real-time collaboration, sharing capabilities, and integration with Google Workspace tools.', 'https://cdn-01.rapida.ai/partners/tools/google-drive-icon.svg',
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/google-drive", "connect_type": "oauth2"}'::jsonb),

(4373891332059251235, 'GitHub', 'Web-based platform for version control using Git, with additional collaboration features.', 'https://cdn-01.rapida.ai/partners/tools/github-icon.svg',
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/github", "connect_type": "oauth2"}'::jsonb),

(6032054350093361929, 'Microsoft OneDrive', 'Cloud storage service for file hosting and synchronization from Microsoft.', 'https://cdn-01.rapida.ai/partners/tools/one-drive-icon.svg',
 '["connection", "data.knowledge"]'::jsonb,
 '{"connect_url": "/connect-knowledge/microsoft-onedrive", "connect_type": "oauth2"}'::jsonb),

(412702929018243296, 'Slack', 'Collaboration hub that connects people with the information they need.', 'https://cdn-01.rapida.ai/partners/tools/slack-icon.svg',
 '["connection", "action.send_message"]'::jsonb,
 '{"connect_url": "/connect-action/slack", "connect_type": "oauth2"}'::jsonb),

(7646645603519189352, 'HubSpot', 'CRM platform for marketing, sales, and customer service. Provides tools for inbound marketing, sales automation, and analytics.', 'https://cdn-01.rapida.ai/partners/tools/hubspot-icon.png',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/hubspot", "connect_type": "oauth2"}'::jsonb),

(7646645603519189353, 'Vonage', 'Cloud communication platform that provides voice, messaging, and video solutions for businesses.', 'https://cdn-01.rapida.ai/partners/vonage.jpeg',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-communication/vonage", "connect_type": "oauth2"}'::jsonb),

(7646645603519189354, 'Exotel', 'Cloud telephony platform offering voice and messaging services for businesses.', 'https://cdn-01.rapida.ai/partners/exotel.jpeg',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-communication/exotel", "connect_type": "oauth2"}'::jsonb),

(2625553142046091699, 'Freshdesk', 'Integrated with ticketing system, AI-driven insights, customizable knowledge base, multilingual support.', 'https://cdn-01.rapida.ai/partners/freshdesk.png',
 '["connection"]'::jsonb,
 '{"form_input": "", "connect_type": "form"}'::jsonb),

(6581214267333847816, 'Zendesk', 'Self-service customer support platform enterprises.', 'https://cdn-01.rapida.ai/partners/zendesk.jpg',
 '["connection"]'::jsonb,
 '{"form_input": "", "connect_type": "form"}'::jsonb),

(7835356314600149384, 'Twilio', 'Cloud communications platform for building SMS, voice, and messaging applications', 'https://cdn-01.rapida.ai/partners/tools/twilio-icon.svg',
 '["connection", "action.call", "action.whatsapp", "action.sms"]'::jsonb,
 '{
   "form_input": [
      {"name": "messaging_service_sid", "type": "string", "label": "Messaging Service SID"},
      {"name": "phone_number", "type": "string", "label": "Phone Number"}
   ],
   "connect_type": "form"
 }'::jsonb),

(7646645603519189356, 'Microsoft Dynamics 365', 'Integrated CRM and ERP platform for customer relationship management and business operations.', 'https://cdn-01.rapida.ai/partners/tools/dynamics365-icon.webp',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/dynamics365", "connect_type": "oauth2"}'::jsonb),

(7646645603519189366, 'Zoho CRM', 'Cloud-based CRM with tools for sales, marketing, and customer support.', 'https://cdn-01.rapida.ai/partners/tools/zoho-icon.png',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/zoho", "connect_type": "oauth2"}'::jsonb),

(7646645603519189367, 'Salesforce', 'Leading CRM platform offering solutions for sales, service, marketing, and analytics.', 'https://cdn-01.rapida.ai/partners/tools/salesforce-icon.png',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/salesforce", "connect_type": "oauth2"}'::jsonb),

(7646645603519189368, 'Pipedrive', 'Sales CRM and pipeline management tool for small and medium-sized businesses.', 'https://cdn-01.rapida.ai/partners/tools/pipedrive-icon.png',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/pipedrive", "connect_type": "oauth2"}'::jsonb),

(7646645603519189369, 'LeadSquared', 'CRM for sales execution and marketing automation tailored for businesses of all sizes.', 'https://cdn-01.rapida.ai/partners/tools/leadsquared-icon.png',
 '["connection"]'::jsonb,
 '{"connect_url": "/connect-crm/leadsquared", "connect_type": "oauth2"}'::jsonb);