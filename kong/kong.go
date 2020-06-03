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
	updatedUpstream, _ := client.Upstreams().UpdateByName(kongSetting.UpStreamName, updateUpstreamRequest)
	fmt.Printf("Upstream: %+v", updatedUpstream)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Target
	targetRequest := &gokong.TargetRequest{
		Target: kongSetting.TargetPath + ":" + kongSetting.TargetPort,
		Weight: kongSetting.TargetWeight,
	}
	createdTarget, _ := client.Targets().CreateFromUpstreamId(updatedUpstream.Id, targetRequest)
	fmt.Printf("Target: %+v", createdTarget)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Service
	serviceRequest := &gokong.ServiceRequest{
		Name:     &kongSetting.ServiceName,
		Protocol: &kongSetting.ServiceProtocol,
		Host:     &kongSetting.UpStreamName,
		Port:     &kongSetting.ServicePort,
	}
	updatedService, _ := client.Services().UpdateServiceByName(*serviceRequest.Name, serviceRequest)
	fmt.Printf("Service: %+v", updatedService)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")

	// Route
	routeRequest := &gokong.RouteRequest{
		Name:         gokong.String(*updatedService.Name + "-Route"),
		Protocols:    gokong.StringSlice(kongSetting.RouteProtocol),
		Methods:      gokong.StringSlice([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS", "HEAD", "TRACE", "CONNECT"}),
		Hosts:        gokong.StringSlice(kongSetting.RouteHost),
		Paths:        gokong.StringSlice([]string{"/(?i)"}),
		StripPath:    gokong.Bool(false),
		PreserveHost: gokong.Bool(false),
		Service:      gokong.ToId(*updatedService.Id),
	}
	updatedRoute, _ := client.Routes().UpdateByName(*routeRequest.Name, routeRequest)
	fmt.Printf("Route: %+v", updatedRoute)
	fmt.Println("")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("------------------Kong注册成功------------------------")
	fmt.Println("-----------------------------------------------------")
}
