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
	Code     int
	Apps     []*Application `json:"apps,omitempty"`
	App      *Application   `json:"app,omitempty"`
	Versions []string       `json:",omitempty"`
	Tasks    []*Task        `json:"tasks,omitempty"`
}

// Application marathon application see :
// https://mesosphere.github.io/marathon/docs/rest-api.html#apps
type Application struct {
	ID            string            `json:"id"`
	Cmd           string            `json:"cmd,omitempty"`
	Constraints   [][]string        `json:"constraints,omitempty"`
	Container     *Container        `json:"container,omitempty"`
	CPUs          float32           `json:"cpus,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	Executor      string            `json:"executor,omitempty"`
	HealtChecks   []*HealthCheck    `json:"healtChecks,omitempty"`
	Instances     int               `json:"instances,omitemptys"`
	Mem           float32           `json:"mem,omitempty"`
	Tasks         []*Task           `json:"tasks,omitempty"`
	Ports         []int             `json:"ports,omitempty"`
	BackoffFactor float32           `json:"backoffFactor,omitempty"`
	TasksRunning  int               `json:"tasksRunning,omitempty"`
	TasksStaged   int               `json:"tasksStaged,omitempty"`
	Uris          []string          `json:"uris,omitempty"`
	Version       string            `json:"version,omitempty"`
}

// Container is docker parameters
// options are passed to container, if you want your options
// to be passed at the end of your docker run
// add // in front of the parameters you want to pass
// Example:
// docker run -ti -p 4343:4343 mysql --listen 0.0.0.0:4343
// options := [ "-p", "4343", "//", "--listen", "0.0.0.0:4343" ]
type Container struct {
	Image   string   `json:"image,omitempty"`
	Options []string `json:"options,omitempty"`
}

// HealthChecks are described here:
// https://mesosphere.github.io/marathon/docs/health-checks.html
type HealthCheck struct {
	Protocol           string `json:"protocol,omitempty"`
	Path               string `json:"path,omitempty"`
	GracePeriodSeconds int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds    int    `json:"intervalSeconds,omitempty"`
	PortIndex          int    `json:"portIndex,omitempty"`
	TimeoutSeconds     int    `json:"timeoutSeconds,omitempty"`
}

// Task is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#tasks
type Task struct {
	AppID     string `json:"appId"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Ports     []int  `json:"ports"`
	StagedAt  string `json:"stagedAt"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

// EventSubscription are described here :
// https://mesosphere.github.io/marathon/docs/rest-api.html#event-subscriptions
type EventSubscription struct {
	CallbackURL  string   `json:"CallbackUrl"`
	ClientIP     string   `json:"ClientIp"`
	EventType    string   `json:"eventType"`
	CallbackURLs []string `json:"CallbackUrls"`
}
