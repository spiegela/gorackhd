package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/client/catalog"
	"github.com/emccode/gorackhd/client/catalogs"
	"github.com/emccode/gorackhd/client/config"
	"github.com/emccode/gorackhd/client/delete"
	"github.com/emccode/gorackhd/client/dhcp"
	"github.com/emccode/gorackhd/client/files"
	"github.com/emccode/gorackhd/client/get"
	"github.com/emccode/gorackhd/client/identify"
	"github.com/emccode/gorackhd/client/lookups"
	"github.com/emccode/gorackhd/client/nodes"
	"github.com/emccode/gorackhd/client/obm"
	"github.com/emccode/gorackhd/client/obms"
	"github.com/emccode/gorackhd/client/patch"
	"github.com/emccode/gorackhd/client/pollers"
	"github.com/emccode/gorackhd/client/post"
	"github.com/emccode/gorackhd/client/profiles"
	"github.com/emccode/gorackhd/client/put"
	"github.com/emccode/gorackhd/client/skus"
	"github.com/emccode/gorackhd/client/tags"
	"github.com/emccode/gorackhd/client/task"
	"github.com/emccode/gorackhd/client/templates"
	"github.com/emccode/gorackhd/client/versions"
	"github.com/emccode/gorackhd/client/whitelist"
	"github.com/emccode/gorackhd/client/workflow"
)

// Default monorail HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new monorail HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Monorail {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/api/1.1", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new monorail client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Monorail {
	cli := new(Monorail)
	cli.Transport = transport

	cli.Catalog = catalog.New(transport, formats)

	cli.Catalogs = catalogs.New(transport, formats)

	cli.Config = config.New(transport, formats)

	cli.Delete = delete.New(transport, formats)

	cli.Dhcp = dhcp.New(transport, formats)

	cli.Files = files.New(transport, formats)

	cli.Get = get.New(transport, formats)

	cli.Identify = identify.New(transport, formats)

	cli.Lookups = lookups.New(transport, formats)

	cli.Nodes = nodes.New(transport, formats)

	cli.Obm = obm.New(transport, formats)

	cli.Obms = obms.New(transport, formats)

	cli.Patch = patch.New(transport, formats)

	cli.Pollers = pollers.New(transport, formats)

	cli.Post = post.New(transport, formats)

	cli.Profiles = profiles.New(transport, formats)

	cli.Put = put.New(transport, formats)

	cli.Skus = skus.New(transport, formats)

	cli.Tags = tags.New(transport, formats)

	cli.Task = task.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

	cli.Versions = versions.New(transport, formats)

	cli.Whitelist = whitelist.New(transport, formats)

	cli.Workflow = workflow.New(transport, formats)

	return cli
}

// Monorail is a client for monorail
type Monorail struct {
	Catalog *catalog.Client

	Catalogs *catalogs.Client

	Config *config.Client

	Delete *delete.Client

	Dhcp *dhcp.Client

	Files *files.Client

	Get *get.Client

	Identify *identify.Client

	Lookups *lookups.Client

	Nodes *nodes.Client

	Obm *obm.Client

	Obms *obms.Client

	Patch *patch.Client

	Pollers *pollers.Client

	Post *post.Client

	Profiles *profiles.Client

	Put *put.Client

	Skus *skus.Client

	Tags *tags.Client

	Task *task.Client

	Templates *templates.Client

	Versions *versions.Client

	Whitelist *whitelist.Client

	Workflow *workflow.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Monorail) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Catalog.SetTransport(transport)

	c.Catalogs.SetTransport(transport)

	c.Config.SetTransport(transport)

	c.Delete.SetTransport(transport)

	c.Dhcp.SetTransport(transport)

	c.Files.SetTransport(transport)

	c.Get.SetTransport(transport)

	c.Identify.SetTransport(transport)

	c.Lookups.SetTransport(transport)

	c.Nodes.SetTransport(transport)

	c.Obm.SetTransport(transport)

	c.Obms.SetTransport(transport)

	c.Patch.SetTransport(transport)

	c.Pollers.SetTransport(transport)

	c.Post.SetTransport(transport)

	c.Profiles.SetTransport(transport)

	c.Put.SetTransport(transport)

	c.Skus.SetTransport(transport)

	c.Tags.SetTransport(transport)

	c.Task.SetTransport(transport)

	c.Templates.SetTransport(transport)

	c.Versions.SetTransport(transport)

	c.Whitelist.SetTransport(transport)

	c.Workflow.SetTransport(transport)

}
