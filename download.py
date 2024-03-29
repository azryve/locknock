import json
import platform
import urllib.request
import sys
import argparse

REPO_URL = "https://api.github.com/repos/azryve/locknock/releases/latest"

def get_os_and_arch():
    os_name = platform.system().lower()
    arch_name = platform.machine().lower()
    if os_name not in ["linux", "darwin"]:
        print(f"Unsupported OS: {os_name}", file=sys.stderr)
        sys.exit(1)
    if arch_name == "x86_64":
        arch_name = "amd64"
    elif arch_name in ["arm64", "aarch64"]:
        arch_name = "arm64"
    else:
        print(f"Unsupported architecture: {arch_name}", file=sys.stderr)
        sys.exit(1)
    return os_name, arch_name

def download_latest_release(os_name, arch_name, output_path):
    asset_name = f"locknock_{os_name}_{arch_name}"
    print(f"Fetching: {REPO_URL}", file=sys.stderr)
    try:
        with urllib.request.urlopen(REPO_URL) as response:
            data = json.loads(response.read().decode())
            assets = data.get("assets", [])
    except Exception as e:
        print(f"Error fetching latest release: {e}", file=sys.stderr)
        sys.exit(1)
    assets_found = {x["name"]: x for x in assets}
    if asset_name not in assets_found:
        print(f"No matching asset '{asset_name}' found for download.", file=sys.stderr)
        return
    asset = assets_found[asset_name]
    download_url = asset["browser_download_url"]
    print(f"Downloading: {download_url}", file=sys.stderr)
    output_file_path = output_path if output_path else asset["name"]
    urllib.request.urlretrieve(download_url, output_file_path)
    print(f"Downloaded to {output_file_path} successfully.")

def parse_args():
    parser = argparse.ArgumentParser(description="Download the latest release binary.")
    parser.add_argument('-o', '--output', type=str, help='Output path for the downloaded binary', default='./locknock')
    args, unknown = parser.parse_known_args()
    return args.output

if __name__ == "__main__":
    output_path = parse_args()
    os_name, arch_name = get_os_and_arch()
    download_latest_release(os_name, arch_name, output_path)
