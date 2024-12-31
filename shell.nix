{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    nodejs_22
    nodePackages.pnpm
    go
  ];
  COMPOSE_FILE="compose-local.yaml";
}
