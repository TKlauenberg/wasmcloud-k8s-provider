// go:generate wit-bindgen-wrpc go -w server --out-dir bindings --package github.com/TKlauenberg/wasmcloud-k8s-provider/bindings wit
package main

import (
	"context"
	"log"

	// server "wasmcloud-k8s-dev/bindings"

	"go.wasmcloud.dev/provider"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatal(err)
	}

	k8sClient := mgr.GetClient()

	provider, err := provider.New("example:kubernetes")
	if err != nil {
		log.Fatal(err)
	}

	// Setup handlers for wasmCloud calls
	provider.Handle("WatchResource", watchResourceHandler(mgr))
	provider.Handle("CreateResource", createResourceHandler(k8sClient))

	go provider.Start()

	if err := mgr.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

// Dynamic watch handler
func watchResourceHandler(mgr manager.Manager) provider.HandlerFunc {
	return func(ctx context.Context, req provider.Request) (provider.Response, error) {
		gvr := schema.GroupVersionResource{
			Group:    req.Body["group"].(string),
			Version:  req.Body["version"].(string),
			Resource: req.Body["resource"].(string),
		}

		if err := setupDynamicWatcher(mgr, gvr); err != nil {
			return nil, err
		}

		return provider.Response{"status": "watcher started"}, nil
	}
}

// Simplified dynamic watcher
func setupDynamicWatcher(mgr manager.Manager, gvr schema.GroupVersionResource) error {
	c, err := controller.NewUnmanaged("dynamic-watcher", mgr, controller.Options{
        Reconciler: reconcile.Func(func(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
            // handle event, forward to wasmCloud actors
            return reconcile.Result{}, nil
        }),
    })
    if err != nil {
        return err
    }

    u := &unstructured.Unstructured{}
    u.SetGroupVersionKind(gvr.GroupVersion().WithKind(gvr.Resource))

    if err := c.Watch(&source.Kind{Type: u}, &handler.EnqueueRequestForObject{}); err != nil {
        return err
    }

    go c.Start(context.Background())
    return nil
}

// Handler to create resources (e.g., ConfigMaps)
func createResourceHandler(cli client.Client) provider.HandlerFunc {
	return func(ctx context.Context, req provider.Request) (provider.Response, error) {
		// Implement resource creation logic
		return provider.Response{"status": "resource created"}, nil
	}
}
