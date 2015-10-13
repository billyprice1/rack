// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kms

import (
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/signer/v4"
)

// AWS Key Management Service (KMS) is an encryption and key management web
// service. This guide describes the KMS actions that you can call programmatically.
// For general information about KMS, see the  AWS Key Management Service Developer
// Guide  (http://docs.aws.amazon.com/kms/latest/developerguide/overview.html)
//
//  AWS provides SDKs that consist of libraries and sample code for various
// programming languages and platforms (Java, Ruby, .Net, iOS, Android, etc.).
// The SDKs provide a convenient way to create programmatic access to KMS and
// AWS. For example, the SDKs take care of tasks such as signing requests (see
// below), managing errors, and retrying requests automatically. For more information
// about the AWS SDKs, including how to download and install them, see Tools
// for Amazon Web Services (http://aws.amazon.com/tools/).   We recommend that
// you use the AWS SDKs to make programmatic API calls to KMS.
//
// Clients must support TLS (Transport Layer Security) 1.0. We recommend TLS
// 1.2. Clients must also support cipher suites with Perfect Forward Secrecy
// (PFS) such as Ephemeral Diffie-Hellman (DHE) or Elliptic Curve Ephemeral
// Diffie-Hellman (ECDHE). Most modern systems such as Java 7 and later support
// these modes.
//
// Signing Requests
//
//  Requests must be signed by using an access key ID and a secret access key.
// We strongly recommend that you do not use your AWS account access key ID
// and secret key for everyday work with KMS. Instead, use the access key ID
// and secret access key for an IAM user, or you can use the AWS Security Token
// Service to generate temporary security credentials that you can use to sign
// requests.
//
//  All KMS operations require Signature Version 4 (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// Recording API Requests
//
//  KMS supports AWS CloudTrail, a service that records AWS API calls and related
// events for your AWS account and delivers them to an Amazon S3 bucket that
// you specify. By using the information collected by CloudTrail, you can determine
// what requests were made to KMS, who made the request, when it was made, and
// so on. To learn more about CloudTrail, including how to turn it on and find
// your log files, see the AWS CloudTrail User Guide (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/whatiscloudtrail.html)
//
// Additional Resources
//
// For more information about credentials and request signing, see the following:
//
//   AWS Security Credentials (http://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html).
// This topic provides general information about the types of credentials used
// for accessing AWS.   AWS Security Token Service (http://docs.aws.amazon.com/STS/latest/UsingSTS/).
// This guide describes how to create and use temporary security credentials.
//   Signing AWS API Requests (http://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html).
// This set of topics walks you through the process of signing a request using
// an access key ID and a secret access key.   Commonly Used APIs
//
//  Of the APIs discussed in this guide, the following will prove the most
// useful for most applications. You will likely perform actions other than
// these, such as creating keys and assigning policies, by using the console.
//  Encrypt Decrypt GenerateDataKey GenerateDataKeyWithoutPlaintext
type KMS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new KMS client.
func New(config *aws.Config) *KMS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "kms",
			APIVersion:   "2014-11-01",
			JSONVersion:  "1.1",
			TargetPrefix: "TrentService",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &KMS{service}
}

// newRequest creates a new request for a KMS operation and runs any
// custom request initialization.
func (c *KMS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
