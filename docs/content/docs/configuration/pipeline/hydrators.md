---
title: "Hydrators"
date: 2022-06-09T18:57:10+02:00
lastmod: 2022-06-09T18:57:10+02:00
draft: true
toc: true
menu:
  docs:
    weight: 30
    parent: "Pipeline"
---

Hydrators enrich the information about the subject obtained in the authenticator step with further information, required by either the endpoint of the upstream service itself or an authorizer step. This can be handy if the actual authentication system doesn't have all information about the subject (which is usually the case in microservice architectures), or if dynamic information about the subject, like the current location based on the IP address, is required.

The following section describes the available hydrator types in more detail.

## Hydrator Types

As of today, there is just one hydrator, which is described below.

### Generic

This handler allows you to communicate to any API you want, to fetch further information about the subject. Typical scenarios is getting specific attributes for later authorization purposes which are not known to the authentication system and thus were not made available in subject's `.Attributes` object. If the API responses with a 2xx HTTP response code, the payload made available in the `.Attributes` property of the subject. To avoid overwriting of existing attributes, this object is however not available on the top level, but under a key named by the `id` of the authorizer (See also the example below). If the `Content-Type` of the response is either ending with `json` or is `application/x-www-form-urlencoded`, the payload is decoded and made available as map, otherwise it is treated as string, but, as written above, is made available as well.

To enable the usage of this hydrator, you have to set the `type` property to `generic`.

Configuration using the `config` property is mandatory. Following properties are available:

| Name              | Type                                                        | Mandatory | Overridable | Description                                                                                                                                                                                                                                                                                                                                                                                                                         |
|-------------------|-------------------------------------------------------------|-----------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `endpoint`        | *[Endpoint]({{< ref "configuration_types.md#endpoint">}})*  | yes       | no          | The API of the service providing additional attributes about the authenticated user. At least the `url` must be configured. This hydrator allows templating of the url. By default this authorizer will use HTTP `POST` to send the rendered payload to this endpoint. You can override this behavior by configuring `method`. Depending on the API requirements you might need to configure further properties, like headers, etc. |
| `forward_headers` | *string array*                                              | no        | yes         | If the API requires any headers from the request to Heimdall, you can forward these unchanged by making use of this property.                                                                                                                                                                                                                                                                                                       |
| `forward_cookies` | *string array*                                              | no        | yes         | If the API requires any cookies from the request to Heimdall, you can forward these unchanged by making use of this property                                                                                                                                                                                                                                                                                                        |
| `payload`         | *string*                                                    | no        | yes         | Your template with definitions required to communicate to the API. See also [Templating]({{< ref "_index.md#templating" >}})                                                                                                                                                                                                                                                                                                        |
| `cache_ttl`       | *[Duration]({{< ref "configuration_types.md#duration" >}})* | no        | yes         | Allows caching of the API responses. Defaults to 10 seconds. The cache key is calculated from the entire configuration of the hydrator instance and the available information about the current subject.                                                                                                                                                                                                                            |

**Example**

In this example the hydrator is configured to call an endpoint using the HTTP `GET` method with the subject id being part of the url path. As the endpoint requires the `X-My-Session-Cookie` cookie for subject authentication purposes, `forward_cookies` is used to achieve this.

```yaml
id: foo
type: generic
config:
  endpoint:
    url: https://some-other.service/users/{{.ID}}
    method: GET
  forward_cookies:
    - X-My-Session-Cookie
```