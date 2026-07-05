# Copyright 2021 OpenSSF Scorecard Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.26.2@sha256:1e598ea5752ae26c093b746fd73c5095af97d6f2d679c43e83e0eac484a33dc3 AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./

FROM base AS authserver
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 make build-github-server

FROM gcr.io/distroless/base:nonroot@sha256:ab7554b6d07ad354fad31957f8a1a813e65dfb93a8ad160568c79c3f2be6884f
COPY --from=authserver /src/clients/githubrepo/roundtripper/tokens/server/github-auth-server clients/githubrepo/roundtripper/tokens/server/github-auth-server
ENTRYPOINT ["clients/githubrepo/roundtripper/tokens/server/github-auth-server"]
