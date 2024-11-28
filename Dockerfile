# Maintainer: Brycen G <brycengranville@outlook.com>

# This container *can* work under NVIDIA.
# For example, you could run the command `distrobox create --name rokon --image ghcr.io/brycensranch/rokon --nvidia`

FROM fedora:latest AS builder

RUN dnf install -y \
    make \
    go \
    gtk4-devel \
    gobject-introspection-devel \
    which \
    clang
RUN dnf clean all

# DO WHATEVER IT TAKES TO BUILD AS FAST AS POSSIBLE!!! TO INFINITY... AND BEYOND
ENV CC=clang
ENV CXX=clang++
ENV CFLAGS="-O0 -w -fno-strict-aliasing -gline-tables-only"


WORKDIR /app
COPY . .

RUN make PACKAGED=true TBPKGFMT=detect NOTB=1 tarball

FROM fedora:latest AS runner

WORKDIR /app

COPY --from=builder /app/tarball .

CMD ["./rokon"]
