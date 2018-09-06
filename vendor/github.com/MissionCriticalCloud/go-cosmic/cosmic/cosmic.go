//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cosmic

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

// UnlimitedResourceID is a special ID to define an unlimited resource
const UnlimitedResourceID = "-1"

var idRegex = regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|-1)$`)

// IsID return true if the passed ID is either a UUID or a UnlimitedResourceID
func IsID(id string) bool {
	return idRegex.MatchString(id)
}

// OptionFunc can be passed to the courtesy helper functions to set additional parameters
type OptionFunc func(*CosmicClient, interface{}) error

type CSError struct {
	ErrorCode   int    `json:"errorcode"`
	CSErrorCode int    `json:"cserrorcode"`
	ErrorText   string `json:"errortext"`
}

func (e *CSError) Error() error {
	return fmt.Errorf("Cosmic API error %d (CSExceptionErrorCode: %d): %s", e.ErrorCode, e.CSErrorCode, e.ErrorText)
}

type CosmicClient struct {
	HTTPGETOnly bool // If `true` only use HTTP GET calls

	client  *http.Client // The http client for communicating
	baseURL string       // The base URL of the API
	apiKey  string       // Api key
	secret  string       // Secret key
	async   bool         // Wait for async calls to finish
	timeout int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds

	Account          *AccountService
	AffinityGroup    *AffinityGroupService
	Alert            *AlertService
	Asyncjob         *AsyncjobService
	Authentication   *AuthenticationService
	Certificate      *CertificateService
	CloudOps         *CloudOpsService
	Cluster          *ClusterService
	Configuration    *ConfigurationService
	DiskOffering     *DiskOfferingService
	Domain           *DomainService
	Event            *EventService
	Firewall         *FirewallService
	GuestOS          *GuestOSService
	Host             *HostService
	Hypervisor       *HypervisorService
	ISO              *ISOService
	ImageStore       *ImageStoreService
	Limit            *LimitService
	LoadBalancer     *LoadBalancerService
	NAT              *NATService
	NetworkACL       *NetworkACLService
	NetworkDevice    *NetworkDeviceService
	NetworkOffering  *NetworkOfferingService
	Network          *NetworkService
	Nic              *NicService
	NiciraNVP        *NiciraNVPService
	Pod              *PodService
	Project          *ProjectService
	PublicIPAddress  *PublicIPAddressService
	Region           *RegionService
	Resourcemetadata *ResourcemetadataService
	Resourcetags     *ResourcetagsService
	Router           *RouterService
	SSH              *SSHService
	SecurityGroup    *SecurityGroupService
	ServiceOffering  *ServiceOfferingService
	Snapshot         *SnapshotService
	StoragePool      *StoragePoolService
	System           *SystemService
	SystemVM         *SystemVMService
	Template         *TemplateService
	Usage            *UsageService
	User             *UserService
	VLAN             *VLANService
	VMGroup          *VMGroupService
	VPC              *VPCService
	VPN              *VPNService
	VirtualMachine   *VirtualMachineService
	Volume           *VolumeService
	Zone             *ZoneService
}

// Creates a new client for communicating with Cosmic
func newClient(apiurl string, apikey string, secret string, async bool, tlsConfig *tls.Config, timeout int64) *CosmicClient {
	cs := &CosmicClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: tlsConfig,
			},
			Timeout: time.Duration(time.Duration(timeout) * time.Second),
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		async:   async,
		timeout: 300,
	}
	cs.Account = NewAccountService(cs)
	cs.AffinityGroup = NewAffinityGroupService(cs)
	cs.Alert = NewAlertService(cs)
	cs.Asyncjob = NewAsyncjobService(cs)
	cs.Authentication = NewAuthenticationService(cs)
	cs.Certificate = NewCertificateService(cs)
	cs.CloudOps = NewCloudOpsService(cs)
	cs.Cluster = NewClusterService(cs)
	cs.Configuration = NewConfigurationService(cs)
	cs.DiskOffering = NewDiskOfferingService(cs)
	cs.Domain = NewDomainService(cs)
	cs.Event = NewEventService(cs)
	cs.Firewall = NewFirewallService(cs)
	cs.GuestOS = NewGuestOSService(cs)
	cs.Host = NewHostService(cs)
	cs.Hypervisor = NewHypervisorService(cs)
	cs.ISO = NewISOService(cs)
	cs.ImageStore = NewImageStoreService(cs)
	cs.Limit = NewLimitService(cs)
	cs.LoadBalancer = NewLoadBalancerService(cs)
	cs.NAT = NewNATService(cs)
	cs.NetworkACL = NewNetworkACLService(cs)
	cs.NetworkDevice = NewNetworkDeviceService(cs)
	cs.NetworkOffering = NewNetworkOfferingService(cs)
	cs.Network = NewNetworkService(cs)
	cs.Nic = NewNicService(cs)
	cs.NiciraNVP = NewNiciraNVPService(cs)
	cs.Pod = NewPodService(cs)
	cs.Project = NewProjectService(cs)
	cs.PublicIPAddress = NewPublicIPAddressService(cs)
	cs.Region = NewRegionService(cs)
	cs.Resourcemetadata = NewResourcemetadataService(cs)
	cs.Resourcetags = NewResourcetagsService(cs)
	cs.Router = NewRouterService(cs)
	cs.SSH = NewSSHService(cs)
	cs.SecurityGroup = NewSecurityGroupService(cs)
	cs.ServiceOffering = NewServiceOfferingService(cs)
	cs.Snapshot = NewSnapshotService(cs)
	cs.StoragePool = NewStoragePoolService(cs)
	cs.System = NewSystemService(cs)
	cs.SystemVM = NewSystemVMService(cs)
	cs.Template = NewTemplateService(cs)
	cs.Usage = NewUsageService(cs)
	cs.User = NewUserService(cs)
	cs.VLAN = NewVLANService(cs)
	cs.VMGroup = NewVMGroupService(cs)
	cs.VPC = NewVPCService(cs)
	cs.VPN = NewVPNService(cs)
	cs.VirtualMachine = NewVirtualMachineService(cs)
	cs.Volume = NewVolumeService(cs)
	cs.Zone = NewZoneService(cs)
	return cs
}

// Default non-async client. So for async calls you need to implement and check the async job result yourself. When using
// HTTPS with a self-signed certificate to connect to your Cosmic API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings. Timeout for the http request is in seconds.
func NewClient(apiurl string, apikey string, secret string, tlsConfig *tls.Config, timeout int64) *CosmicClient {
	cs := newClient(apiurl, apikey, secret, false, tlsConfig, timeout)
	return cs
}

// For sync API calls this client behaves exactly the same as a standard client call, but for async API calls
// this client will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return actual object received from the API and nil, but when the timout is
// reached it will return the initial object containing the async job ID for the running job and a warning. Timeout for the http request is in seconds.
func NewAsyncClient(apiurl string, apikey string, secret string, tlsConfig *tls.Config, timeout int64) *CosmicClient {
	cs := newClient(apiurl, apikey, secret, true, tlsConfig, timeout)
	return cs
}

// When using the async client an api call will wait for the async call to finish before returning. The default is to poll for 300 seconds
// seconds, to check if the async job is finished.
func (cs *CosmicClient) AsyncTimeout(timeoutInSeconds int64) {
	cs.timeout = timeoutInSeconds
}

var AsyncTimeoutErr = errors.New("Timeout while waiting for async job to finish")

// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured
// timeout, the async job returns a AsyncTimeoutErr.
func (cs *CosmicClient) GetAsyncJobResult(jobid string, timeout int64) (json.RawMessage, error) {
	var timer time.Duration
	currentTime := time.Now().Unix()

	for {
		p := cs.Asyncjob.NewQueryAsyncJobResultParams(jobid)
		r, err := cs.Asyncjob.QueryAsyncJobResult(p)
		if err != nil {
			return nil, err
		}

		// Status 1 means the job is finished successfully
		if r.Jobstatus == 1 {
			return r.Jobresult, nil
		}

		// When the status is 2, the job has failed
		if r.Jobstatus == 2 {
			if r.Jobresulttype == "text" {
				return nil, fmt.Errorf(string(r.Jobresult))
			} else {
				return nil, fmt.Errorf("Undefined error: %s", string(r.Jobresult))
			}
		}

		if time.Now().Unix()-currentTime > timeout {
			return nil, AsyncTimeoutErr
		}

		// Add an (extremely simple) exponential backoff like feature to prevent
		// flooding the Cosmic API
		if timer < 15 {
			timer++
		}

		time.Sleep(timer * time.Second)
	}
}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occured. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *CosmicClient) newRequest(api string, params url.Values) (json.RawMessage, error) {
	params.Set("apiKey", cs.apiKey)
	params.Set("command", api)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters, URL encoding only values and sort them by key, done by encodeValues
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with Cosmic secret
	// * URL encode the string and convert to base64
	s := encodeValues(params)
	s2 := strings.ToLower(s)
	s3 := strings.Replace(s2, "+", "%20", -1)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	var err error
	var resp *http.Response
	if !cs.HTTPGETOnly && (api == "deployVirtualMachine" || api == "updateVirtualMachine") {
		// The deployVirtualMachine API should be called using a POST call
		// so we don't have to worry about the userdata size

		// Add the unescaped signature to the POST params
		params.Set("signature", signature)

		// Make a POST call
		resp, err = cs.client.PostForm(cs.baseURL, params)
	} else {
		// Create the final URL before we issue the request
		url := cs.baseURL + "?" + s + "&signature=" + url.QueryEscape(signature)

		// Make a GET call
		resp, err = cs.client.Get(url)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Need to get the raw value to make the result play nice
	b, err = getRawValue(b)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e CSError
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}
		return nil, e.Error()
	}
	return b, nil
}

// Custom version of net/url Encode that only URL escapes values
// Unmodified portions here remain under BSD license of The Go Authors: https://go.googlesource.com/go/+/master/LICENSE
func encodeValues(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// Generic function to get the first raw value from a response as json.RawMessage
func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for _, v := range m {
		return v, nil
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

// ProjectIDSetter is an interface that every type that can set a project ID must implement
type ProjectIDSetter interface {
	SetProjectid(string)
}

// WithProject takes either a project name or ID and sets the `projectid` parameter
func WithProject(project string) OptionFunc {
	return func(cs *CosmicClient, p interface{}) error {
		ps, ok := p.(ProjectIDSetter)

		if !ok || project == "" {
			return nil
		}

		if !IsID(project) {
			id, _, err := cs.Project.GetProjectID(project)
			if err != nil {
				return err
			}
			project = id
		}

		ps.SetProjectid(project)

		return nil
	}
}

// VPCIDSetter is an interface that every type that can set a vpc ID must implement
type VPCIDSetter interface {
	SetVpcid(string)
}

// WithVPCID takes a vpc ID and sets the `vpcid` parameter
func WithVPCID(id string) OptionFunc {
	return func(cs *CosmicClient, p interface{}) error {
		vs, ok := p.(VPCIDSetter)

		if !ok || id == "" {
			return nil
		}

		vs.SetVpcid(id)

		return nil
	}
}

type AccountService struct {
	cs *CosmicClient
}

func NewAccountService(cs *CosmicClient) *AccountService {
	return &AccountService{cs: cs}
}

type AffinityGroupService struct {
	cs *CosmicClient
}

func NewAffinityGroupService(cs *CosmicClient) *AffinityGroupService {
	return &AffinityGroupService{cs: cs}
}

type AlertService struct {
	cs *CosmicClient
}

func NewAlertService(cs *CosmicClient) *AlertService {
	return &AlertService{cs: cs}
}

type AsyncjobService struct {
	cs *CosmicClient
}

func NewAsyncjobService(cs *CosmicClient) *AsyncjobService {
	return &AsyncjobService{cs: cs}
}

type AuthenticationService struct {
	cs *CosmicClient
}

func NewAuthenticationService(cs *CosmicClient) *AuthenticationService {
	return &AuthenticationService{cs: cs}
}

type CertificateService struct {
	cs *CosmicClient
}

func NewCertificateService(cs *CosmicClient) *CertificateService {
	return &CertificateService{cs: cs}
}

type CloudOpsService struct {
	cs *CosmicClient
}

func NewCloudOpsService(cs *CosmicClient) *CloudOpsService {
	return &CloudOpsService{cs: cs}
}

type ClusterService struct {
	cs *CosmicClient
}

func NewClusterService(cs *CosmicClient) *ClusterService {
	return &ClusterService{cs: cs}
}

type ConfigurationService struct {
	cs *CosmicClient
}

func NewConfigurationService(cs *CosmicClient) *ConfigurationService {
	return &ConfigurationService{cs: cs}
}

type DiskOfferingService struct {
	cs *CosmicClient
}

func NewDiskOfferingService(cs *CosmicClient) *DiskOfferingService {
	return &DiskOfferingService{cs: cs}
}

type DomainService struct {
	cs *CosmicClient
}

func NewDomainService(cs *CosmicClient) *DomainService {
	return &DomainService{cs: cs}
}

type EventService struct {
	cs *CosmicClient
}

func NewEventService(cs *CosmicClient) *EventService {
	return &EventService{cs: cs}
}

type FirewallService struct {
	cs *CosmicClient
}

func NewFirewallService(cs *CosmicClient) *FirewallService {
	return &FirewallService{cs: cs}
}

type GuestOSService struct {
	cs *CosmicClient
}

func NewGuestOSService(cs *CosmicClient) *GuestOSService {
	return &GuestOSService{cs: cs}
}

type HostService struct {
	cs *CosmicClient
}

func NewHostService(cs *CosmicClient) *HostService {
	return &HostService{cs: cs}
}

type HypervisorService struct {
	cs *CosmicClient
}

func NewHypervisorService(cs *CosmicClient) *HypervisorService {
	return &HypervisorService{cs: cs}
}

type ISOService struct {
	cs *CosmicClient
}

func NewISOService(cs *CosmicClient) *ISOService {
	return &ISOService{cs: cs}
}

type ImageStoreService struct {
	cs *CosmicClient
}

func NewImageStoreService(cs *CosmicClient) *ImageStoreService {
	return &ImageStoreService{cs: cs}
}

type LimitService struct {
	cs *CosmicClient
}

func NewLimitService(cs *CosmicClient) *LimitService {
	return &LimitService{cs: cs}
}

type LoadBalancerService struct {
	cs *CosmicClient
}

func NewLoadBalancerService(cs *CosmicClient) *LoadBalancerService {
	return &LoadBalancerService{cs: cs}
}

type NATService struct {
	cs *CosmicClient
}

func NewNATService(cs *CosmicClient) *NATService {
	return &NATService{cs: cs}
}

type NetworkACLService struct {
	cs *CosmicClient
}

func NewNetworkACLService(cs *CosmicClient) *NetworkACLService {
	return &NetworkACLService{cs: cs}
}

type NetworkDeviceService struct {
	cs *CosmicClient
}

func NewNetworkDeviceService(cs *CosmicClient) *NetworkDeviceService {
	return &NetworkDeviceService{cs: cs}
}

type NetworkOfferingService struct {
	cs *CosmicClient
}

func NewNetworkOfferingService(cs *CosmicClient) *NetworkOfferingService {
	return &NetworkOfferingService{cs: cs}
}

type NetworkService struct {
	cs *CosmicClient
}

func NewNetworkService(cs *CosmicClient) *NetworkService {
	return &NetworkService{cs: cs}
}

type NicService struct {
	cs *CosmicClient
}

func NewNicService(cs *CosmicClient) *NicService {
	return &NicService{cs: cs}
}

type NiciraNVPService struct {
	cs *CosmicClient
}

func NewNiciraNVPService(cs *CosmicClient) *NiciraNVPService {
	return &NiciraNVPService{cs: cs}
}

type PodService struct {
	cs *CosmicClient
}

func NewPodService(cs *CosmicClient) *PodService {
	return &PodService{cs: cs}
}

type ProjectService struct {
	cs *CosmicClient
}

func NewProjectService(cs *CosmicClient) *ProjectService {
	return &ProjectService{cs: cs}
}

type PublicIPAddressService struct {
	cs *CosmicClient
}

func NewPublicIPAddressService(cs *CosmicClient) *PublicIPAddressService {
	return &PublicIPAddressService{cs: cs}
}

type RegionService struct {
	cs *CosmicClient
}

func NewRegionService(cs *CosmicClient) *RegionService {
	return &RegionService{cs: cs}
}

type ResourcemetadataService struct {
	cs *CosmicClient
}

func NewResourcemetadataService(cs *CosmicClient) *ResourcemetadataService {
	return &ResourcemetadataService{cs: cs}
}

type ResourcetagsService struct {
	cs *CosmicClient
}

func NewResourcetagsService(cs *CosmicClient) *ResourcetagsService {
	return &ResourcetagsService{cs: cs}
}

type RouterService struct {
	cs *CosmicClient
}

func NewRouterService(cs *CosmicClient) *RouterService {
	return &RouterService{cs: cs}
}

type SSHService struct {
	cs *CosmicClient
}

func NewSSHService(cs *CosmicClient) *SSHService {
	return &SSHService{cs: cs}
}

type SecurityGroupService struct {
	cs *CosmicClient
}

func NewSecurityGroupService(cs *CosmicClient) *SecurityGroupService {
	return &SecurityGroupService{cs: cs}
}

type ServiceOfferingService struct {
	cs *CosmicClient
}

func NewServiceOfferingService(cs *CosmicClient) *ServiceOfferingService {
	return &ServiceOfferingService{cs: cs}
}

type SnapshotService struct {
	cs *CosmicClient
}

func NewSnapshotService(cs *CosmicClient) *SnapshotService {
	return &SnapshotService{cs: cs}
}

type StoragePoolService struct {
	cs *CosmicClient
}

func NewStoragePoolService(cs *CosmicClient) *StoragePoolService {
	return &StoragePoolService{cs: cs}
}

type SystemService struct {
	cs *CosmicClient
}

func NewSystemService(cs *CosmicClient) *SystemService {
	return &SystemService{cs: cs}
}

type SystemVMService struct {
	cs *CosmicClient
}

func NewSystemVMService(cs *CosmicClient) *SystemVMService {
	return &SystemVMService{cs: cs}
}

type TemplateService struct {
	cs *CosmicClient
}

func NewTemplateService(cs *CosmicClient) *TemplateService {
	return &TemplateService{cs: cs}
}

type UsageService struct {
	cs *CosmicClient
}

func NewUsageService(cs *CosmicClient) *UsageService {
	return &UsageService{cs: cs}
}

type UserService struct {
	cs *CosmicClient
}

func NewUserService(cs *CosmicClient) *UserService {
	return &UserService{cs: cs}
}

type VLANService struct {
	cs *CosmicClient
}

func NewVLANService(cs *CosmicClient) *VLANService {
	return &VLANService{cs: cs}
}

type VMGroupService struct {
	cs *CosmicClient
}

func NewVMGroupService(cs *CosmicClient) *VMGroupService {
	return &VMGroupService{cs: cs}
}

type VPCService struct {
	cs *CosmicClient
}

func NewVPCService(cs *CosmicClient) *VPCService {
	return &VPCService{cs: cs}
}

type VPNService struct {
	cs *CosmicClient
}

func NewVPNService(cs *CosmicClient) *VPNService {
	return &VPNService{cs: cs}
}

type VirtualMachineService struct {
	cs *CosmicClient
}

func NewVirtualMachineService(cs *CosmicClient) *VirtualMachineService {
	return &VirtualMachineService{cs: cs}
}

type VolumeService struct {
	cs *CosmicClient
}

func NewVolumeService(cs *CosmicClient) *VolumeService {
	return &VolumeService{cs: cs}
}

type ZoneService struct {
	cs *CosmicClient
}

func NewZoneService(cs *CosmicClient) *ZoneService {
	return &ZoneService{cs: cs}
}
