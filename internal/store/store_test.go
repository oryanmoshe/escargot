package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oryanmoshe/escargot/internal/config"
	"github.com/oryanmoshe/escargot/internal/store"
)

var _ = Describe("Store", func() {
	Describe("initialization", func() {
		var (
			storeConfig config.StoreConfig
			storeObject store.Store
		)

		BeforeEach(func() {
			storeConfig = config.StoreConfig{
				ID:   "oryan-activities-store",
				Name: "oryan-activities-redis-cluster-eu",
				Host: "redis-cluster-host.redis.com",
				Pass: "password1234",
			}
		})

		Context("by providing StoreConfig", func() {
			JustBeforeEach(func() {
				storeObject = store.New(storeConfig)
			})

			It("should return a struct implementing the Store interface", func() {
				var expected store.Store = store.New(config.StoreConfig{})
				Expect(storeObject).To(BeAssignableToTypeOf(expected))
			})

			It("should have the correct name returned by GetName()", func() {
				Expect(storeObject.GetName()).To(Equal(storeConfig.Name))
			})

			It("should not fail when calling Whoop()", func() {
				storeObject.Whoop("test")
			})
		})

		Context("by providing a store ID", func() {
			JustBeforeEach(func() {
				storeObject = store.FromID(storeConfig.ID)
			})

			It("should return a struct implementing the Store interface", func() {
				var expected store.Store = store.New(config.StoreConfig{})
				Expect(storeObject).To(BeAssignableToTypeOf(expected))
			})

			It("should have the correct name returned by GetName()", func() {
				Expect(storeObject.GetName()).To(Equal(storeConfig.Name))
			})

			It("should not fail when calling Whoop()", func() {
				storeObject.Whoop("test")
			})
		})
	})
})
