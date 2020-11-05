package test_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	"k8s.io/klog"
	"os"
)

var _ = Describe("Test", func() {
	It("test", func() {
		fmt.Println("-----------")
		data := os.Getenv("VERSION")
		klog.Info(data)
		if data == "a" {
			fmt.Println("true")
		}

		fmt.Println("----------")
		fmt.Println(os.Getenv("KUBECONFIG"))
	})
})
