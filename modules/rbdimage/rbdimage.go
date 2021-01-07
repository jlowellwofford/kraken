/* cpuburn.go: this module induces artificial load on the system.  It can throttle itself based on thermal condition.
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

//go:generate protoc -I ../../core/proto/include -I proto --go_out=plugins=grpc:proto proto/rbdimage.proto

package cpuburn

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/hpc/kraken/core"
	"github.com/hpc/kraken/lib"

	cpb "github.com/hpc/kraken/core/proto"
	imgpb "github.com/hpc/kraken/extensions/ImageState/proto"
	pb "github.com/hpc/kraken/modules/rbdimage/proto"
)

// SrvStateURL is the URL For this service instance
const (
	SrvStateURL   = "/Services/rbdimage/State"
	imageStateURL = "type.googleapis.com/proto.ImageState/State"
	busyStateURL  = "type.googleapis.com/proto.ImageState/Busy"
)

var _ lib.Module = (*RBDImage)(nil)
var _ lib.ModuleSelfService = (*RBDImage)(nil)
var _ lib.ModuleWithConfig = (*RBDImage)(nil)
var _ lib.ModuleWithDiscovery = (*RBDImage)(nil)

// A RBDImage manages RBD images
type RBDImage struct {
	api   lib.APIClient
	cfg   *pb.RBDImageConfig
	dchan chan<- lib.Event
}

// Name returns a unique URL name for the module
// Module
func (*RBDImage) Name() string { return "github.com/hpc/kraken/modules/cpuburn" }

// Entry provides the module entry point
// ModuleSelfService
func (c *RBDImage) Entry() {
	url := lib.NodeURLJoin(c.api.Self().String(), SrvStateURL)
	ev := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			URL:     url,
			ValueID: "RUN",
		},
	)
	c.dchan <- ev
}

// Init is run before Entry, provides an APIClient
// ModuleSelfService
func (c *RBDImage) Init(api lib.APIClient) {
	c.api = api
	c.cfg = c.NewConfig().(*pb.RBDImageConfig)
}

// Stop should gracefully exit a running service instance
// ModuleSelfService
func (*RBDImage) Stop() {
	os.Exit(0)
}

// NewConfig should return a guaranteed sane (default) config
// ModuleWithConfig
func (*RBDImage) NewConfig() (r proto.Message) {
	r = &pb.RBDImageConfig{
		Namespace: "",
		Id:        "admin",
		Secret:    "",
		Monitors: [][]byte{
			net.IPv4(127, 0, 0, 1),
		},
	}
	return
}

// UpdateConfig should update the running module config
// ModuleWithConfig
func (c *RBDImage) UpdateConfig(cfg proto.Message) (e error) {
	if ccfg, ok := cfg.(*pb.RBDImageConfig); ok {
		c.cfg = ccfg
		return
	}
	// no active changes; current images use the old, new will use the new
	return fmt.Errorf("invalid config type")
}

// ConfigURL should return the unique url for the Config
// ModuleWithConfig
func (*RBDImage) ConfigURL() string {
	cfg := &pb.RBDImageConfig{}
	any, _ := ptypes.MarshalAny(cfg)
	return any.GetTypeUrl()
}

// SetDiscoveryChan takes a discovery channel (and presumably stores it)
// ModuleWithDiscovery
func (c *RBDImage) SetDiscoveryChan(dc chan<- lib.Event) { c.dchan = dc }

/*
 * Module specific functions
 */

// TBD

type ismut struct {
	f       imgpb.ImageState_State
	t       imgpb.ImageState_State
	timeout string
}

var ismuts = map[string]ismut{
	"NONEtoINACTIVE": {
		imgpb.ImageState_NONE,
		imgpb.ImageState_INACTIVE,
		"5s",
	},
	"INACTIVEtoLOADING": {
		imgpb.ImageState_INACTIVE,
		imgpb.ImageState_LOADING,
		"5s",
	},
	"LOADINGtoACTIVE": {
		imgpb.ImageState_LOADING,
		imgpb.ImageState_ACTIVE,
		"30s",
	},
}

