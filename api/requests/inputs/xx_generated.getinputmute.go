// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputMute request.
type GetInputMuteParams struct {
	// Name of input to get the mute state of
	InputName *string `json:"inputName,omitempty"`
}

func NewGetInputMuteParams() *GetInputMuteParams {
	return &GetInputMuteParams{}
}
func (o *GetInputMuteParams) WithInputName(x string) *GetInputMuteParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *GetInputMuteParams) GetRequestName() string {
	return "GetInputMute"
}

// Represents the response body for the GetInputMute request.
type GetInputMuteResponse struct {
	_response

	// Whether the input is muted
	InputMuted bool `json:"inputMuted,omitempty"`
}

// Gets the audio mute state of an input.
func (c *Client) GetInputMute(params *GetInputMuteParams) (*GetInputMuteResponse, error) {
	data := &GetInputMuteResponse{}
	return data, c.client.SendRequest(params, data)
}
