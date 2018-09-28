# Crypto Utility

## USAGE:
---
#### Example
```
TOKEN=0c8d3980ba92fef743433c024f145fd0
```

#### Encrypt my token 0c8d3980ba92fef743433c024f145fd0 with my RSA public key:
```shell
docker run --rm -v $(pwd)/key.pub:/key.pub edimarlnx/go-crypto-util enc 0c8d3980ba92fef743433c024f145fd0
```
#### Encrypt Result
```
FntlfpBFGAg8JvFX6quVfaCwgLDOwSO7+YFrKPgmlYW/nJ+Vef01Rc1frk8qiTIp7wLsRxX4+nlujQpVnwsfsufnl91PfxuUnMJflve4DgXykPw70PEPacg4sCeNU0ZXyiajsOerVSzFovSXT+0oDkEFZr6YtxFrgh2s3eVCHvPudXj/jtwkQHJvgjL3fGWqlTb4mHv91m7AuK5naF3UrJhn6HiEWTpTbuN5AaDhA8emPI5LAtrCi1wahtWOmdJevnVEowWXM9wnUzlhmw+/NLTCaMpdKVrK+IU9kYkRCn0/fR546FYgQpymk/drOQb5KqTgYhv+XF/eOeEiSlEe1g==
```

#### Decrypt my token 0c8d3980ba92fef743433c024f145fd0 with my RSA public key:
```shell
docker run --rm -v $(pwd)/privkey.pem:/privkey.pem edimarlnx/go-crypto-util dec FntlfpBFGAg8JvFX6quVfaCwgLDOwSO7+YFrKPgmlYW/nJ+Vef01Rc1frk8qiTIp7wLsRxX4+nlujQpVnwsfsufnl91PfxuUnMJflve4DgXykPw70PEPacg4sCeNU0ZXyiajsOerVSzFovSXT+0oDkEFZr6YtxFrgh2s3eVCHvPudXj/jtwkQHJvgjL3fGWqlTb4mHv91m7AuK5naF3UrJhn6HiEWTpTbuN5AaDhA8emPI5LAtrCi1wahtWOmdJevnVEowWXM9wnUzlhmw+/NLTCaMpdKVrK+IU9kYkRCn0/fR546FYgQpymk/drOQb5KqTgYhv+XF/eOeEiSlEe1g==
```

#### Decrypt Result
```
0c8d3980ba92fef743433c024f145fd0
```
