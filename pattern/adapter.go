package main

import "fmt"

type Client interface {
	List(group string)
	Get(name string, group string)
	Delete(name string, group string)
	DeleteTag(name string, group string, tag string)
}

type RegistryClient interface {
	list(group string)
	get(name string, group string)
	delete(name string, group string)
	deleteTag(name string, group string, tag string)
}

type SWRRegistryClient struct{}
type HarborRegistryClient struct{}

type Adapter struct {
	RegistryClient
}

func (src SWRRegistryClient) list(group string) {
	fmt.Println("It is SWR List")
}

func (sr SWRRegistryClient) get(name string, group string) {
	fmt.Println("It is SWR Get")
}

func (sr SWRRegistryClient) delete(name string, group string) {
	fmt.Println("It is SWR Delete")
}

func (sr SWRRegistryClient) deleteTag(name string, group string, tag string) {
	fmt.Println("It is SWR DeleteTag")
}

func (hrc HarborRegistryClient) list(group string) {
	fmt.Println("It is Harbor List")
}

func (hrc HarborRegistryClient) get(name string, group string) {
	fmt.Println("It is Harbor Get")
}

func (hrc HarborRegistryClient) delete(name string, group string) {
	fmt.Println("It is Harbor Delete")
}

func (hrc HarborRegistryClient) deleteTag(name string, group string, tag string) {
	fmt.Println("It is Harbor DeleteTag")
}

func (adapter Adapter) List(group string) {
	adapter.list(group)
}

func NewAdapter(rc RegistryClient) Adapter {
	return Adapter{
		rc,
	}
}

func NewRegistryClient(registry string) RegistryClient {
	switch registry {
	case "swr":
		return SWRRegistryClient{}
	case "harbor":
		return HarborRegistryClient{}
	default:
		return nil
	}
}

func main() {
	swrRegistry := NewRegistryClient("swr")
	registryClient := NewAdapter(swrRegistry)
	registryClient.List("SGT")

	fmt.Println("=======================")
	harborRegistry := NewRegistryClient("harbor")
	registryClient = NewAdapter(harborRegistry)
	registryClient.List("SBS")
}
