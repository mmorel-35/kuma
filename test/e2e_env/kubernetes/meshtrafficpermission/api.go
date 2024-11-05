package meshtrafficpermission

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/pkg/plugins/policies/meshtrafficpermission/api/v1alpha1"
	. "github.com/kumahq/kuma/test/framework"
	"github.com/kumahq/kuma/test/framework/envs/kubernetes"
)

func API() {
	meshName := "meshtrafficpermission-api"

	BeforeAll(func() {
		err := NewClusterSetup().
			Install(MeshKubernetes(meshName)).
			Setup(kubernetes.Cluster)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEachFailure(func() {
		DebugKube(kubernetes.Cluster, meshName, Config.KumaNamespace)
	})

	E2EAfterEach(func() {
		Expect(DeleteMeshResources(kubernetes.Cluster, meshName, v1alpha1.MeshTrafficPermissionResourceTypeDescriptor)).To(Succeed())
	})

	E2EAfterAll(func() {
		Expect(kubernetes.Cluster.DeleteMesh(meshName)).To(Succeed())
	})

	It("should create MeshTrafficPermission policy", func() {
		// given no MeshTrafficPermissions
		mtps, err := kubernetes.Cluster.GetKumactlOptions().KumactlList("meshtrafficpermissions", meshName)
		Expect(err).ToNot(HaveOccurred())
		Expect(mtps).To(BeEmpty())

		// when
		Expect(YamlK8s(fmt.Sprintf(`
apiVersion: kuma.io/v1alpha1
kind: MeshTrafficPermission
metadata:
  name: mtp1
  namespace: %s
  labels:
    kuma.io/mesh: meshtrafficpermission-api
spec:
  targetRef:
    kind: MeshService
    name: backend
  from:
    - targetRef:
        kind: Mesh
      default:
        action: Allow
    - targetRef:
        kind: MeshService
        name: backend
      default:
        action: AllowWithShadowDeny
    - targetRef:
        kind: MeshServiceSubset
        name: backend
        tags:
          version: v1
      default:
        action: Deny
`, Config.KumaNamespace))(kubernetes.Cluster)).To(Succeed())

		// then
		mtps, err = kubernetes.Cluster.GetKumactlOptions().KumactlList("meshtrafficpermissions", meshName)
		Expect(err).ToNot(HaveOccurred())
		Expect(mtps).To(HaveLen(1))
		Expect(mtps[0]).To(Equal("mtp1." + Config.KumaNamespace))
	})
}