var isnobusy = map[string]ismut{
	"ACTIVEtoINACTIVE": {
		imgpb.ImageState_ACTIVE,
		imgpb.ImageState_INACTIVE,
		"10s",
	},
	"UPDATEtoLOADING": {
		imgpb.ImageState_UPDATE,
		imgpb.ImageState_LOADING,
		"10s",
	},
}

func init() {
	module := &RBDImage{}
	si := core.NewServiceInstance("rbdimage", module.Name(), module.Entry)
	mutations := map[string]lib.StateMutation{}
	for k, v := range ismuts {
		dur, _ := time.ParseDuration(v.timeout)
		mutations[k] = core.NewStateMutation(
			map[string][2]reflect.Value{
				imageStateURL: {
					reflect.ValueOf(v.f),
					reflect.ValueOf(v.t),
				},
			},
			map[string]reflect.Value{"/RunState": reflect.ValueOf(cpb.Node_SYNC)},
			map[string]reflect.Value{},
			lib.StateMutationContext_SELF,
			dur,
			[3]string{si.ID(), imageStateURL, imgpb.ImageState_ERROR.Enum().String()},
		)
	}
	for k, v := range isnobusy {
		dur, _ := time.ParseDuration(v.timeout)
		mutations[k] = core.NewStateMutation(
			map[string][2]reflect.Value{
				imageStateURL: {
					reflect.ValueOf(v.f),
					reflect.ValueOf(v.t),
				},
			},
			map[string]reflect.Value{
				"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
				busyStateURL: reflect.ValueOf(imgpb.ImageState_IDLE),
			},
			map[string]reflect.Value{},
			lib.StateMutationContext_SELF,
			dur,
			[3]string{si.ID(), imageStateURL, imgpb.ImageState_ERROR.String()},
		)
	}
	mutations["UNKNOWNtoIDLE"] = core.NewStateMutation(
		map[string][2]reflect.Value{
			busyStateURL: {
				reflect.ValueOf(imgpb.ImageState_UNKNOWN),
				reflect.ValueOf(imgpb.ImageState_IDLE),
			},
		},
		map[string]reflect.Value{
			"/RunState": reflect.ValueOf(cpb.Node_SYNC),
		},
		map[string]reflect.Value{
			imageStateURL: reflect.ValueOf(imgpb.ImageState_NONE),
		},
		lib.StateMutationContext_SELF,
		time.Second*5,
		[3]string{si.ID(), imageStateURL, imgpb.ImageState_ERROR.String()},
	)

	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterMutations(si, mutations)
	core.Registry.RegisterDiscoverable(si, map[string]map[string]reflect.Value{
		imageStateURL: {
			imgpb.ImageState_NONE.String():     reflect.ValueOf(imgpb.ImageState_NONE),
			imgpb.ImageState_INACTIVE.String(): reflect.ValueOf(imgpb.ImageState_INACTIVE),
			imgpb.ImageState_LOADING.String():  reflect.ValueOf(imgpb.ImageState_LOADING),
			imgpb.ImageState_ACTIVE.String():   reflect.ValueOf(imgpb.ImageState_ACTIVE),
			imgpb.ImageState_UPDATE.String():   reflect.ValueOf(imgpb.ImageState_UPDATE),
			imgpb.ImageState_ERROR.String():    reflect.ValueOf(imgpb.ImageState_ERROR),
		},
		busyStateURL: {
			imgpb.ImageState_UNKNOWN.String(): reflect.ValueOf(imgpb.ImageState_UNKNOWN),
			imgpb.ImageState_IDLE.String():    reflect.ValueOf(imgpb.ImageState_IDLE),
			imgpb.ImageState_BUSY.String():    reflect.ValueOf(imgpb.ImageState_BUSY),
		},
		SrvStateURL: {
			"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN),
		},
	})
}
