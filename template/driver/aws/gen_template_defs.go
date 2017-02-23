/* Copyright 2017 WALLIX

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

// DO NOT EDIT
// This file was automatically generated with go generate
package aws

import (
	"github.com/wallix/awless/template"
)

var AWSTemplatesDefinitions = map[string]template.TemplateDefinition{
	"createvpc": {
		Action:         "create",
		Entity:         "vpc",
		Api:            "ec2",
		RequiredParams: []string{"cidr"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletevpc": {
		Action:         "delete",
		Entity:         "vpc",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createsubnet": {
		Action:         "create",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "vpc"},
		ExtraParams:    []string{"zone"},
		TagsMapping:    []string{},
	},
	"updatesubnet": {
		Action:         "update",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{"public-vms"},
		TagsMapping:    []string{},
	},
	"deletesubnet": {
		Action:         "delete",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createinstance": {
		Action:         "create",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"image", "type", "count", "count", "subnet"},
		ExtraParams:    []string{"lock", "key", "ip", "group", "userdata"},
		TagsMapping:    []string{"name"},
	},
	"updateinstance": {
		Action:         "update",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{"lock", "group", "type"},
		TagsMapping:    []string{},
	},
	"deleteinstance": {
		Action:         "delete",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"startinstance": {
		Action:         "start",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"stopinstance": {
		Action:         "stop",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"checkinstance": {
		Action:         "check",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id", "state", "timeout"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createsecuritygroup": {
		Action:         "create",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"description", "name", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"updatesecuritygroup": {
		Action:         "update",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "id", "protocol"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletesecuritygroup": {
		Action:         "delete",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createvolume": {
		Action:         "create",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"zone", "size"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletevolume": {
		Action:         "delete",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachvolume": {
		Action:         "attach",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"device", "instance", "id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createinternetgateway": {
		Action:         "create",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteinternetgateway": {
		Action:         "delete",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachinternetgateway": {
		Action:         "attach",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachinternetgateway": {
		Action:         "detach",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createroutetable": {
		Action:         "create",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteroutetable": {
		Action:         "delete",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachroutetable": {
		Action:         "attach",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"id", "subnet"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachroutetable": {
		Action:         "detach",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"association"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createroute": {
		Action:         "create",
		Entity:         "route",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "gateway", "table"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteroute": {
		Action:         "delete",
		Entity:         "route",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "table"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createtags": {
		Action:         "create",
		Entity:         "tags",
		Api:            "ec2",
		RequiredParams: []string{"resource"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createkeypair": {
		Action:         "create",
		Entity:         "keypair",
		Api:            "ec2",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletekeypair": {
		Action:         "delete",
		Entity:         "keypair",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createuser": {
		Action:         "create",
		Entity:         "user",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteuser": {
		Action:         "delete",
		Entity:         "user",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"creategroup": {
		Action:         "create",
		Entity:         "group",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletegroup": {
		Action:         "delete",
		Entity:         "group",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachpolicy": {
		Action:         "attach",
		Entity:         "policy",
		Api:            "iam",
		RequiredParams: []string{"arn", "user"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachpolicy": {
		Action:         "detach",
		Entity:         "policy",
		Api:            "iam",
		RequiredParams: []string{"arn", "user"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createbucket": {
		Action:         "create",
		Entity:         "bucket",
		Api:            "s3",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletebucket": {
		Action:         "delete",
		Entity:         "bucket",
		Api:            "s3",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createstorageobject": {
		Action:         "create",
		Entity:         "storageobject",
		Api:            "s3",
		RequiredParams: []string{"file", "bucket"},
		ExtraParams:    []string{"name"},
		TagsMapping:    []string{},
	},
	"deletestorageobject": {
		Action:         "delete",
		Entity:         "storageobject",
		Api:            "s3",
		RequiredParams: []string{"bucket", "key"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createtopic": {
		Action:         "create",
		Entity:         "topic",
		Api:            "sns",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletetopic": {
		Action:         "delete",
		Entity:         "topic",
		Api:            "sns",
		RequiredParams: []string{"arn"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createsubscription": {
		Action:         "create",
		Entity:         "subscription",
		Api:            "sns",
		RequiredParams: []string{"endpoint", "protocol", "topic"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletesubscription": {
		Action:         "delete",
		Entity:         "subscription",
		Api:            "sns",
		RequiredParams: []string{"arn"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
}
