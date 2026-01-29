# Quick Linux Setup

This is a short, Linux-focused setup for developing and running the testnet locally.

Prereqs
- Linux (Ubuntu 22.04+ recommended)
- 8GB RAM (16GB recommended)
- Internet access

1) Install Go (>= 1.25)

Replace `GO_VERSION` with the latest 1.25+ patch as needed.

```bash
# remove old Go (if present)
sudo rm -rf /usr/local/go

GO_VERSION=1.25.4
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
rm go${GO_VERSION}.linux-amd64.tar.gz

# Add to shell profile (~/.profile, ~/.bashrc or ~/.zshrc)
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"' >> ~/.profile
source ~/.profile

# verify
go version
```

2) Install Ignite CLI (version 28.6.1)

Install the requested Ignite version. Replace `28.6.1` if you need a different release.

```bash
IGNITE_VERSION=28.6.1
# Example download (Linux AMD64). Adjust URL if the project uses a different release naming.
curl -L "https://github.com/ignite/cli/releases/download/v${IGNITE_VERSION}/ignite-${IGNITE_VERSION}-linux-amd64.tar.gz" -o ignite.tar.gz
tar -xzf ignite.tar.gz
sudo mv ignite /usr/local/bin/ || sudo mv ignite* /usr/local/bin/ || true
rm ignite.tar.gz

# or via Go (alternate)
go install github.com/ignite/cli/v28@v${IGNITE_VERSION} || true

ignite version
```

3) Clone repository and get dependencies

```bash
git clone https://github.com/rohitroyrr8/restakechain-go-testnet.git
cd restakechain-go-testnet
go mod download
```

4) Build / install project binary

This project contains a Makefile; use it if available.

```bash
# preferred (uses project Makefile)
make install

# or direct Go install for the CLI binary
go install ./cmd/testchaind

# or generate using ignite
ignite chain build

# quick verification
testchaind version || echo "testchaind not in PATH; check $GOPATH/bin or Makefile targets"
```

5) Quick verification

```bash
go version        # should be 1.25+
ignite version    # should be 28.6.1 (or the version you installed)
testchaind version
```

Notes
- Keep `$HOME/go/bin` and `/usr/local/go/bin` in your `PATH`.
- If a specific Ignite release name differs on GitHub, adjust the download URL accordingly.
- This doc intentionally kept short — ask if you want expanded steps for running a local multi-node testnet or CI.


## System Requirements

- **OS**: Linux, macOS, or Windows (WSL2 recommended)
- **RAM**: 8GB minimum (16GB recommended)
- **Disk Space**: 10GB minimum
- **Internet**: Stable connection required

## Getting Help

- **Cosmos SDK Docs**: https://docs.cosmos.network/
- **Ignite CLI Docs**: https://docs.ignite.com/
- **Go Documentation**: https://golang.org/doc/
- **Project Issues**: https://github.com/rohitroyrr8/restakechain-go-testnet/issues
- **Cosmos Discord**: https://discord.gg/cosmosnetwork

---

## Next Steps

After setting up the environment:

1. Read [challenge.md](./challenge.md) for project requirements
2. Check [README.md](./readme.md) for project overview
3. Review [GOVERNANCE.md](./GOVERNANCE.md) for governance operations
4. Start the local testnet with `ignite chain serve`
5. Begin development and testing

---

**Last Updated**: January 2026  
**Tested On**: Go 1.25.4, Ignite v28.6.1, Ubuntu 24.04 LTS
