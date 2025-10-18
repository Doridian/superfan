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
          vendorHash =  "sha256-7xNRovpxxS1C2Y4kOS61zyur0DcWoJpue/0lZLNWpCA=";
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
