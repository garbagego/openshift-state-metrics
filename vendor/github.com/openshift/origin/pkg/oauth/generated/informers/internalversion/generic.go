// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"fmt"

	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=oauth.openshift.io, Version=internalVersion
	case oauth.SchemeGroupVersion.WithResource("oauthaccesstokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Oauth().InternalVersion().OAuthAccessTokens().Informer()}, nil
	case oauth.SchemeGroupVersion.WithResource("oauthauthorizetokens"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Oauth().InternalVersion().OAuthAuthorizeTokens().Informer()}, nil
	case oauth.SchemeGroupVersion.WithResource("oauthclients"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Oauth().InternalVersion().OAuthClients().Informer()}, nil
	case oauth.SchemeGroupVersion.WithResource("oauthclientauthorizations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Oauth().InternalVersion().OAuthClientAuthorizations().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
