package lats_test

import (
	"lats/helpers"
	"plumbing/conversion"
	v2 "plumbing/v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/sonde-go/events"
)

var _ = Describe("Container Metrics Endpoint", func() {
	It("can receive container metrics", func() {
		envelope := createContainerMetric("test-id")
		helpers.EmitToMetronV1(envelope)

		f := func() []*events.ContainerMetric {
			return helpers.RequestContainerMetrics("test-id")
		}
		Eventually(f).Should(ContainElement(envelope.ContainerMetric))
	})

	Describe("emit v2 and consume via reverse log proxy", func() {
		It("can receive container metrics", func() {
			envelope := createContainerMetric("test-id")
			v2Env := conversion.ToV2(envelope)
			helpers.EmitToMetronV2(v2Env)

			f := func() []*v2.Envelope {
				return helpers.ReadContainerFromRLP("test-id")
			}
			Eventually(f).Should(ContainElement(v2Env))
		})
	})
})
