package httpclient

import (
    "fmt"
    "strings"
    "encoding/json"
    "testing"
)


type BasicDeployment struct {
    ApiVersion string `json:"apiVersion"`
    Kind string `json:"kind"`
    Metadata metadataDL `json:"metadata"`
    Spec specDL `json:"spec"`
}

type metadataDL struct {
    Name string `json:"name"`
    Namespace string `json:"namespace"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations,omitempty"`
}

type specDL struct {
    Replicas float64 `json:"replicas"`
    Template templateSpec `json:"template"`
}

type templateSpec struct {
    Metadata metadataTSDL `json:"metadata"`
    Spec specTSDL `json:"spec"`
}

type metadataTSDL struct {
    Name string `json:"name"`
    Labels map[string] string `json:"labels"`
}

type specTSDL struct {
    Volumes []volumesSTSDL `json:"volumes,omitempty"`
    Containers []ContainersSTSDL `json:"containers"` 
}

type volumesSTSDL struct {
    Name string `json:"name"`
    HostPath *hostPathVSTSDL `json:"hostPath,omitempty"`
    EmptyDir *emptyDirVSTSDL `json:"emptyDir,omitempty"`
    PersistentVolumeClaim *pvVolumeClaimVSTSDL `json:"persistentVolumeClaim,omitempty"`
    Rbd *rbdVSTSDL `json:"rbd,omitempty"`
    ConfigMap *configMapVSTSDL `json:"configMap,omitempty"`
}

type hostPathVSTSDL struct {
    Path string `json:"path"`
}

type emptyDirVSTSDL struct {
    Medium string `json:"medium"`
}

type pvVolumeClaimVSTSDL struct {
    ClaimName string `json:"claimName"`
    ReadOnly bool `json:"readOnly"`
}

type rbdVSTSDL struct {
    Monitors []string `json:"monitors"`
    Image string `json:"image"`
    FsType string `json:"fsType"`
    Pool string `json:"pool"`
    User string `json:"user"`
    Keyring string `json:"keyring"`
    SecretRef *secretRefRVSTSDL `json:"secretRef"`
    ReadOnly bool `json:"readOnly"` 
}

type secretRefRVSTSDL struct {
    Name string `json: name"`
}

type configMapVSTSDL struct {
    Name string `json:"name"` 
    Items []itemsConfigMap `json:"items"`
}

type itemsConfigMap struct {
    Key string `json:"key"`
    Path string `json:"Path"`
}

type ContainerType struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command,omitempty"`
    Args []string `json:"args,omitempty"`
    Ports []portContainer `json:"ports"`     
    Env []envContainer `json:"env,omitempty"`
    Resources *resourcesContainer `json:"resources,omitempty"`
    VolumeMounts []volumeMountsContainer `json:"volumeMounts,omitempty"` 
    LivenessProbe *livenessProbeContainer `json:"livenessProbe,omitempty"`
    ReadinessProbe *readinessProbeContainer `json:"readinessProbe,omitempty"`
    Lifecycle *lifecycleContainer `json:"lifecycle,omitempty"`
 
}


type ContainersSTSDL struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command,omitempty"`
    Args []string `json:"args,omitempty"`
    Ports []portContainer `json:"ports"`     
    Env []envContainer `json:"env,omitempty"`
    Resources *resourcesContainer `json:"resources,omitempty"`
    VolumeMounts []volumeMountsContainer `json:"volumeMounts,omitempty"` 
    LivenessProbe *livenessProbeContainer `json:"livenessProbe,omitempty"`
    ReadinessProbe *readinessProbeContainer `json:"readinessProbe,omitempty"`
    Lifecycle *lifecycleContainer `json:"lifecycle,omitempty"`
 
}

type portContainer struct {
    Name string `json:"name"`
    HostPort float64 `json:"hostPort"`
    ContainerPort float64 `json:"containerPort"`
    Protocol string `json:"protocol"`
    HostIP string `json:"hostIP"` 
}

type envContainer struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type resourcesContainer struct {
    Limits map[string] string `json:"limits"`
    Requests map[string] string `json:"requests"`
}
    
type volumeMountsContainer struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"` 
}

type livenessProbeContainer struct {
    Exec *execLiveProbeType `json:"exec,omitempty"`
    HttpGet *httpGetLiveProbeType `json:"httpGet,omitempty"`
    TcpSocket *tcpLiveProbeType `json:"tcpSocket,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}

type execLiveProbeType struct {
    Command []string `json:"command"` 
}

type httpGetLiveProbeType struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []headersGLPType `json:"httpHeaders"`    
}

type headersGLPType struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type tcpLiveProbeType struct {
    Port float64 `json:"port"`
}

type readinessProbeContainer struct {
    HttpGet *httpGetReadProbeType `json:"httpGet,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}


type httpGetReadProbeType struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []headersGRPType `json:"httpHeaders"`    
}

type headersGRPType struct {
    Name string `json:"name"`
    Value string `json:"value"`
}


type lifecycleContainer struct {
    PostStart *postStartLCType `json:"postStart,omitempty"`
    PreStop *preStopLCType `json:"preStop,omitempty"`    
}

type postStartLCType struct {
    Exec *execPSLCType `json:"exec,omitempty"`
}

type execPSLCType struct {
    Command []string `json:"command"`
}

type preStopLCType struct {
    Exec *execPrSLCType `json:"exec,omitempty"`
}

type execPrSLCType struct {
    Command []string `json:"command"`
}





func Test_HttpClient_New(t *testing.T) {
    client := NewHttpClient("http://master", "8080")
    if client == nil {
        t.Error("HttpClient_New didn't pass")

    } else {
        fmt.Printf("%v\n", client)
        t.Log("test 1 passed")
    }
}

func Test_HttpClient_Get(t *testing.T) {
    client := NewHttpClient("http://master", "8080")
    resp, err := client.Get("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments")
    if err != nil {
        t.Error("Get pods error")
    } else {
        fmt.Println(string(resp))
    }
}


func Test_HttpClient_Post(t *testing.T) {

    tempDeploy := new(BasicDeployment)      

    defaultlabels := make(map[string]string, 1)
    defaultlabels["appname"] = "nginx-test" 

    tempDeploy.ApiVersion = "extensions/v1beta1"
    tempDeploy.Kind = "Deployment"

    tempDeploy.Metadata.Name = "nginx-test" 
    tempDeploy.Spec.Replicas = 1 
    tempDeploy.Spec.Template.Metadata.Labels = defaultlabels
    tempDeploy.Spec.Template.Spec.Containers = make([]ContainersSTSDL, 1)
    tempDeploy.Spec.Template.Spec.Containers[0].Name = "nginx-test" 
    tempDeploy.Spec.Template.Spec.Containers[0].Image = "nginx:1.7.9" 

    var result []byte
        result, _ = json.MarshalIndent(tempDeploy, "", "    ")

     
    client := NewHttpClient("http://master", "8080")
	resp, err := client.Post("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments", strings.NewReader(string(result)))

    if err != nil {
        t.Error("Post pods error")
    } else {
        fmt.Println(string(resp))
    }

}

func Test_HttpClient_Delete(t *testing.T) {
    client := NewHttpClient("http://master", "8080")
    resp, err := client.Delete("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments/nginx-test")
    if err != nil {
        t.Error("Delete pods error")
    } else {
        fmt.Println(string(resp))
    }   
}


