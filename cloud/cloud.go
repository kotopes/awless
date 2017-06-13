/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloud

import (
	"errors"
	"fmt"
	"strings"

	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/template/driver"
)

var ErrFetchAccessDenied = errors.New("access denied to cloud resource")

// Resources
const (
	Region string = "region"
	//infra
	Vpc              string = "vpc"
	Subnet           string = "subnet"
	Image            string = "image"
	ImportImageTask  string = "importimagetask"
	SecurityGroup    string = "securitygroup"
	AvailabilityZone string = "availabilityzone"
	Keypair          string = "keypair"
	Volume           string = "volume"
	Instance         string = "instance"
	InstanceProfile  string = "instanceprofile"
	InternetGateway  string = "internetgateway"
	RouteTable       string = "routetable"
	ElasticIP        string = "elasticip"
	Snapshot         string = "snapshot"
	//loadbalancer
	LoadBalancer string = "loadbalancer"
	TargetGroup  string = "targetgroup"
	Listener     string = "listener"
	//database
	Database      string = "database"
	DbSubnetGroup string = "dbsubnetgroup"
	//access
	User         string = "user"
	Role         string = "role"
	Group        string = "group"
	Policy       string = "policy"
	AccessKey    string = "accesskey"
	LoginProfile string = "loginprofile"
	//storage
	Bucket   string = "bucket"
	S3Object string = "s3object"
	Acl      string = "storageacl"
	//notification
	Subscription string = "subscription"
	Topic        string = "topic"
	//queue
	Queue string = "queue"
	//dns
	Zone   string = "zone"
	Record string = "record"
	//lambda
	Function string = "function"
	//autoscaling
	LaunchConfiguration string = "launchconfiguration"
	ScalingGroup        string = "scalinggroup"
	ScalingPolicy       string = "scalingpolicy"
	//monitoring
	Metric string = "metric"
	Alarm  string = "alarm"
	//cdn
	Distribution string = "distribution"
	//cloudformation
	Stack string = "stack"
	//container
	Repository              string = "repository"
	Registry                string = "registry"
	ContainerCluster        string = "containercluster"
	ContainerService        string = "containerservice"
	ContainerTaskDefinition string = "containertaskdefinition"
	Container               string = "container"
)

type Service interface {
	Name() string
	Drivers() []driver.Driver
	ResourceTypes() []string
	FetchResources() (*graph.Graph, error)
	IsSyncDisabled() bool
	FetchByType(t string) (*graph.Graph, error)
}

type Services []Service

func (srvs Services) Names() (names []string) {
	for _, srv := range srvs {
		names = append(names, srv.Name())
	}
	return
}

var ServiceRegistry = make(map[string]Service)

func GetServiceForType(t string) (Service, error) {
	for _, srv := range ServiceRegistry {
		for _, typ := range srv.ResourceTypes() {
			if typ == t {
				return srv, nil
			}
		}
	}
	return nil, fmt.Errorf("cannot find cloud service for resource type %s", t)
}

func PluralizeResource(singular string) string {
	if strings.HasSuffix(singular, "cy") || strings.HasSuffix(singular, "ry") {
		return strings.TrimSuffix(singular, "y") + "ies"
	}
	return singular + "s"
}
