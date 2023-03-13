/*
   Copyright The containerd Authors.
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

package proxy

import (
	"context"

	verifierapi "github.com/containerd/containerd/api/services/verifier/v1"
	"github.com/containerd/containerd/verifier"
)

func NewVerifier(client verifierapi.VerifierClient) verifier.Verifier {
	return &proxyVerifier{
		client: client,
	}
}

type proxyVerifier struct {
	client verifierapi.VerifierClient
}

func (p *proxyVerifier) VerifyImage(ctx context.Context, name string, digest string) (verifier.Judgement, error) {
	resp, err := p.client.VerifyImage(ctx, &verifierapi.VerifyImageRequest{
		ImageName:   name,
		ImageDigest: digest,
	})
	if err != nil {
		// TODO: fail open option?
		return verifier.Judgement{}, err
	}

	return verifier.Judgement{
		OK:     resp.OK,
		Reason: resp.Reason,
	}, nil
}
