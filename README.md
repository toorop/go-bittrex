go-bittrex
==========

Go API for the Bittrex crypto-currency exchange.

##Warning
Bittrex API presents somes problems (see below), due to lack of support from the bittrex team (bug report sent but no response), i've decided to stop the developpement of this lib.

### Problems with Bittrex 
#### Not RFC compliant

According to RFC 2616 :
http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2

"Each header field consists of a name followed by a colon (":") and the field value. Field names are case-insensitive."

But it's not the case,with Bittrex API we must use the uppercase form "APIKEY".


curl -H "APIKEY: XXXXXXX" 'https://bittrex.com/api/v1.1/account/getorderhistory?count=10&market=BTC-QBC'

{"success":true,"message":"","result":[{"OrderUuid":"d93ffdc7-f11f-4c2c-83ca-ea01f756d2f2","Exchange":"BTC-QBC","TimeStamp":"2014-04-22T19:15:22.807","OrderType":"LIMIT_SELL","Limit":0.00055000,"Quantity":14.92518703,"QuantityRemaining":14.92518703,"Commission":0.00000000,"Price":0.00000000,"PricePerUnit":null},{"OrderUuid":"fa3f1b8e-41e7-441c-8132-0ccf42f4636f","Exchange":"BTC-QBC","TimeStamp":"2014-04-22T09:37:57.18","OrderType":"LIMIT_BUY","Limit":0.00040000,"Quantity":14.92518703,"QuantityRemaining":0.00000000,"Commission":0.00001492,"Price":0.00597007,"PricePerUnit":0.00039999},{"OrderUuid":"e2d0fd10-fd30-44b1-bd4c-6e1136762731","Exchange":"BTC-QBC","TimeStamp":"2014-04-22T07:06:29.51","OrderType":"LIMIT_SELL","Limit":0.00060000,"Quantity":10.00000000,"QuantityRemaining":0.00000000,"Commission":0.00001500,"Price":0.00600000,"PricePerUnit":0.00060000}]}


But :

curl -H "Apikey: XXXXXXX" -v 'https://bittrex.com/api/v1.1/account/getorderhistory?count=10&market=QBC'

{"success":false,"message":"Value cannot be null.\r\nParameter name: source","result":null}


That's a problems because a lot of standards libraries rewrite header's name by following this rule : converts the first letter and any letter following a hyphen to upper case; the rest are converted to lowercase.

It's the case for the std "net" libary in Go:
http://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey 

I could hack this, but it would be better if Bittrex correct its API.

#### It's seems that they don't know where they are going with the API
For example, there is 2 authentications procedures, depending on the method called. And they've announced a third one...

#### Doc needs updates
See for example /account/getdeposithistory (new in v1.1)

Request is just a copy & past of getwithdrawalhistory

If you try a /account/getdeposithistory you will have a 404

#### Lack of support
A 5 days ago i've send an email about RFC compliance problem. No response yet.


I hope all this problems will be fixed soon.
Wake up guys !