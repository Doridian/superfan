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
          vendorHash =  "sha256-JDH/An1bB/TG8ZB3PbYP0ne2jDN3h6qD/+eYAUShSoA=";
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
