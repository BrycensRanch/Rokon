# Why is boss music playing?
FROM fedora:rawhide

# We need to create a non-root user

# Install the necessary packages

RUN dnf copr enable -y atim/starship && dnf copr enable elxreno/zsync -y

RUN dnf up -y --allowerasing

RUN dnf install -y golang gtk4-devel gobject-introspection-devel git @development-tools @development-libs @c-development @rpm-development-tools flatpak-builder apt-get python3-pip sudo starship fedora-packager upx patchelf fuse clang go-rpm-macros yq wget

RUN           wget -c https://github.com/$(wget -q https://github.com/probonopd/go-appimage/releases/expanded_assets/continuous -O - | grep "appimagetool-.*-x86_64.AppImage" | head -n 1 | cut -d '"' -f 2)
RUN  chmod +x appimagetool-*.AppImage
RUN  mv appimagetool-* appimagetool
RUN  mv appimagetool /usr/local/bin
RUN  wget -c https://github.com/$(wget -q https://github.com/probonopd/go-appimage/releases/expanded_assets/continuous -O - | grep "mkappimage-.*-x86_64.AppImage" | head -n 1 | cut -d '"' -f 2)
RUN chmod +x mkappimage-*.AppImage
RUN mv mkappimage-* mkappimage
RUN mv mkappimage /usr/local/bin

# AppImage Builder dependencies
RUN dnf install -y binutils coreutils desktop-file-utils fakeroot fuse patchelf python3-pip python3-setuptools squashfs-tools strace util-linux zsync


# Flatpak dependencies
RUN flatpak remote-add --user --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
RUN flatpak remote-add --user --if-not-exists flathub-beta https://flathub.org/beta-repo/flathub-beta.flatpakrepo
RUN flatpak install --user -y flathub org.flatpak.Builder

RUN pip3 install appimage-builder --break-system-packages

# Set the working directory

RUN useradd -m -s /bin/bash dev
RUN echo "dev ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/dev
# add them to the wheel and sudo groups
RUN usermod -aG wheel dev
RUN groupadd -f sudo && usermod -aG sudo dev

RUN echo 'eval "$(starship init bash)"' >> /home/dev/.bashrc

USER dev

WORKDIR /app

# Copy the source code

COPY --chown=dev:dev . .


# Drop developer into their development environment

CMD ["/usr/bin/bash"]
