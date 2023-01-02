# TinyTime

_`Tiny Timestamps That Last Forever.`_

TinyTime lets you choose how precise you want your timestamp to be, enabling shorter TinyTime stamps to be used when length/size is a concern. With TinyTime you can use a 32bit integer timestamp saving 50% of length compared to standard Unix timestamp. This can be achieved through storing **minutes rather than seconds**, or hours instead of minutes, days instead of hours etc. Even though the timestamps are smaller they will still last for thousands of years, something a Unix timestamp will not in a 32bit integer. 

## The Problem This Solves

Communicating a unix timestamp in readable ASCII requires **16 bytes** (in HEX), something which can make stateless tokens, urls or other usages too large, bloating and wasting characters. With TinyTime that same time stamp can be reduced to only **4 bytes** by storing the hour. In most cases down to the second precision is unnecessary. For example, a user token expiration; we only need to know at `X` hour it expires, the exact second is unimportant.

When storing TinyTime as bytes the smallest timestamp can be `2 bytes` or `3 bytes` if stored as base64.

## Install Now
*Theres no TinyTime like the present.*

```sh
go get github.com/go-compile/tinytime
```