{ stdenv, fetchzip, clang, gtk4, go, git, lib, pkgs }:


stdenv.mkDerivation rec {
  pname = "rokon";
  version = "1.0.0";

  src = fetchzip {
    url = "https://nightly.link/BrycensRanch/Rokon/actions/runs/11943137586/rokon-vendored-source.zip";
    sha256 = "sha256-h9MAKYCg2Hff8gk2dGEiDDGtkb1QmfZYN9nYbzHPZG8=";
    stripRoot = false;
  };

  buildInputs = [ clang gtk4 go git ];

  nativeBuildInputs = [];

  configureFlags = [];

  buildPhase = ''
    make build
  '';

  installPhase = ''
    make install PREFIX=$out
  '';

  meta = with lib; {
    description = "Rokon, a custom package for your project";
    license = licenses.agpl3Plus;
    maintainers = with maintainers; [ johnDoe ];
  };
}
