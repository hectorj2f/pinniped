// Copyright 2020-2021 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// StrategyType enumerates a type of "strategy" used to implement credential access on a cluster.
// +kubebuilder:validation:Enum=KubeClusterSigningCertificate;ImpersonationProxy
type StrategyType string

// FrontendType enumerates a type of "frontend" used to provide access to users of a cluster.
// +kubebuilder:validation:Enum=TokenCredentialRequestAPI;ImpersonationProxy
type FrontendType string

// StrategyStatus enumerates whether a strategy is working on a cluster.
// +kubebuilder:validation:Enum=Success;Error
type StrategyStatus string

// StrategyReason enumerates the detailed reason why a strategy is in a particular status.
// +kubebuilder:validation:Enum=Listening;Pending;Disabled;ErrorDuringSetup;CouldNotFetchKey;CouldNotGetClusterInfo;FetchedKey
type StrategyReason string

const (
	KubeClusterSigningCertificateStrategyType = StrategyType("KubeClusterSigningCertificate")
	ImpersonationProxyStrategyType            = StrategyType("ImpersonationProxy")

	TokenCredentialRequestAPIFrontendType = FrontendType("TokenCredentialRequestAPI")
	ImpersonationProxyFrontendType        = FrontendType("ImpersonationProxy")

	SuccessStrategyStatus = StrategyStatus("Success")
	ErrorStrategyStatus   = StrategyStatus("Error")

	ListeningStrategyReason              = StrategyReason("Listening")
	PendingStrategyReason                = StrategyReason("Pending")
	DisabledStrategyReason               = StrategyReason("Disabled")
	ErrorDuringSetupStrategyReason       = StrategyReason("ErrorDuringSetup")
	CouldNotFetchKeyStrategyReason       = StrategyReason("CouldNotFetchKey")
	CouldNotGetClusterInfoStrategyReason = StrategyReason("CouldNotGetClusterInfo")
	FetchedKeyStrategyReason             = StrategyReason("FetchedKey")
)

// CredentialIssuerStatus describes the status of the Concierge.
type CredentialIssuerStatus struct {
	// List of integration strategies that were attempted by Pinniped.
	Strategies []CredentialIssuerStrategy `json:"strategies"`

	// Information needed to form a valid Pinniped-based kubeconfig using this credential issuer.
	// This field is deprecated and will be removed in a future version.
	// +optional
	KubeConfigInfo *CredentialIssuerKubeConfigInfo `json:"kubeConfigInfo,omitempty"`
}

// CredentialIssuerKubeConfigInfo provides the information needed to form a valid Pinniped-based kubeconfig using this credential issuer.
// This type is deprecated and will be removed in a future version.
type CredentialIssuerKubeConfigInfo struct {
	// The K8s API server URL.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^https://|^http://`
	Server string `json:"server"`

	// The K8s API server CA bundle.
	// +kubebuilder:validation:MinLength=1
	CertificateAuthorityData string `json:"certificateAuthorityData"`
}

// CredentialIssuerStrategy describes the status of an integration strategy that was attempted by Pinniped.
type CredentialIssuerStrategy struct {
	// Type of integration attempted.
	Type StrategyType `json:"type"`

	// Status of the attempted integration strategy.
	Status StrategyStatus `json:"status"`

	// Reason for the current status.
	Reason StrategyReason `json:"reason"`

	// Human-readable description of the current status.
	// +kubebuilder:validation:MinLength=1
	Message string `json:"message"`

	// When the status was last checked.
	LastUpdateTime metav1.Time `json:"lastUpdateTime"`

	// Frontend describes how clients can connect using this strategy.
	Frontend *CredentialIssuerFrontend `json:"frontend,omitempty"`
}

// CredentialIssuerFrontend describes how to connect using a particular integration strategy.
type CredentialIssuerFrontend struct {
	// Type describes which frontend mechanism clients can use with a strategy.
	Type FrontendType `json:"type"`

	// TokenCredentialRequestAPIInfo describes the parameters for the TokenCredentialRequest API on this Concierge.
	// This field is only set when Type is "TokenCredentialRequestAPI".
	TokenCredentialRequestAPIInfo *TokenCredentialRequestAPIInfo `json:"tokenCredentialRequestInfo,omitempty"`

	// ImpersonationProxyInfo describes the parameters for the impersonation proxy on this Concierge.
	// This field is only set when Type is "ImpersonationProxy".
	ImpersonationProxyInfo *ImpersonationProxyInfo `json:"impersonationProxyInfo,omitempty"`
}

// TokenCredentialRequestAPIInfo describes the parameters for the TokenCredentialRequest API on this Concierge.
type TokenCredentialRequestAPIInfo struct {
	// Server is the Kubernetes API server URL.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^https://|^http://`
	Server string `json:"server"`

	// CertificateAuthorityData is the base64-encoded Kubernetes API server CA bundle.
	// +kubebuilder:validation:MinLength=1
	CertificateAuthorityData string `json:"certificateAuthorityData"`
}

// ImpersonationProxyInfo describes the parameters for the impersonation proxy on this Concierge.
type ImpersonationProxyInfo struct {
	// Endpoint is the HTTPS endpoint of the impersonation proxy.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^https://`
	Endpoint string `json:"endpoint"`

	// CertificateAuthorityData is the base64-encoded PEM CA bundle of the impersonation proxy.
	// +kubebuilder:validation:MinLength=1
	CertificateAuthorityData string `json:"certificateAuthorityData"`
}

// CredentialIssuer describes the configuration and status of the Pinniped Concierge credential issuer.
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=pinniped,scope=Cluster
// +kubebuilder:subresource:status
type CredentialIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Status of the credential issuer.
	// +optional
	Status CredentialIssuerStatus `json:"status"`
}

// CredentialIssuerList is a list of CredentialIssuer objects.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CredentialIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []CredentialIssuer `json:"items"`
}
