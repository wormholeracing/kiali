package business

import (
	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/kubernetes/kubetest"
	osappsv1 "github.com/openshift/api/apps/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	batch_v1 "k8s.io/api/batch/v1"
	batch_v1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	"testing"
)

func setupAppService(k8s *kubetest.K8SClientMock) AppService {
	return AppService{k8s: k8s}
}

func TestGetAppListFromDeployments(t *testing.T) {
	assert := assert.New(t)
	conf := config.NewConfig()
	config.Set(conf)

	// Setup mocks
	k8s := new(kubetest.K8SClientMock)
	// Auxiliar fake* tests defined in workload_test.go
	k8s.On("IsOpenShift").Return(true)
	k8s.On("GetDeployments", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(FakeDeployments(), nil)
	k8s.On("GetDeploymentConfigs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]osappsv1.DeploymentConfig{}, nil)
	k8s.On("GetReplicaSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.ReplicaSet{}, nil)
	k8s.On("GetReplicationControllers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.ReplicationController{}, nil)
	k8s.On("GetStatefulSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.StatefulSet{}, nil)
	k8s.On("GetJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1.Job{}, nil)
	k8s.On("GetCronJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1beta1.CronJob{}, nil)
	k8s.On("GetPods", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.Pod{}, nil)
	svc := setupAppService(k8s)

	appList, _ := svc.GetAppList("Namespace")

	assert.Equal("Namespace", appList.Namespace.Name)

	assert.Equal(1, len(appList.Apps))
	assert.Equal("httpbin", appList.Apps[0].Name)
}

func TestGetAppFromDeployments(t *testing.T) {
	assert := assert.New(t)
	conf := config.NewConfig()
	config.Set(conf)

	// Setup mocks
	k8s := new(kubetest.K8SClientMock)
	k8s.On("IsOpenShift").Return(true)
	k8s.On("GetDeployments", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(FakeDeployments(), nil)
	k8s.On("GetDeploymentConfigs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]osappsv1.DeploymentConfig{}, nil)
	k8s.On("GetReplicaSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.ReplicaSet{}, nil)
	k8s.On("GetReplicationControllers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.ReplicationController{}, nil)
	k8s.On("GetStatefulSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.StatefulSet{}, nil)
	k8s.On("GetPods", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.Pod{}, nil)
	k8s.On("GetJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1.Job{}, nil)
	k8s.On("GetCronJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1beta1.CronJob{}, nil)
	k8s.On("GetServices", mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string")).Return(FakeServices(), nil)
	svc := setupAppService(k8s)

	appDetails, _ := svc.GetApp("Namespace", "httpbin")

	assert.Equal("Namespace", appDetails.Namespace.Name)
	assert.Equal("httpbin", appDetails.Name)

	assert.Equal(2, len(appDetails.Workloads))
	assert.Equal("httpbin-v1", appDetails.Workloads[0].WorkloadName)
	assert.Equal("httpbin-v2", appDetails.Workloads[1].WorkloadName)
}

func TestGetAppListFromReplicaSets(t *testing.T) {
	assert := assert.New(t)
	conf := config.NewConfig()
	config.Set(conf)

	// Setup mocks
	k8s := new(kubetest.K8SClientMock)
	// Auxiliar fake* tests defined in workload_test.go
	k8s.On("IsOpenShift").Return(true)
	k8s.On("GetDeployments", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta1.Deployment{}, nil)
	k8s.On("GetDeploymentConfigs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]osappsv1.DeploymentConfig{}, nil)
	k8s.On("GetReplicaSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(FakeReplicaSets(), nil)
	k8s.On("GetReplicationControllers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.ReplicationController{}, nil)
	k8s.On("GetStatefulSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.StatefulSet{}, nil)
	k8s.On("GetJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1.Job{}, nil)
	k8s.On("GetCronJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1beta1.CronJob{}, nil)
	k8s.On("GetPods", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.Pod{}, nil)
	svc := setupAppService(k8s)

	appList, _ := svc.GetAppList("Namespace")

	assert.Equal("Namespace", appList.Namespace.Name)

	assert.Equal(1, len(appList.Apps))
	assert.Equal("httpbin", appList.Apps[0].Name)
}

func TestGetAppFromReplicaSets(t *testing.T) {
	assert := assert.New(t)
	conf := config.NewConfig()
	config.Set(conf)

	// Setup mocks
	k8s := new(kubetest.K8SClientMock)
	k8s.On("IsOpenShift").Return(true)
	k8s.On("GetDeployments", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta1.Deployment{}, nil)
	k8s.On("GetDeploymentConfigs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]osappsv1.DeploymentConfig{}, nil)
	k8s.On("GetReplicaSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(FakeReplicaSets(), nil)
	k8s.On("GetReplicationControllers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.ReplicationController{}, nil)
	k8s.On("GetStatefulSets", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1beta2.StatefulSet{}, nil)
	k8s.On("GetPods", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]v1.Pod{}, nil)
	k8s.On("GetJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1.Job{}, nil)
	k8s.On("GetCronJobs", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]batch_v1beta1.CronJob{}, nil)
	k8s.On("GetServices", mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string")).Return(FakeServices(), nil)
	svc := setupAppService(k8s)

	appDetails, _ := svc.GetApp("Namespace", "httpbin")

	assert.Equal("Namespace", appDetails.Namespace.Name)
	assert.Equal("httpbin", appDetails.Name)

	assert.Equal(2, len(appDetails.Workloads))
	assert.Equal("httpbin-v1", appDetails.Workloads[0].WorkloadName)
	assert.Equal("httpbin-v2", appDetails.Workloads[1].WorkloadName)
}
