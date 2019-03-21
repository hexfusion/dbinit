FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 AS builder

WORKDIR /go/src/github.com/hexfusion/dbinit

COPY . .

RUN make bin/dbinit

# stage 2
FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base

ENTRYPOINT ["/usr/bin/dbinit"]

COPY --from=builder /go/src/github.com/hexfusion/bin/dbinit /usr/bin/
