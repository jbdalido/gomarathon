package gomarathon

// RequestOptions passed for query api
type RequestOptions struct {
	Method string
	Path   string
	Datas  interface{}
	Params *Parameters
}

// Parameters to build url query
type Parameters struct {
	Cmd         string
	Host        string
	Scale       bool
	CallbackURL string
}

// Response representation of a full marathon response
type Response struct {
	Code         int
	Apps         []*Application `json:"apps,omitempty"`
	App          *Application   `json:"app,omitempty"`
	Versions     []string       `json:",omitempty"`
	Tasks        []*Task        `json:"tasks,omitempty"`
	DeploymentId string         `json:"deployment_id,omitempty"`
}

// Application marathon application see :
// https://github.com/mesosphere/marathon/blob/master/REST.md#apps
type Application struct {
	ID              string            `json:"id"`
	Cmd             string            `json:"cmd,omitempty"`
	Constraints     [][]string        `json:"constraints,omitempty"`
	Container       *Container        `json:"container,omitempty"`
	CPUs            float32           `json:"cpus,omitempty"`
	Deployments     []*Deployment     `json:"deployments,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
	Executor        string            `json:"executor,omitempty"`
	HealthChecks    []*HealthCheck    `json:"healthChecks,omitempty"`
	Instances       int               `json:"instances,omitemptys"`
	Mem             float32           `json:"mem,omitempty"`
	Tasks           []*Task           `json:"tasks,omitempty"`
	Ports           []int             `json:"ports,omitempty"`
	RequirePorts    bool              `json:"requirePorts,omitempty"`
	BackoffFactor   float32           `json:"backoffFactor,omitempty"`
	TasksRunning    int               `json:"tasksRunning,omitempty"`
	TasksStaged     int               `json:"tasksStaged,omitempty"`
	UpgradeStrategy *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Uris            []string          `json:"uris,omitempty"`
	Version         string            `json:"version,omitempty"`
        Labels          map[string]string `json:"labels,omitempty"`
}

// Container is docker parameters
type Container struct {
	Type    string    `json:"type,omitempty"`
	Docker  *Docker   `json:"docker,omitempty"`
	Volumes []*Volume `json:"volumes,omitempty"`
}

// Docker options
type Docker struct {
	Image        string         `json:"image,omitempty"`
	Network      string         `json:"network,omitempty"`
	PortMappings []*PortMapping `json:"portMappings,omitempty"`
}

// Volume is used for mounting a host directory as a container volume
type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

// Container PortMappings
type PortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// UpgradeStrategy has a minimumHealthCapacity which defines the minimum number of healty nodes
type UpgradeStrategy struct {
	MinimumHealthCapacity float32 `json:"minimumHealthCapacity,omitempty"`
}

// HealthCheck is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#healthchecks
type HealthCheck struct {
	Protocol               string `json:"protocol,omitempty"`
	Path                   string `json:"path,omitempty"`
	GracePeriodSeconds     int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int    `json:"intervalSeconds,omitempty"`
	PortIndex              int    `json:"portIndex,omitempty"`
	TimeoutSeconds         int    `json:"timeoutSeconds,omitempty"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures"`
}

type HealthCheckResult struct {
	Alive               bool   `json:"alive,omitempty"`
	ConsecutiveFailures int    `json:"consecutiveFailures,omitempty"`
	FirstSuccess        string `json:"firstSuccess,omitempty"`
	LastFailure         string `json:"lastFailure,omitempty"`
	LastSuccess         string `json:"lastSuccess,omitempty"`
	TaskID              string `json:"taskId,omitempty"`
}

// Task is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#tasks
type Task struct {
	AppID              string               `json:"appId"`
	Host               string               `json:"host"`
	ID                 string               `json:"id"`
	Ports              []int                `json:"ports"`
	StagedAt           string               `json:"stagedAt"`
	StartedAt          string               `json:"startedAt"`
	Version            string               `json:"version"`
	HealthCheckResults []*HealthCheckResult `json:"healthCheckResults"`
}

// Deployment is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#get-/v2/deployments
type Deployment struct {
	AffectedApps   []string          `json:"affectedApps"`
	ID             string            `json:"id"`
	Steps          []*DeploymentStep `json:"steps"`
	CurrentActions []*DeploymentStep `json:"currentActions"`
	CurrentStep    int               `json:"currentStep"`
	TotalSteps     int               `json:"totalSteps"`
	Version        string            `json:"version"`
}

// Deployment steps
type DeploymentStep struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// EventSubscription is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#event-subscriptions
type EventSubscription struct {
	CallbackURL  string   `json:"CallbackUrl"`
	ClientIP     string   `json:"ClientIp"`
	EventType    string   `json:"eventType"`
	CallbackURLs []string `json:"CallbackUrls"`
}
