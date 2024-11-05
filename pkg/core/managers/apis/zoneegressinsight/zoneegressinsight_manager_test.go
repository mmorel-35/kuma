package zoneegressinsight_test

import (
	"context"

	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	kuma_cp "github.com/kumahq/kuma/pkg/config/app/kuma-cp"
	"github.com/kumahq/kuma/pkg/core/managers/apis/zoneegressinsight"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/plugins/resources/memory"
)

var _ = Describe("ZoneEgressInsight Manager", func() {
	It("should limit the number of subscription", func() {
		// setup
		s := memory.NewStore()
		cfg := &kuma_cp.DataplaneMetrics{
			SubscriptionLimit: 3,
		}
		manager := zoneegressinsight.NewZoneEgressInsightManager(s, cfg)

		err := s.Create(context.Background(), mesh.NewZoneEgressResource(), store.CreateByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		input := mesh.NewZoneEgressInsightResource()
		for i := 0; i < 10; i++ {
			input.Spec.Subscriptions = append(input.Spec.Subscriptions, &mesh_proto.DiscoverySubscription{
				Id: strconv.Itoa(i),
			})
		}

		// when
		err = manager.Create(context.Background(), input, store.CreateByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		actual := mesh.NewZoneEgressInsightResource()
		err = s.Get(context.Background(), actual, store.GetByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		// then
		Expect(actual.Spec.Subscriptions).To(HaveLen(3))
		Expect(actual.Spec.Subscriptions[0].Id).To(Equal("7"))
		Expect(actual.Spec.Subscriptions[1].Id).To(Equal("8"))
		Expect(actual.Spec.Subscriptions[2].Id).To(Equal("9"))
	})

	It("should cleanup subscriptions if limit is 0", func() {
		// setup
		s := memory.NewStore()
		cfg := &kuma_cp.DataplaneMetrics{
			SubscriptionLimit: 0,
		}
		manager := zoneegressinsight.NewZoneEgressInsightManager(s, cfg)

		err := s.Create(context.Background(), mesh.NewZoneEgressResource(), store.CreateByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		input := mesh.NewZoneEgressInsightResource()
		for i := 0; i < 10; i++ {
			input.Spec.Subscriptions = append(input.Spec.Subscriptions, &mesh_proto.DiscoverySubscription{
				Id: strconv.Itoa(i),
			})
		}

		// when
		err = manager.Create(context.Background(), input, store.CreateByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		actual := mesh.NewZoneEgressInsightResource()
		err = s.Get(context.Background(), actual, store.GetByKey("di1", model.NoMesh))
		Expect(err).ToNot(HaveOccurred())

		// then
		Expect(actual.Spec.Subscriptions).To(BeEmpty())
		Expect(actual.Spec.Subscriptions).To(BeNil())
	})
})
