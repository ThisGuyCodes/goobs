// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputAudioSyncOffset request.
type GetInputAudioSyncOffsetParams struct {
	// Name of the input to get the audio sync offset of
	InputName *string `json:"inputName,omitempty"`
}

func NewGetInputAudioSyncOffsetParams() *GetInputAudioSyncOffsetParams {
	return &GetInputAudioSyncOffsetParams{}
}
func (o *GetInputAudioSyncOffsetParams) WithInputName(x string) *GetInputAudioSyncOffsetParams {
	o.InputName = &x
	return o
}

// Returns the associated request.
func (o *GetInputAudioSyncOffsetParams) GetRequestName() string {
	return "GetInputAudioSyncOffset"
}

// Represents the response body for the GetInputAudioSyncOffset request.
type GetInputAudioSyncOffsetResponse struct {
	_response

	// Audio sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`
}

/*
Gets the audio sync offset of an input.

Note: The audio sync offset can be negative too!
*/
func (c *Client) GetInputAudioSyncOffset(
	params *GetInputAudioSyncOffsetParams,
) (*GetInputAudioSyncOffsetResponse, error) {
	data := &GetInputAudioSyncOffsetResponse{}
	return data, c.client.SendRequest(params, data)
}
