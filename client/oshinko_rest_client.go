package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/radanalyticsio/oshinko-rest/client/clusters"
	"github.com/radanalyticsio/oshinko-rest/client/server"
)

// Default oshinko rest HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new oshinko rest HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OshinkoRest {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new oshinko rest client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OshinkoRest {
	cli := new(OshinkoRest)
	cli.Transport = transport

	cli.Clusters = clusters.New(transport, formats)

	cli.Server = server.New(transport, formats)

	return cli
}

// OshinkoRest is a client for oshinko rest
type OshinkoRest struct {
	Clusters *clusters.Client

	Server *server.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OshinkoRest) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Clusters.SetTransport(transport)

	c.Server.SetTransport(transport)

}
