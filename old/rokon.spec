Name:           rokon
Version:        1.0.0
Release:        4{?dist}
Summary:        Control your Roku device with your desktop!

License:        AGPL-3.0-or-later
URL:            https://github.com/BrycensRanch/Rokon
Source:        %{url}/archive/master.tar.gz

BuildRequires:  git
BuildRequires:  go
BuildRequires:  rust
BuildRequires:  cargo
BuildRequires:  cmake
BuildRequires:  gcc
BuildRequires:  gcc-c++
BuildRequires:  make
BuildRequires:  gtk4-devel
BuildRequires:  glib2-devel
BuildRequires:  gobject-introspection-devel

Provides:       %{name} = %{version}-%{release}

%description
Rokon is a GTK4 application that allows you to control your Roku device with your desktop. It is written in Golang and uses the Roku External Control API to communicate with your Roku device.

%global debug_package %{nil}

%prep
%autosetup -n Rokon-master



%build
go mod download all
ls
cd old/lib/sysinfo
cargo build --release
mv ./target/release/liblibrokon_rust_sysinfo.a ../
mv ./target/librokon_rust_sysinfo.h ../
cd ../..
go build -v -o %{name}

%install
install -Dpm 0755 old/%{name} %{buildroot}%{_bindir}/%{name}


%files
%{_bindir}/%{name}
%license LICENSE.md



%changelog
* Mon Sep 2 2024 Brycen <brycengranville@outlook.com> 1.0.0-3
- Initial package