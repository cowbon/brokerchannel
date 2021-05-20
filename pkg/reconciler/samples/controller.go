/*
Copyright 2019 The Knative Authors

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

package samples

import (
	"context"
	brokerchannelreconciler "github.com/cowbon/brokerchannel/pkg/client/injection/reconciler/samples/v1alpha1/brokerchannel"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	brokerchannelinformer "github.com/cowbon/brokerchannel/pkg/client/injection/informers/samples/v1alpha1/brokerchannel"
	"knative.dev/pkg/injection/clients/dynamicclient"
	"knative.dev/pkg/resolver"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logging.FromContext(ctx).Error("Start running")
	brokerChannelInformer := brokerchannelinformer.Get(ctx)

	r := &Reconciler{
		dynamicClientSet:   dynamicclient.Get(ctx),
	}
	impl := brokerchannelreconciler.NewImpl(ctx, r)
	r.sinkResolver = resolver.NewURIResolver(ctx, impl.EnqueueKey)
	logging.FromContext(ctx).Info("Setting up event handlers")
	brokerChannelInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))


	return impl
}
