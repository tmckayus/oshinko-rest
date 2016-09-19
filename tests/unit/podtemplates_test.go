package unittest

import (
	"gopkg.in/check.v1"
	kapi "k8s.io/kubernetes/pkg/api"

	"github.com/radanalyticsio/oshinko-rest/helpers/containers"
	"github.com/radanalyticsio/oshinko-rest/helpers/podtemplates"
)

// This function is named TestCreatePodTemplateSpec because there is another
// function named TestPodTemplateSpec.
func (s *OshinkoUnitTestSuite) TestCreatePodTemplateSpec(c *check.C) {
	newPodTemplateSpec := podtemplates.PodTemplateSpec()
	c.Assert(*newPodTemplateSpec, check.FitsTypeOf, podtemplates.OPodTemplateSpec{})
}

func (s *OshinkoUnitTestSuite) TestPodTemplateSetLabels(c *check.C) {
	expectedLabels := map[string]string{"test": "value"}
	newPodTemplateSpec := podtemplates.PodTemplateSpec()
	newPodTemplateSpec.SetLabels(expectedLabels)
	c.Assert(newPodTemplateSpec.PodTemplateSpec.GetLabels(), check.DeepEquals, expectedLabels)
}

func (s *OshinkoUnitTestSuite) TestPodTemplateLabel(c *check.C) {
	expectedLabels := map[string]string{"test": "value"}
	newPodTemplateSpec := podtemplates.PodTemplateSpec()
	newPodTemplateSpec.SetLabels(map[string]string{})
	newPodTemplateSpec.Label("test", "value")
	c.Assert(newPodTemplateSpec.PodTemplateSpec.GetLabels(), check.DeepEquals, expectedLabels)
}

func (s *OshinkoUnitTestSuite) TestContainers(c *check.C) {
	expectedContainers := []kapi.Container{
		kapi.Container{Name: "container1"},
		kapi.Container{Name: "container2"}}
	newPodTemplateSpec := podtemplates.PodTemplateSpec()
	newPodTemplateSpec.Containers(
		&containers.OContainer{Container: expectedContainers[0]},
		&containers.OContainer{Container: expectedContainers[1]})
	c.Assert(newPodTemplateSpec.Spec.Containers, check.DeepEquals, expectedContainers)
}
