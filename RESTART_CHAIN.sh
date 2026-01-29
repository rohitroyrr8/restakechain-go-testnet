#!/bin/bash
# Script to rebuild and restart the testchain with a clean state

set -e

echo "================================================"
echo "Rebuilding testchaind binary..."
echo "================================================"
go install ./cmd/testchaind

echo ""
echo "================================================"
echo "Removing old chain state..."
echo "================================================"
rm -rf ~/.testchain || true

echo ""
echo "================================================"
echo "Starting fresh chain..."
echo "================================================"
echo ""
echo "The chain will start. Give it a moment to initialize."
echo "Then in another terminal, you can run:"
echo ""
echo "  # Check faucet module account exists"
echo "  testchaind query auth module-accounts"
echo ""
echo "  # Check faucet module balance"
echo "  testchaind query bank balances zig1zynz2d7f0yp5s42pjl3qgp2luw5nlw8uc3lnf5"
echo ""
echo "  # Request from faucet"
echo "  testchaind tx faucet request 100000stake --from alice --chain-id testchain -y"
echo ""
echo "  # Query requests"
echo "  testchaind query faucet by-address zig1qvuhm5m644660nd8377d6l7yz9e9hhm9evmx3x"
echo ""
echo "================================================"

ignite chain serve
