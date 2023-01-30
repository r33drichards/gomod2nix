{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
}:

let
  goEnv = pkgs.mkGoEnv { pwd = ./.; };
in
pkgs.mkShell {
  packages = [
    goEnv
    pkgs.gomod2nix
    pkgs.gopls
    pkgs.tmux
    pkgs.gofumpt
    pkgs.gosec
    pkgs.delve
    pkgs.go-tools
    pkgs.gotests
    pkgs.gomodifytags
    pkgs.terraform
    pkgs.just
  ];
}
