
## Build the binary and start chain
```
rm -rf ~/.testchain
ignite chain serve
```
## funding faucet module

```
testchaind query auth module-accounts

testchaind tx bank send alice zig17s95c5jpc6x2l3edwh4dm8yhac68yru7efu8u9 --chain-id testchain -y

```

## Basic Faucet Request Command
```
testchaind tx faucet request 100000 \
  --from alice \
  --chain-id testchain \
  --yes
```

## Check request status
```
testchaind query faucet params


// check your balance
testchaind query bank balances alice

// chain total requested by an address
testchaind query faucet by-address alice

```

## Common Options
```
--gas auto                  # Estimate gas automatically
--gas-adjustment 1.2       # Increase gas estimate by 20%
--fees 5000          # Custom fees
--broadcast-mode block     # Wait for block confirmation
--output json              # JSON output format
```


## Geting the address
```
# List all accounts
testchaind keys list

# Show alice's full details
testchaind keys show alice

# Show only alice's address
testchaind keys show alice -a

# Show bob's address
testchaind keys show bob -a
```

## Query the address
```
# Query balance
testchaind query bank balances cosmos1...

# Make a faucet request
testchaind tx faucet request 100000 \
  --from alice \
  --chain-id testchain \
  --yes

# Query faucet requests
testchaind query faucet by-address cosmos1...
```