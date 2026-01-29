# Governance: Updating Faucet Parameters

This guide explains how to create and vote on governance proposals to update the faucet module's parameters (`maxPerRequest` and `maxPerAddress`) on the testnet.

## Parameters Overview

- **maxPerRequest**: Maximum amount of tokens a user can request in a single transaction
- **maxPerAddress**: Maximum total amount of tokens a user can request across all transactions (lifetime limit)
- **defaultDenom**: The denomination of tokens to distribute (default: "stake")
- **running**: Boolean flag to enable/disable the faucet module

## Prerequisites

Before submitting a proposal, ensure you have:

1. **testchaind** binary installed and in your PATH
2. Enough balance to pay for the proposal deposit (typically 10 STAKE)
3. Access to a node (local or remote)
4. The governance module enabled on the chain

## Step 1: Create a Governance Proposal

### Option A: Using MsgUpdateParams Message (Recommended)

Create a proposal JSON file with the new parameter values:

**File: `update_faucet_params.json`**

```json
{
  "title": "Update Faucet Module Parameters",
  "description": "Proposal to update maxPerRequest and maxPerAddress limits for the faucet module",
  "changes": [
    {
      "subspace": "faucet",
      "key": "MaxPerRequest",
      "value": "100000"
    },
    {
      "subspace": "faucet",
      "key": "MaxPerAddress",
      "value": "1000000"
    }
  ],
  "deposit": "10000000stake"
}
```

**Parameter Values:**
- `maxPerRequest`: Set to the maximum tokens allowed per single request (e.g., `100000`)
- `maxPerAddress`: Set to the maximum total tokens across all requests (e.g., `1000000`)
- `deposit`: Amount of tokens locked as proposal deposit (typically 10 STAKE = 10000000 smallest units)

### Step 2: Submit the Proposal

Submit the governance proposal using the testchain CLI:

```bash
testchaind tx gov submit-proposal update-params \
  --from <your-account> \
  --params '{"subspace":"faucet","key":"MaxPerRequest","value":"100000"}' \
  --params '{"subspace":"faucet","key":"MaxPerAddress","value":"1000000"}' \
  --title "Update Faucet Module Parameters" \
  --description "Proposal to update maxPerRequest and maxPerAddress limits for the faucet module" \
  --deposit 10000000stake \
  --chain-id testchain \
  --node tcp://localhost:26657
```

Or using a proposal file:

```bash
testchaind tx gov submit-proposal \
  update-params \
  update_faucet_params.json \
  --from <your-account> \
  --chain-id testchain \
  --node tcp://localhost:26657
```

**Expected Output:**
You'll see a transaction hash. Wait for confirmation and note the **proposal ID**.

## Step 3: Query the Proposal

Check the proposal details:

```bash
testchaind query gov proposal <proposal-id> \
  --node tcp://localhost:26657
```

**Output will show:**
- Proposal ID
- Title and description
- Current status (Voting Period)
- Voting end time
- Deposit amount

## Step 4: Vote on the Proposal

### Vote YES (Support the proposal)

```bash
testchaind tx gov vote <proposal-id> yes \
  --from <your-account> \
  --chain-id testchain \
  --node tcp://localhost:26657
```

### Vote NO (Reject the proposal)

```bash
testchaind tx gov vote <proposal-id> no \
  --from <your-account> \
  --chain-id testchain \
  --node tcp://localhost:26657
```

### Vote ABSTAIN

```bash
testchaind tx gov vote <proposal-id> abstain \
  --from <your-account> \
  --chain-id testchain \
  --node tcp://localhost:26657
```

## Step 5: Check Voting Status

```bash
testchaind query gov votes <proposal-id> \
  --node tcp://localhost:26657
```

## Step 6: Wait for Voting Period to End

The voting period is configured in the chain's governance settings (typically 1 week on mainnet, shorter on testnet).

Once the voting period ends, the proposal will be:
- **PASSED**: If >50% votes (quorum met) and more YES than NO votes
- **REJECTED**: If quorum not met or more NO votes than YES

## Step 7: Execute the Proposal

Once passed, the parameters are automatically updated. Verify the new values:

```bash
testchaind query faucet params \
  --node tcp://localhost:26657
```

**Expected Output:**
```yaml
params:
  maxPerAddress: "1000000"
  maxPerRequest: "100000"
  running: true
  defaultDenom: stake
```

## Example Workflow

### Local Testnet Example

```bash
# 1. Submit proposal
PROPOSAL_ID=$(testchaind tx gov submit-proposal update-params \
  --from alice \
  --params '{"subspace":"faucet","key":"MaxPerRequest","value":"50000"}' \
  --params '{"subspace":"faucet","key":"MaxPerAddress","value":"500000"}' \
  --title "Increase Faucet Limits" \
  --description "Increase daily faucet limits for better user experience" \
  --deposit 10000000stake \
  --chain-id testchain \
  --yes \
  -o json | jq -r '.logs[0].events[] | select(.type=="submit_proposal") | .attributes[] | select(.key=="proposal_id") | .value')

echo "Proposal ID: $PROPOSAL_ID"

# 2. Vote
testchaind tx gov vote $PROPOSAL_ID yes \
  --from alice \
  --chain-id testchain \
  --yes

# 3. Check status (after voting period)
testchaind query gov proposal $PROPOSAL_ID

# 4. Verify new parameters
testchaind query faucet params
```

## Configuration Parameters Reference

All faucet module parameters:

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `maxPerRequest` | uint64 | 0 | Max tokens per request (0 = unlimited) |
| `maxPerAddress` | uint64 | 0 | Max tokens per address lifetime (0 = unlimited) |
| `defaultDenom` | string | "stake" | Token denomination to distribute |
| `running` | bool | false | Enable/disable faucet module |

## Governance Constants

| Setting | Value | Notes |
|---------|-------|-------|
| Min Deposit | 10 STAKE | Required to activate voting |
| Voting Period | Chain-dependent | Check chain params |
| Quorum | 33.34% | Minimum participation required |
| Threshold | 50% | YES votes needed to pass |

## Troubleshooting

### Proposal Not Found
```bash
# Check if proposal exists
testchaind query gov proposals --node tcp://localhost:26657

# Check specific proposal with full details
testchaind query gov proposal <id> -o json | jq .
```

### Insufficient Balance for Deposit
```bash
# Check your balance
testchaind query bank balances <your-address>

# Get tokens from faucet first
testchaind tx faucet request \
  --amount 1000000stake \
  --from alice \
  --chain-id testchain
```

### Transaction Failed
```bash
# Check logs
testchaind query tx <tx-hash>

# Common issues:
# - Insufficient gas
# - Wrong chain-id
# - Account doesn't exist
# - Invalid parameter values
```

## Best Practices

1. **Test on Testnet First**: Always test governance changes on testnet before mainnet
2. **Clear Description**: Provide detailed rationale in proposal description
3. **Community Discussion**: Discuss with community before submitting
4. **Monitor Voting**: Track voting progress and adjust if needed
5. **Parameter Validation**: Ensure new values make sense:
   - `maxPerRequest` ≤ `maxPerAddress`
   - Values should not be 0 unless intentionally disabling limits
   - Consider token supply when setting limits

## Additional Resources

- [Cosmos SDK Governance Docs](https://docs.cosmos.network/main/build/modules/gov)
- [Chain Governance Parameters](./config.yml)
- [Faucet Module Source](./x/faucet)
