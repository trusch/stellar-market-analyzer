stellar-market-analyzer
=======================

This tool shows some stats about arbitrary markets in the stellar network.

## Examples

### Get Market Data for ABDT/XLM

``` bash
sma \
    -buying-code native \
    -selling-code ABDT \
    -selling-issuer GDZURZR6RZKIQVOWZFWPVAUBMLLBQGXP2K5E5G7PEOV75IYPDFA36WK4
```

### Get Market Data for ABDT/DOP

``` bash
sma \
    -buying-code DOP \
    -buying-issuer GCURFAOIY7TYEMXXZ2LFSZJ2ALGFFPSU3EDGMFPCWVDESQ2NZWN4GJ6R \
    -selling-code ABDT \
    -selling-issuer GDZURZR6RZKIQVOWZFWPVAUBMLLBQGXP2K5E5G7PEOV75IYPDFA36WK4
```


