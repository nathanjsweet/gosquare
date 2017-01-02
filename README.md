# gosquare

##Summary
Square API for Golang. See the [api documentation](https://docs.connect.squareup.com/api/connect/v1/) for further details.
See the [go docs](https://godoc.org/github.com/nathanjsweet/gosquare) for details on this lib.

##Notes
Every method in the the Connect V1 API is accounted for in the lib.
Additionally every method has a corresponding `*BatchRequest` method,
that returns a `*BatchRequest` object that can be added to an array
and sent as a BatchRequest using the `SubmitBatchRequest` method.

There are several utilities and functions you should be aware of for your benefit:
1. Square will sometimes paginate results on large get request. On any method for
which this is possible, the method will return a `NextRequest` object, which
will be `nil` if there is no "next" result but will be present when there is.
The `NextRequest` object has two convenient methods `GetNextRequest`, which
takes an expected result object as an argument, and will, itself, return another
`NextRequest` object if there is an additional one, and `GetNextRequestAsBatchRequest`,
which will return the request as a `BatchRequest` object which can be sent along
with other `BatchRequest`s in `SubmitBatchRequest`. Finally if any `BatchRequest`
object has pagination the `BatchResponse` object will be populated with a
`NextRequest` member, which can be used the same as discussed above.

2. `GeneratePermissionURL` is a method you can use to generate a url, based
on your square client id, that will give redirect anyone who clicks it to
a square signin screen asking if they would like to give your application permission
to access their data (along with the specific permissions you ask for in the `scope`
argument). Should the user grant your application permission you can also add a redirect
url in the method (the `locale` argument) along with any state you'd like to transfer
between requests (the `state` argument). See [Square's Oauth docs](https://docs.connect.squareup.com/api/oauth/) for details.

3. The `GetToken` method gets a first time token based on the authorization code
you get when someone grants your application permission to access their square account.
The `RenewToken` method, obviously, allows you to renew the token based on the expired token
(the grace period for an old, expired token being able to get your application a renewed token,
is 30 days).

4. `ValidateWebhook` validates a Square-initiated webhook request to your application
is.

