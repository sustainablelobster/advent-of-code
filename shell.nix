{ pkgs ? import <nixpkgs> { } }:

let
  ruby = pkgs.ruby;
  go = pkgs.go;
  defaultGemDir = "${ruby}/lib/ruby/gems/${ruby.version}";
in
pkgs.mkShell {
  packages = [
    ruby
    pkgs.pre-commit
    go
    pkgs.gopls
    pkgs.gotools
    pkgs.bundler
    pkgs.pkg-config
    pkgs.openssl
    pkgs.zlib
    pkgs.libyaml
    pkgs.libffi
    pkgs.gmp
    pkgs.readline
    pkgs.libxml2
    pkgs.libxslt
    pkgs.postgresql
  ];

  shellHook = ''
    export BUNDLE_PATH=$PWD/.bundle
    export BUNDLE_BIN=$BUNDLE_PATH/bin
    export GEM_HOME=$BUNDLE_PATH
    export GEM_PATH=$GEM_HOME:${defaultGemDir}
    export PATH=$BUNDLE_BIN:$PATH
  '';
}
