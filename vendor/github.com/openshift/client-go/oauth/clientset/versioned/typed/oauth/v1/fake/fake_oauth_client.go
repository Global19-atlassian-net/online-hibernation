package fake

import (
	v1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOauthV1 struct {
	*testing.Fake
}

func (c *FakeOauthV1) OAuthAccessTokens() v1.OAuthAccessTokenInterface {
	return &FakeOAuthAccessTokens{c}
}

func (c *FakeOauthV1) OAuthAuthorizeTokens() v1.OAuthAuthorizeTokenInterface {
	return &FakeOAuthAuthorizeTokens{c}
}

func (c *FakeOauthV1) OAuthClients() v1.OAuthClientInterface {
	return &FakeOAuthClients{c}
}

func (c *FakeOauthV1) OAuthClientAuthorizations() v1.OAuthClientAuthorizationInterface {
	return &FakeOAuthClientAuthorizations{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOauthV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
