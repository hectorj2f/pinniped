// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"k8s.io/klog/v2"

	"go.pinniped.dev/internal/testutil/testlogger"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientauthv1beta1 "k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"

	"go.pinniped.dev/internal/certauthority"
	"go.pinniped.dev/internal/here"
	"go.pinniped.dev/internal/testutil"
	"go.pinniped.dev/pkg/conciergeclient"
)

func TestLoginStaticCommand(t *testing.T) {
	cfgDir := mustGetConfigDir()

	testCA, err := certauthority.New("Test CA", 1*time.Hour)
	require.NoError(t, err)
	tmpdir := testutil.TempDir(t)
	testCABundlePath := filepath.Join(tmpdir, "testca.pem")
	require.NoError(t, ioutil.WriteFile(testCABundlePath, testCA.Bundle(), 0600))

	tests := []struct {
		name             string
		args             []string
		env              map[string]string
		loginErr         error
		conciergeErr     error
		wantError        bool
		wantStdout       string
		wantStderr       string
		wantOptionsCount int
		wantLogs         []string
	}{
		{
			name: "help flag passed",
			args: []string{"--help"},
			wantStdout: here.Doc(`
				Login using a static token

				Usage:
				  static [--token TOKEN] [--token-env TOKEN_NAME] [flags]

				Flags:
				      --concierge-api-group-suffix string     Concierge API group suffix (default "pinniped.dev")
				      --concierge-authenticator-name string   Concierge authenticator name
				      --concierge-authenticator-type string   Concierge authenticator type (e.g., 'webhook', 'jwt')
				      --concierge-ca-bundle-data string       CA bundle to use when connecting to the Concierge
				      --concierge-endpoint string             API base for the Concierge endpoint
				      --credential-cache string               Path to cluster-specific credentials cache ("" disables the cache) (default "` + cfgDir + `/credentials.yaml")
				      --enable-concierge                      Use the Concierge to login
				  -h, --help                                  help for static
				      --token string                          Static token to present during login
				      --token-env string                      Environment variable containing a static token
			`),
		},
		{
			name:      "missing required flags",
			args:      []string{},
			wantError: true,
			wantStderr: here.Doc(`
				Error: one of --token or --token-env must be set
			`),
		},
		{
			name: "missing concierge flags",
			args: []string{
				"--token", "test-token",
				"--enable-concierge",
			},
			wantError: true,
			wantStderr: here.Doc(`
				Error: invalid Concierge parameters: endpoint must not be empty
			`),
		},
		{
			name: "missing env var",
			args: []string{
				"--token-env", "TEST_TOKEN_ENV",
			},
			wantError: true,
			wantStderr: here.Doc(`
				Error: --token-env variable "TEST_TOKEN_ENV" is not set
			`),
		},
		{
			name: "empty env var",
			args: []string{
				"--token-env", "TEST_TOKEN_ENV",
			},
			env: map[string]string{
				"TEST_TOKEN_ENV": "",
			},
			wantError: true,
			wantStderr: here.Doc(`
				Error: --token-env variable "TEST_TOKEN_ENV" is empty
			`),
		},
		{
			name: "env var token success",
			args: []string{
				"--token-env", "TEST_TOKEN_ENV",
			},
			env: map[string]string{
				"TEST_TOKEN_ENV": "test-token",
			},
			wantStdout: `{"kind":"ExecCredential","apiVersion":"client.authentication.k8s.io/v1beta1","spec":{"interactive":false},"status":{"token":"test-token"}}` + "\n",
		},
		{
			name: "concierge failure",
			args: []string{
				"--token", "test-token",
				"--enable-concierge",
				"--concierge-endpoint", "https://127.0.0.1/",
				"--concierge-authenticator-type", "webhook",
				"--concierge-authenticator-name", "test-authenticator",
			},
			conciergeErr: fmt.Errorf("some concierge error"),
			env:          map[string]string{"PINNIPED_DEBUG": "true"},
			wantError:    true,
			wantStderr: here.Doc(`
				Error: could not complete Concierge credential exchange: some concierge error
			`),
			wantLogs: []string{"\"level\"=0 \"msg\"=\"Pinniped login: exchanging static token for cluster credential\"  \"authenticator name\"=\"test-authenticator\" \"authenticator type\"=\"webhook\" \"endpoint\"=\"https://127.0.0.1/\""},
		},
		{
			name: "invalid API group suffix",
			args: []string{
				"--token", "test-token",
				"--enable-concierge",
				"--concierge-api-group-suffix", ".starts.with.dot",
				"--concierge-authenticator-type", "jwt",
				"--concierge-authenticator-name", "test-authenticator",
				"--concierge-endpoint", "https://127.0.0.1:1234/",
			},
			wantError: true,
			wantStderr: here.Doc(`
				Error: invalid Concierge parameters: invalid API group suffix: a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')
			`),
		},
		{
			name: "static token success",
			args: []string{
				"--token", "test-token",
			},
			env:        map[string]string{"PINNIPED_DEBUG": "true"},
			wantStdout: `{"kind":"ExecCredential","apiVersion":"client.authentication.k8s.io/v1beta1","spec":{"interactive":false},"status":{"token":"test-token"}}` + "\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testLogger := testlogger.NewLegacy(t) //nolint: staticcheck  // old test with lots of log statements
			klog.SetLogger(testLogger.Logger)
			cmd := staticLoginCommand(staticLoginDeps{
				lookupEnv: func(s string) (string, bool) {
					v, ok := tt.env[s]
					return v, ok
				},
				exchangeToken: func(ctx context.Context, client *conciergeclient.Client, token string) (*clientauthv1beta1.ExecCredential, error) {
					require.Equal(t, token, "test-token")
					if tt.conciergeErr != nil {
						return nil, tt.conciergeErr
					}
					return &clientauthv1beta1.ExecCredential{
						TypeMeta: metav1.TypeMeta{
							Kind:       "ExecCredential",
							APIVersion: "client.authentication.k8s.io/v1beta1",
						},
						Status: &clientauthv1beta1.ExecCredentialStatus{
							Token: "exchanged-token",
						},
					}, nil
				},
			})
			require.NotNil(t, cmd)

			var stdout, stderr bytes.Buffer
			cmd.SetOut(&stdout)
			cmd.SetErr(&stderr)
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			if tt.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wantStdout, stdout.String(), "unexpected stdout")
			require.Equal(t, tt.wantStderr, stderr.String(), "unexpected stderr")

			require.Equal(t, tt.wantLogs, testLogger.Lines())
		})
	}
}
