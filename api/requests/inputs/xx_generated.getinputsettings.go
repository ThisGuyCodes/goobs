// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputSettings request.
type GetInputSettingsParams struct {
	// Name of the input to get the settings of
	InputName *string `json:"inputName,omitempty"`
}

func NewGetInputSettingsParams() *GetInputSettingsParams {
	return &GetInputSettingsParams{}
}
func (o *GetInputSettingsParams) WithInputName(x string) *GetInputSettingsParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *GetInputSettingsParams) GetRequestName() string {
	return "GetInputSettings"
}

// Represents the response body for the GetInputSettings request.
type GetInputSettingsResponse struct {
	_response

	// The kind of the input
	InputKind string `json:"inputKind,omitempty"`

	// Object of settings for the input
	InputSettings map[string]any `json:"inputSettings,omitempty"`
}

/*
Gets the settings of an input.

Note: Does not include defaults. To create the entire settings object, overlay `inputSettings` over the `defaultInputSettings` provided by `GetInputDefaultSettings`.
*/
func (c *Client) GetInputSettings(params *GetInputSettingsParams) (*GetInputSettingsResponse, error) {
	data := &GetInputSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
