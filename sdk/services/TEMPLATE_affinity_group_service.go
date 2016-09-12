package services

import (
	ovirt_http "github.com/CpuID/ovirt-engine-sdk-go/sdk/http"
	"github.com/CpuID/ovirt-engine-sdk-go/sdk/types"
)

// This file was written by hand, to use as a reference example of what the golang output should look like.
// I then went ahead and started adjusting the Java codebase to generate this output at a minimum.
// Once this has been achieved, this file can be removed.

type AffinityGroupService struct {
}

// Creates a new implementation of the service.
func NewAffinityGroupService(c *ovirt_http.Connection, p string) *AffinityGroupService {

}

// Returns the representation of the object managed by this service.
func (a *AffinityGroupService) Get(opts map[string]string) *types.AffinityGroup {

}

func (a *AffinityGroupService) Remove(opts map[string]string) bool {

}

//
// TODO: what is the return type here? can't identify it easily from other SDK implementations? bool?
func (a *AffinityGroupService) Update(group *types.AffinityGroup) string {

}

// Locates the `vms` service.
func (a *AffinityGroupService) VmsService() *AffinityGroupVmsService {

}

// TODO: conflicting function name and type...? make them unique?
func (a *AffinityGroupService) Service(path string) *Service {

}

func (a *AffinityGroupService) ToString() string {

}
