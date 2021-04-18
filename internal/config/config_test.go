package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/oryanmoshe/escargot/internal/config"
)

var _ = Describe("Config", func() {
	Describe("initialization using viper", func() {
		var (
			storeId     string
			storeConfig config.StoreConfig
		)

		BeforeEach(func() {
			storeId = "oryan-activities-store"
		})

		Context("by providing a storeId", func() {
			JustBeforeEach(func() {
				storeConfig = config.GetConfigByStoreID(storeId)
			})

			It("should return a valid StoreConfig object", func() {
				var expected config.StoreConfig = config.StoreConfig{}
				Expect(storeConfig).To(BeAssignableToTypeOf(expected))
			})

			It("should have the correct fields set in the configuration", func() {
				expected := config.StoreConfig{
					ID:   storeId,
					Name: "oryan-activities-redis-cluster-eu",
					Host: "redis-cluster-host.redis.com",
					Pass: "password1234",
				}

				Expect(storeConfig).To(Equal(expected))
			})
		})
	})
})
