package kong

import (
	"encoding/json"
	"fmt"

	"github.com/kevholditch/gokong"
)

type Kong struct {
	KongHost        string
	UpStreamName    string
	TargetPath      string
	TargetPort      string
	TargetWeight    int
	ServiceName     string
	ServiceProtocol string
	ServicePort     int
	RouteProtocol   []string
	RouteHost       []string
	RoutePath       string
}

func InitKong(kongSetting Kong) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	config := gokong.Config{HostAddress: kongSetting.KongHost}
	client := gokong.NewClient(&config)
	// Upstream
	updateUpstreamRequest := &gokong.UpstreamRequest{}
	json.Unmarshal(upstreamJson, updateUpstreamRequest)
	updateUpstreamRequest.Name = kongSetting.UpStreamName
	updateUpstreamRequest.HealthChecks.Active.HttpPath = fmt.Sprintf("/%s/kong/healthchecks", kongSetting.RoutePath)
	var updatedUpstream *gokong.Upstream
	if upstream, _ := client.Upstreams().GetByName(kongSetting.UpStreamName); upstream != nil {
		updatedUpstream, _ = client.Upstreams().UpdateByName(kongSetting.UpStreamName, updateUpstreamRequest)
	} else {
		updatedUpstream, _ = client.Upstreams().Create(updateUpstreamRequest)
	}
	fmt.Printf("Upstream: %s，%+v", updatedUpstream.Id, updatedUpstream)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Target
	targetRequest := &gokong.TargetRequest{
		Target: kongSetting.TargetPath + ":" + kongSetting.TargetPort,
		Weight: kongSetting.TargetWeight,
	}
	createdTarget, err := client.Targets().CreateFromUpstreamId(updatedUpstream.Id, targetRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Target: %s，%+v", *createdTarget.Id, createdTarget)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Service
	serviceRequest := &gokong.ServiceRequest{
		Name:     &kongSetting.ServiceName,
		Protocol: &kongSetting.ServiceProtocol,
		Host:     &kongSetting.UpStreamName,
		Port:     &kongSetting.ServicePort,
	}

	var updatedService *gokong.Service
	if service, _ := client.Services().GetServiceByName(kongSetting.ServiceName); service != nil {
		updatedService, _ = client.Services().UpdateServiceByName(kongSetting.ServiceName, serviceRequest)
	} else {
		updatedService, _ = client.Services().Create(serviceRequest)
	}

	fmt.Printf("Service: %s，%+v", *updatedService.Id, updatedService)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Route
	routeRequest := &gokong.RouteRequest{
		Name:         gokong.String(*updatedService.Name + "-Route"),
		Protocols:    gokong.StringSlice(kongSetting.RouteProtocol),
		Methods:      gokong.StringSlice([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS", "HEAD", "TRACE", "CONNECT"}),
		Hosts:        gokong.StringSlice(kongSetting.RouteHost),
		Paths:        gokong.StringSlice([]string{fmt.Sprintf("/%s/(?i)", kongSetting.RoutePath)}),
		StripPath:    gokong.Bool(false),
		PreserveHost: gokong.Bool(false),
		Service:      gokong.ToId(*updatedService.Id),
	}

	var updatedRoute *gokong.Route
	if route, _ := client.Routes().GetByName(*routeRequest.Name); route != nil {
		updatedRoute, _ = client.Routes().UpdateByName(*routeRequest.Name, routeRequest)
	} else {
		updatedRoute, _ = client.Routes().Create(routeRequest)
	}

	fmt.Printf("Route: %s，%+v", *updatedRoute.Id, updatedRoute)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("------------------Kong注册成功------------------------")
	fmt.Println("-----------------------------------------------------")
}
