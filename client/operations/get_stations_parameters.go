// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetStationsParams creates a new GetStationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStationsParams() *GetStationsParams {
	return &GetStationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStationsParamsWithTimeout creates a new GetStationsParams object
// with the ability to set a timeout on a request.
func NewGetStationsParamsWithTimeout(timeout time.Duration) *GetStationsParams {
	return &GetStationsParams{
		timeout: timeout,
	}
}

// NewGetStationsParamsWithContext creates a new GetStationsParams object
// with the ability to set a context for a request.
func NewGetStationsParamsWithContext(ctx context.Context) *GetStationsParams {
	return &GetStationsParams{
		Context: ctx,
	}
}

// NewGetStationsParamsWithHTTPClient creates a new GetStationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStationsParamsWithHTTPClient(client *http.Client) *GetStationsParams {
	return &GetStationsParams{
		HTTPClient: client,
	}
}

/* GetStationsParams contains all the parameters to send to the API endpoint
   for the get stations operation.

   Typically these are written to a http.Request.
*/
type GetStationsParams struct {

	/* Allegiance.

	   Name of the allegiance.
	*/
	Allegiance *string

	/* BeginsWith.

	   Starting characters of the station.
	*/
	BeginsWith *string

	/* Count.

	   Number of history records. Disables timeMin and timeMax
	*/
	Count *string

	/* Economy.

	   Station economy.
	*/
	Economy *string

	/* EddbID.

	   EDDB ID of the station.
	*/
	EddbID *string

	/* Government.

	   Name of the government type.
	*/
	Government *string

	/* ID.

	   ID of the document.
	*/
	ID *string

	/* Name.

	   Station name.
	*/
	Name *string

	/* Page.

	   Page no of response.
	*/
	Page *int64

	/* State.

	   State the station is in.
	*/
	State *string

	/* System.

	   System name the station is in.
	*/
	System *string

	/* SystemID.

	   Filter by system id.
	*/
	SystemID *string

	/* TimeMax.

	   Maximum time for the station history in milliseconds.
	*/
	TimeMax *string

	/* TimeMin.

	   Minimum time for the station history in milliseconds.
	*/
	TimeMin *string

	/* Type.

	   Station type.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStationsParams) WithDefaults() *GetStationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stations params
func (o *GetStationsParams) WithTimeout(timeout time.Duration) *GetStationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stations params
func (o *GetStationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stations params
func (o *GetStationsParams) WithContext(ctx context.Context) *GetStationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stations params
func (o *GetStationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stations params
func (o *GetStationsParams) WithHTTPClient(client *http.Client) *GetStationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stations params
func (o *GetStationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllegiance adds the allegiance to the get stations params
func (o *GetStationsParams) WithAllegiance(allegiance *string) *GetStationsParams {
	o.SetAllegiance(allegiance)
	return o
}

// SetAllegiance adds the allegiance to the get stations params
func (o *GetStationsParams) SetAllegiance(allegiance *string) {
	o.Allegiance = allegiance
}

// WithBeginsWith adds the beginsWith to the get stations params
func (o *GetStationsParams) WithBeginsWith(beginsWith *string) *GetStationsParams {
	o.SetBeginsWith(beginsWith)
	return o
}

// SetBeginsWith adds the beginsWith to the get stations params
func (o *GetStationsParams) SetBeginsWith(beginsWith *string) {
	o.BeginsWith = beginsWith
}

// WithCount adds the count to the get stations params
func (o *GetStationsParams) WithCount(count *string) *GetStationsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get stations params
func (o *GetStationsParams) SetCount(count *string) {
	o.Count = count
}

// WithEconomy adds the economy to the get stations params
func (o *GetStationsParams) WithEconomy(economy *string) *GetStationsParams {
	o.SetEconomy(economy)
	return o
}

// SetEconomy adds the economy to the get stations params
func (o *GetStationsParams) SetEconomy(economy *string) {
	o.Economy = economy
}

// WithEddbID adds the eddbID to the get stations params
func (o *GetStationsParams) WithEddbID(eddbID *string) *GetStationsParams {
	o.SetEddbID(eddbID)
	return o
}

// SetEddbID adds the eddbId to the get stations params
func (o *GetStationsParams) SetEddbID(eddbID *string) {
	o.EddbID = eddbID
}

// WithGovernment adds the government to the get stations params
func (o *GetStationsParams) WithGovernment(government *string) *GetStationsParams {
	o.SetGovernment(government)
	return o
}

// SetGovernment adds the government to the get stations params
func (o *GetStationsParams) SetGovernment(government *string) {
	o.Government = government
}

// WithID adds the id to the get stations params
func (o *GetStationsParams) WithID(id *string) *GetStationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stations params
func (o *GetStationsParams) SetID(id *string) {
	o.ID = id
}

// WithName adds the name to the get stations params
func (o *GetStationsParams) WithName(name *string) *GetStationsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get stations params
func (o *GetStationsParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the get stations params
func (o *GetStationsParams) WithPage(page *int64) *GetStationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get stations params
func (o *GetStationsParams) SetPage(page *int64) {
	o.Page = page
}

// WithState adds the state to the get stations params
func (o *GetStationsParams) WithState(state *string) *GetStationsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get stations params
func (o *GetStationsParams) SetState(state *string) {
	o.State = state
}

// WithSystem adds the system to the get stations params
func (o *GetStationsParams) WithSystem(system *string) *GetStationsParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get stations params
func (o *GetStationsParams) SetSystem(system *string) {
	o.System = system
}

// WithSystemID adds the systemID to the get stations params
func (o *GetStationsParams) WithSystemID(systemID *string) *GetStationsParams {
	o.SetSystemID(systemID)
	return o
}

// SetSystemID adds the systemId to the get stations params
func (o *GetStationsParams) SetSystemID(systemID *string) {
	o.SystemID = systemID
}

// WithTimeMax adds the timeMax to the get stations params
func (o *GetStationsParams) WithTimeMax(timeMax *string) *GetStationsParams {
	o.SetTimeMax(timeMax)
	return o
}

// SetTimeMax adds the timeMax to the get stations params
func (o *GetStationsParams) SetTimeMax(timeMax *string) {
	o.TimeMax = timeMax
}

// WithTimeMin adds the timeMin to the get stations params
func (o *GetStationsParams) WithTimeMin(timeMin *string) *GetStationsParams {
	o.SetTimeMin(timeMin)
	return o
}

// SetTimeMin adds the timeMin to the get stations params
func (o *GetStationsParams) SetTimeMin(timeMin *string) {
	o.TimeMin = timeMin
}

// WithType adds the typeVar to the get stations params
func (o *GetStationsParams) WithType(typeVar *string) *GetStationsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get stations params
func (o *GetStationsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetStationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Allegiance != nil {

		// query param allegiance
		var qrAllegiance string

		if o.Allegiance != nil {
			qrAllegiance = *o.Allegiance
		}
		qAllegiance := qrAllegiance
		if qAllegiance != "" {

			if err := r.SetQueryParam("allegiance", qAllegiance); err != nil {
				return err
			}
		}
	}

	if o.BeginsWith != nil {

		// query param beginsWith
		var qrBeginsWith string

		if o.BeginsWith != nil {
			qrBeginsWith = *o.BeginsWith
		}
		qBeginsWith := qrBeginsWith
		if qBeginsWith != "" {

			if err := r.SetQueryParam("beginsWith", qBeginsWith); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount string

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := qrCount
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Economy != nil {

		// query param economy
		var qrEconomy string

		if o.Economy != nil {
			qrEconomy = *o.Economy
		}
		qEconomy := qrEconomy
		if qEconomy != "" {

			if err := r.SetQueryParam("economy", qEconomy); err != nil {
				return err
			}
		}
	}

	if o.EddbID != nil {

		// query param eddbId
		var qrEddbID string

		if o.EddbID != nil {
			qrEddbID = *o.EddbID
		}
		qEddbID := qrEddbID
		if qEddbID != "" {

			if err := r.SetQueryParam("eddbId", qEddbID); err != nil {
				return err
			}
		}
	}

	if o.Government != nil {

		// query param government
		var qrGovernment string

		if o.Government != nil {
			qrGovernment = *o.Government
		}
		qGovernment := qrGovernment
		if qGovernment != "" {

			if err := r.SetQueryParam("government", qGovernment); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.System != nil {

		// query param system
		var qrSystem string

		if o.System != nil {
			qrSystem = *o.System
		}
		qSystem := qrSystem
		if qSystem != "" {

			if err := r.SetQueryParam("system", qSystem); err != nil {
				return err
			}
		}
	}

	if o.SystemID != nil {

		// query param systemId
		var qrSystemID string

		if o.SystemID != nil {
			qrSystemID = *o.SystemID
		}
		qSystemID := qrSystemID
		if qSystemID != "" {

			if err := r.SetQueryParam("systemId", qSystemID); err != nil {
				return err
			}
		}
	}

	if o.TimeMax != nil {

		// query param timeMax
		var qrTimeMax string

		if o.TimeMax != nil {
			qrTimeMax = *o.TimeMax
		}
		qTimeMax := qrTimeMax
		if qTimeMax != "" {

			if err := r.SetQueryParam("timeMax", qTimeMax); err != nil {
				return err
			}
		}
	}

	if o.TimeMin != nil {

		// query param timeMin
		var qrTimeMin string

		if o.TimeMin != nil {
			qrTimeMin = *o.TimeMin
		}
		qTimeMin := qrTimeMin
		if qTimeMin != "" {

			if err := r.SetQueryParam("timeMin", qTimeMin); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
