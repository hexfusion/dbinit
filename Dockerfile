FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 AS builder

WORKDIR /go/src/github.com/hexfusion/dbutil

COPY . .

RUN make bin/dbutil

# stage 2
FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base

ENTRYPOINT ["/usr/bin/dbutil"]

COPY --from=builder /go/src/github.com/hexfusion/bin/dbutil /usr/bin/
