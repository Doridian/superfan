{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        package = pkgs.buildGoModule {
          pname = "superfan";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-E2AmHtnZ2+Q6vUXO4niMYdT+8saGTMMa4lEDkSB/jO4=";
          buildInputs = [];
        };
      in
      {
        packages = {
          default = package;
          superfan = package;
        };
      });
}
