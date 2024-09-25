package common

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"

	hdfsv1alpha1 "github.com/zncdatadev/hdfs-operator/api/v1alpha1"
	authv1alpha1 "github.com/zncdatadev/operator-go/pkg/apis/authentication/v1alpha1"
	"github.com/zncdatadev/operator-go/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	oidcLogger = ctrl.Log.WithName("oidc")
)

func MakeOidcContainer(
	ctx context.Context,
	client ctrlclient.Client,
	instance *hdfsv1alpha1.HdfsCluster,
	port int32,
) (*corev1.Container, error) {
	authClass := &authv1alpha1.AuthenticationClass{}
	if err := client.Get(ctx, ctrlclient.ObjectKey{Namespace: instance.Namespace, Name: instance.Spec.ClusterConfigSpec.Authentication.AuthenticationClass}, authClass); err != nil {
		if ctrlclient.IgnoreNotFound(err) != nil {
			return nil, err
		}
		return nil, nil
	}

	if authClass.Spec.AuthenticationProvider.OIDC == nil || instance.Spec.ClusterConfigSpec.Authentication.Oidc == nil {
		oidcLogger.Info("OIDC provider is not configured", "OidcProvider", authClass.Spec.AuthenticationProvider.OIDC, "OidcCredential", instance.Spec.ClusterConfigSpec.Authentication.Oidc)
		return nil, nil
	}

	oidc := NewOidcContainerBuilder(
		client,
		instance,
		authClass.Spec.AuthenticationProvider.OIDC,
		instance.Spec.ClusterConfigSpec.Authentication.Oidc,
		port,
	)
	obj := oidc.Build(oidc)
	return &obj, nil
}

type OidcContainerBuilder struct {
	ContainerBuilder
	client       client.Client
	instanceUid  string
	port         int32
	oidcProvider *authv1alpha1.OIDCProvider
	oidc         *hdfsv1alpha1.OidcSpec
}

func NewOidcContainerBuilder(
	client client.Client,
	instance *hdfsv1alpha1.HdfsCluster,
	oidcProvider *authv1alpha1.OIDCProvider,
	oidc *hdfsv1alpha1.OidcSpec,
	port int32,
) *OidcContainerBuilder {
	imageSpec := instance.Spec.Image
	image := hdfsv1alpha1.TransformImage(imageSpec)
	return &OidcContainerBuilder{
		ContainerBuilder: *NewContainerBuilder(image.String(), image.GetPullPolicy(), corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("512Mi"),
			},
		}),
		client:       client,
		instanceUid:  string(instance.UID),
		port:         port,
		oidcProvider: oidcProvider,
		oidc:         oidc,
	}
}

func (o *OidcContainerBuilder) ContainerName() string {
	return "oidc"
}

func (o *OidcContainerBuilder) ContainerPorts() []corev1.ContainerPort {
	return []corev1.ContainerPort{
		{
			Name:          "oidc",
			ContainerPort: 4180,
		},
	}
}

func (o *OidcContainerBuilder) ContainerEnv() []corev1.EnvVar {

	oidcProvider := o.oidcProvider

	scopes := []string{"openid", "email", "profile"}

	if o.oidc.ExtraScopes != nil {
		scopes = append(scopes, o.oidc.ExtraScopes...)
	}

	issuer := url.URL{
		Scheme: "http",
		Host:   oidcProvider.Hostname,
		Path:   oidcProvider.RootPath,
	}

	if oidcProvider.Port != 0 && oidcProvider.Port != 80 {
		issuer.Host += ":" + strconv.Itoa(oidcProvider.Port)
	}

	provisioner := oidcProvider.Provisioner
	// TODO: fix support keycloak-oidc
	if provisioner == "keycloak" {
		provisioner = "keycloak-oidc"
	}

	clientCredentialsSecretName := o.oidc.ClientCredentialsSecret

	hash := sha256.Sum256([]byte(o.instanceUid))
	hashStr := hex.EncodeToString(hash[:])
	tokenBytes := []byte(hashStr[:16])

	cookieSecret := base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString(tokenBytes)))

	return []corev1.EnvVar{
		{
			Name:  "OAUTH2_PROXY_COOKIE_SECRET",
			Value: cookieSecret,
		},
		{
			Name: "OAUTH2_PROXY_CLIENT_ID",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: clientCredentialsSecretName,
					},
					Key: "CLIENT_ID",
				},
			},
		},
		{
			Name: "OAUTH2_PROXY_CLIENT_SECRET",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: clientCredentialsSecretName,
					},
					Key: "CLIENT_SECRET",
				},
			},
		},
		{
			Name: "POD_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name:  "OAUTH2_PROXY_OIDC_ISSUER_URL",
			Value: issuer.String(),
		},
		{
			Name:  "OAUTH2_PROXY_SCOPE",
			Value: strings.Join(scopes, " "),
		},
		{
			Name:  "OAUTH2_PROXY_PROVIDER",
			Value: provisioner,
		},
		{
			Name:  "UPSTREAM",
			Value: "http://$(POD_IP):" + strconv.Itoa(int(o.port)),
		},
		{
			Name:  "OAUTH2_PROXY_HTTP_ADDRESS",
			Value: "0.0.0.0:4180",
		},
		{
			Name:  "OAUTH2_PROXY_REDIRECT_URL",
			Value: "http://localhost:4180/oauth2/callback",
		},
		{
			Name:  "OAUTH2_PROXY_CODE_CHALLENGE_METHOD",
			Value: "S256",
		},
		{
			Name:  "OAUTH2_PROXY_EMAIL_DOMAINS",
			Value: "*",
		},
	}

}

func (o *OidcContainerBuilder) Command() []string {
	script := `
	set -ex
	ARCH=$(uname -m)
	ARCH="${ARCH/x86_64/amd64}"
	ARCH="${ARCH/aarch64/arm64}"
	
	mkdir /kubedoop/oauth2-proxy
	cd /kubedoop/oauth2-proxy
	
	curl -sSfL \
		https://github.com/oauth2-proxy/oauth2-proxy/releases/download/v7.6.0/oauth2-proxy-v7.6.0.linux-${ARCH}.tar.gz \
		| tar xzf - --strip-components=1
	
	/kubedoop/oauth2-proxy/oauth2-proxy \
		--upstream=${UPSTREAM}
	`
	return []string{
		"bash",
		"-c",
		util.IndentTab4Spaces(script),
	}
}
