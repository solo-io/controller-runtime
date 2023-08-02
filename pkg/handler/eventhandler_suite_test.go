/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/controller-runtime/pkg/envtest"
	logf "github.com/solo-io/controller-runtime/pkg/log"
	"github.com/solo-io/controller-runtime/pkg/log/zap"
	"k8s.io/client-go/rest"
)

func TestEventhandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eventhandler Suite")
}

var testenv *envtest.Environment
var cfg *rest.Config

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	testenv = &envtest.Environment{}
	var err error
	cfg, err = testenv.Start()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(testenv.Stop()).To(Succeed())
})
