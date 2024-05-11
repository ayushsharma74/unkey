// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type V1ApisGetAPIRequest struct {
	APIID string `queryParam:"style=form,explode=true,name=apiId"`
}

func (o *V1ApisGetAPIRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

// V1ApisGetAPIResponseBody - The configuration for an api
type V1ApisGetAPIResponseBody struct {
	// The id of the key
	ID string `json:"id"`
	// The id of the workspace that owns the api
	WorkspaceID string `json:"workspaceId"`
	// The name of the api. This is internal and your users will not see this.
	Name *string `json:"name,omitempty"`
}

func (o *V1ApisGetAPIResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *V1ApisGetAPIResponseBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *V1ApisGetAPIResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
