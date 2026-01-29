```
We have prepared a new shiny blockchain using ignite v28.6.1 and created one module, a
faucet. Please modify the faucet module with one request message and one query so it
is possible to request tokens from the faucet, received in the requested address and
to query the requests already done to the faucet per address.

Take into consideration the following:
- The faucet works as a transfer faucet. Coins are not minted.
- The request message has one argument, “amount” which is uint - that is the amount of
stake coins being sent to the requester (signer); it should check at least against
module parafmeters maxPerRequest and maxPerAddress (address is signer). If any of the
limits is reached, it should return an error. Denom is not needed, only default denom
is used.

- Params maxPerRequest and maxPerAddress are modifiable via governance. Default values
are set up on the code.
- maxPerAddress refers to the maximum in total requested by the user in all time

- The query should be called by Address and takes the provided address and returns the
requests done by the address indicating the amount and the height when the request was
done.
- Anything added additionally will be evaluated positively.
```