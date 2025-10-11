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
          vendorHash = "sha256-8f32OPXLPSmXg4OxCS3U73ctIRraI3TMMez8zsU8/xQ=";
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
