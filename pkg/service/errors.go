// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"github.com/livekit/psrpc"
)

var (
	ErrEgressNotFound        = psrpc.NewErrorf(psrpc.NotFound, "egress does not exist")
	ErrEgressNotConnected    = psrpc.NewErrorf(psrpc.Internal, "egress not connected (redis required)")
	ErrIdentityEmpty         = psrpc.NewErrorf(psrpc.InvalidArgument, "identity cannot be empty")
	ErrIngressNotConnected   = psrpc.NewErrorf(psrpc.Internal, "ingress not connected (redis required)")
	ErrIngressNotFound       = psrpc.NewErrorf(psrpc.NotFound, "ingress does not exist")
	ErrIngressNonReusable    = psrpc.NewErrorf(psrpc.InvalidArgument, "ingress is not reusable and cannot be modified")
	ErrMetadataExceedsLimits = psrpc.NewErrorf(psrpc.InvalidArgument, "metadata size exceeds limits")
	ErrOperationFailed       = psrpc.NewErrorf(psrpc.Internal, "operation cannot be completed")
	ErrParticipantNotFound   = psrpc.NewErrorf(psrpc.NotFound, "participant does not exist")
	ErrRoomNotFound          = psrpc.NewErrorf(psrpc.NotFound, "requested room does not exist")
	ErrRoomLockFailed        = psrpc.NewErrorf(psrpc.Internal, "could not lock room")
	ErrRoomUnlockFailed      = psrpc.NewErrorf(psrpc.Internal, "could not unlock room, lock token does not match")
	ErrTrackNotFound         = psrpc.NewErrorf(psrpc.NotFound, "track is not found")
	ErrWebHookMissingAPIKey  = psrpc.NewErrorf(psrpc.InvalidArgument, "api_key is required to use webhooks")
)
