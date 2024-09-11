/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package organizationmember

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/controller/handler"
	"github.com/crossplane/upjet/pkg/metrics"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-equinix/apis/metal/v1alpha1"
	features "github.com/crossplane-contrib/provider-jet-equinix/internal/features"
)

// Setup adds a controller that reconciles OrganizationMember managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.OrganizationMember_GroupVersionKind.String())
	var initializers managed.InitializerChain
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK, connection.WithTLSConfig(o.ESSOptions.TLSConfig)))
	}
	eventHandler := handler.NewEventHandler(handler.WithLogger(o.Logger.WithValues("gvk", v1alpha1.OrganizationMember_GroupVersionKind)))
	ac := tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1alpha1.OrganizationMember_GroupVersionKind), tjcontroller.WithEventHandler(eventHandler), tjcontroller.WithStatusUpdates(false))
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(
			tjcontroller.NewTerraformPluginFrameworkAsyncConnector(mgr.GetClient(), o.OperationTrackerStore, o.SetupFn, o.Provider.Resources["equinix_metal_organization_member"],
				tjcontroller.WithTerraformPluginFrameworkAsyncLogger(o.Logger),
				tjcontroller.WithTerraformPluginFrameworkAsyncConnectorEventHandler(eventHandler),
				tjcontroller.WithTerraformPluginFrameworkAsyncCallbackProvider(ac),
				tjcontroller.WithTerraformPluginFrameworkAsyncMetricRecorder(metrics.NewMetricRecorder(v1alpha1.OrganizationMember_GroupVersionKind, mgr, o.PollInterval)),
				tjcontroller.WithTerraformPluginFrameworkAsyncManagementPolicies(o.Features.Enabled(features.EnableBetaManagementPolicies)))),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(tjcontroller.NewOperationTrackerFinalizer(o.OperationTrackerStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3 * time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	}
	if o.PollJitter != 0 {
		opts = append(opts, managed.WithPollJitterHook(o.PollJitter))
	}
	if o.Features.Enabled(features.EnableBetaManagementPolicies) {
		opts = append(opts, managed.WithManagementPolicies())
	}
	if o.MetricOptions != nil {
		opts = append(opts, managed.WithMetricRecorder(o.MetricOptions.MRMetrics))
	}

	// register webhooks for the kind v1alpha1.OrganizationMember
	// if they're enabled.
	if o.StartWebhooks {
		if err := ctrl.NewWebhookManagedBy(mgr).
			For(&v1alpha1.OrganizationMember{}).
			Complete(); err != nil {
			return errors.Wrap(err, "cannot register webhook for the kind v1alpha1.OrganizationMember")
		}
	}

	if o.MetricOptions != nil && o.MetricOptions.MRStateMetrics != nil {
		stateMetricsRecorder := statemetrics.NewMRStateRecorder(
			mgr.GetClient(), o.Logger, o.MetricOptions.MRStateMetrics, &v1alpha1.OrganizationMemberList{}, o.MetricOptions.PollStateMetricInterval,
		)
		if err := mgr.Add(stateMetricsRecorder); err != nil {
			return errors.Wrap(err, "cannot register MR state metrics recorder for kind v1alpha1.OrganizationMemberList")
		}
	}

	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1alpha1.OrganizationMember_GroupVersionKind), opts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(xpresource.DesiredStateChanged()).
		Watches(&v1alpha1.OrganizationMember{}, eventHandler).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
