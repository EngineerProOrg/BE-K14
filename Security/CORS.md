Browser security prevents a web page from making requests to a different domain than the one that served the web page. This restriction is called the same-origin policy. The same-origin policy prevents a malicious site from reading sensitive data from another site. Sometimes, you might want to allow other sites to make cross-origin requests to your app

- Cross Origin Resource Sharing (CORS):
    - Is a W3C standard that allows a server to relax the same-origin policy.
    - Is not a security feature, CORS relaxes security. An API is not safer by allowing CORS.
    - Allows a server to explicitly allow some cross-origin requests while rejecting others.
    - Is safer and more flexible than earlier techniques

### Same origin
Two URLs have the same origin if they have identical schemes, hosts, and ports.
These two URLs have the same origin:
- https://example.com/foo.html
- https://example.com/bar.html
These URLs have different origins than the previous two URLs:
- https://example.net: Different domain
- https://contoso.example.com/foo.html: Different subdomain
- http://example.com/foo.html: Different scheme
- https://example.com:9000/foo.html: Different port

### Enable CORS
There are three ways to enable CORS:
In middleware using a named policy or default policy.
- Using endpoint routing.
- With the [EnableCors] attribute.
- Using the [EnableCors] attribute with a named policy provides the finest control in limiting endpoints that support CORS.

#### CORS with named policy and middleware
CORS Middleware handles cross-origin requests. The following code applies a CORS policy to all the app's endpoints with the specified origins:

```c#
var  MyAllowSpecificOrigins = "_myAllowSpecificOrigins";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy(name: MyAllowSpecificOrigins,
                      policy  =>
                      {
                          policy.WithOrigins("http://example.com",
                                              "http://www.contoso.com");
                      });
});

// services.AddResponseCaching();

builder.Services.AddControllers();

var app = builder.Build();
app.UseHttpsRedirection();
app.UseStaticFiles();
app.UseRouting();

app.UseCors(MyAllowSpecificOrigins);

app.UseAuthorization();

app.MapControllers();

app.Run();
```

The preceding code:
- Sets the policy name to _myAllowSpecificOrigins. The policy name is arbitrary.
- Calls the UseCors extension method and specifies the _myAllowSpecificOrigins CORS policy. UseCors adds the CORS middleware. The call to UseCors must be placed after UseRouting, but before UseAuthorization. For more information, see Middleware order.
- Calls AddCors with a lambda expression. The lambda takes a CorsPolicyBuilder object. Configuration options, such as WithOrigins, are described later in this article.
- Enables the _myAllowSpecificOrigins CORS policy for all controller endpoints. See endpoint routing to apply a CORS policy to specific endpoints.
- When using Response Caching Middleware, call UseCors before UseResponseCaching.
With endpoint routing, the CORS middleware must be configured to execute between the calls to UseRouting and UseEndpoints.
![plot](../img/middleware-pipeline.svg)

#### Enable Cors with endpoint routing
With endpoint routing, CORS can be enabled on a per-endpoint basis using the RequireCors set of extension methods:
```c#
var MyAllowSpecificOrigins = "_myAllowSpecificOrigins";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy(name: MyAllowSpecificOrigins,
                      policy =>
                      {
                          policy.WithOrigins("http://example.com",
                                              "http://www.contoso.com");
                      });
});

builder.Services.AddControllers();
builder.Services.AddRazorPages();

var app = builder.Build();

app.UseHttpsRedirection();
app.UseStaticFiles();
app.UseRouting();

app.UseCors();

app.UseAuthorization();

app.UseEndpoints(endpoints =>
{
    endpoints.MapGet("/echo",
        context => context.Response.WriteAsync("echo"))
        .RequireCors(MyAllowSpecificOrigins);

    endpoints.MapControllers()
             .RequireCors(MyAllowSpecificOrigins);

    endpoints.MapGet("/echo2",
        context => context.Response.WriteAsync("echo2"));

    endpoints.MapRazorPages();
});

app.Run();
```

In the preceding code:
- app.UseCors enables the CORS middleware. Because a default policy hasn't been configured, app.UseCors() alone doesn't enable CORS.
- The /echo and controller endpoints allow cross-origin requests using the specified policy.
- The /echo2 and Razor Pages endpoints do not allow cross-origin requests because no default policy was specified.
The [DisableCors] attribute does not disable CORS that has been enabled by endpoint routing with RequireCors.

#### Enable CORS with attributes
Enabling CORS with the [EnableCors] attribute and applying a named policy to only those endpoints that require CORS provides the finest control.
The [EnableCors] attribute provides an alternative to applying CORS globally. The [EnableCors] attribute enables CORS for selected endpoints, rather than all endpoints:
- [EnableCors] specifies the default policy.
- [EnableCors("{Policy String}")] specifies a named policy.
- The [EnableCors] attribute can be applied to:

Razor Page PageModel
- Controller
- Controller action method
- Different policies can be applied to controllers, page models, or action methods with the [EnableCors] attribute. When the [EnableCors] attribute is applied to a controller, page model, or action method, and CORS is enabled in middleware, both policies are applied. We recommend against combining policies. Use the [EnableCors] attribute or middleware, not both in the same app.

The following code applies a different policy to each method:
```c#
[Route("api/[controller]")]
[ApiController]
public class WidgetController : ControllerBase
{
    // GET api/values
    [EnableCors("AnotherPolicy")]
    [HttpGet]
    public ActionResult<IEnumerable<string>> Get()
    {
        return new string[] { "green widget", "red widget" };
    }

    // GET api/values/5
    [EnableCors("Policy1")]
    [HttpGet("{id}")]
    public ActionResult<string> Get(int id)
    {
        return id switch
        {
            1 => "green widget",
            2 => "red widget",
            _ => NotFound(),
        };
    }
}
```
The following code creates two CORS policies:
```c#
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy("Policy1",
        policy =>
        {
            policy.WithOrigins("http://example.com",
                                "http://www.contoso.com");
        });

    options.AddPolicy("AnotherPolicy",
        policy =>
        {
            policy.WithOrigins("http://www.contoso.com")
                                .AllowAnyHeader()
                                .AllowAnyMethod();
        });
});

builder.Services.AddControllers();

var app = builder.Build();

app.UseHttpsRedirection();

app.UseRouting();

app.UseCors();

app.UseAuthorization();

app.MapControllers();

app.Run();
```

### CORS policy options
This section describes the various options that can be set in a CORS policy:
- Set the allowed origins
- Set the allowed HTTP methods
- Set the allowed request headers
- Set the exposed response headers
- Credentials in cross-origin requests
- Set the preflight expiration time

#### Set the allowed origins
AllowAnyOrigin: Allows CORS requests from all origins with any scheme (http or https). AllowAnyOrigin is insecure because any website can make cross-origin requests to the app.
AllowAnyOrigin affects preflight requests and the Access-Control-Allow-Origin header.
SetIsOriginAllowedToAllowWildcardSubdomains: Sets the IsOriginAllowed property of the policy to be a function that allows origins to match a configured wildcard domain when evaluating if the origin is allowed.
```c#
var MyAllowSpecificOrigins = "_MyAllowSubdomainPolicy";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy(name: MyAllowSpecificOrigins,
        policy =>
        {
            policy.WithOrigins("https://*.example.com")
                .SetIsOriginAllowedToAllowWildcardSubdomains();
        });
});

builder.Services.AddControllers();

var app = builder.Build();
```

#### Set the allowed HTTP methods
AllowAnyMethod:
- Allows any HTTP method:
- Affects preflight requests and the Access-Control-Allow-Methods header.

#### Set the allowed request headers
To allow specific headers to be sent in a CORS request, called author request headers, call WithHeaders and specify the allowed headers:
```c#
using Microsoft.Net.Http.Headers;

var MyAllowSpecificOrigins = "_MyAllowSubdomainPolicy";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy(name: MyAllowSpecificOrigins,
       policy =>
       {
           policy.WithOrigins("http://example.com")
                  .WithHeaders(HeaderNames.ContentType, "x-custom-header");
       });
});

builder.Services.AddControllers();

var app = builder.Build();
```
AllowAnyHeader affects preflight requests and the Access-Control-Request-Headers header. 

A CORS Middleware policy match to specific headers specified by WithHeaders is only possible when the headers sent in Access-Control-Request-Headers exactly match the headers stated in WithHeaders.

For instance, consider an app configured as follows:

```c#
app.UseCors(policy => policy.WithHeaders(HeaderNames.CacheControl));
```

CORS Middleware declines a preflight request with the following request header because Content-Language (HeaderNames.ContentLanguage) isn't listed in WithHeaders:
```
Access-Control-Request-Headers: Cache-Control, Content-Language
```
The app returns a 200 OK response but doesn't send the CORS headers back. Therefore, the browser doesn't attempt the cross-origin request.

#### Set the exposed response headers
By default, the browser doesn't expose all of the response headers to the app. For more information, see W3C Cross-Origin Resource Sharing (Terminology): Simple Response Header.
The response headers that are available by default are:
- Cache-Control
- Content-Language
- Content-Type
- Expires
- Last-Modified
- Pragma
The CORS specification calls these headers simple response headers. To make other headers available to the app, call WithExposedHeaders:
```c#
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy("MyExposeResponseHeadersPolicy",
        policy =>
        {
            policy.WithOrigins("https://*.example.com")
                   .WithExposedHeaders("x-custom-header");
        });
});

builder.Services.AddControllers();

var app = builder.Build();
```

#### Credentials in cross-origin requests
Credentials require special handling in a CORS request. By default, the browser doesn't send credentials with a cross-origin request. Credentials include cookies and HTTP authentication schemes. To send credentials with a cross-origin request, the client must set XMLHttpRequest.withCredentials to true.

The server must allow the credentials. To allow cross-origin credentials, call AllowCredentials:
```c#
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy("MyMyAllowCredentialsPolicy",
        policy =>
        {
            policy.WithOrigins("http://example.com")
                   .AllowCredentials();
        });
});

builder.Services.AddControllers();

var app = builder.Build();
```
The HTTP response includes an Access-Control-Allow-Credentials header, which tells the browser that the server allows credentials for a cross-origin request.

If the browser sends credentials but the response doesn't include a valid Access-Control-Allow-Credentials header, the browser doesn't expose the response to the app, and the cross-origin request fails.

Allowing cross-origin credentials is a security risk. A website at another domain can send a signed-in user's credentials to the app on the user's behalf without the user's knowledge.

The CORS specification also states that setting origins to "*" (all origins) is invalid if the Access-Control-Allow-Credentials header is present.
##### Why Is It a Security Risk?
The risk is not in credentials itself, but in misconfiguration:
- Overly Permissive Origins: If the server allows Access-Control-Allow-Origin: * with Allow-Credentials: true, it's a security violation (and browsers block it).
- Cross-Site Request Forgery (CSRF): If a server trusts cookies from any origin without proper CSRF protection, attackers can forge authenticated requests from malicious sites.

#### Preflight request
For some CORS requests, the browser sends an additional OPTIONS request before making the actual request. This request is called a preflight request. The browser can skip the preflight request if all the following conditions are true:
- The request method is GET, HEAD, or POST.
- The app doesn't set request headers other than Accept, Accept-Language, Content-Language, Content-Type, or Last-Event-ID.
The Content-Type header, if set, has one of the following values:
- application/x-www-form-urlencoded
- multipart/form-data
- text/plain
The rule on request headers set for the client request applies to headers that the app sets by calling setRequestHeader on the XMLHttpRequest object. The CORS specification calls these headers author request headers. The rule doesn't apply to headers the browser can set, such as User-Agent, Host, or Content-Length.

The following is an example response similar to the preflight request made from the [Put test] button in the Test CORS section of this document.
```
General:
Request URL: https://cors3.azurewebsites.net/api/values/5
Request Method: OPTIONS
Status Code: 204 No Content

Response Headers:
Access-Control-Allow-Methods: PUT,DELETE,GET
Access-Control-Allow-Origin: https://cors1.azurewebsites.net
Server: Microsoft-IIS/10.0
Set-Cookie: ARRAffinity=8f8...8;Path=/;HttpOnly;Domain=cors1.azurewebsites.net
Vary: Origin

Request Headers:
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Access-Control-Request-Method: PUT
Connection: keep-alive
Host: cors3.azurewebsites.net
Origin: https://cors1.azurewebsites.net
Referer: https://cors1.azurewebsites.net/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0
```

The preflight request uses the HTTP OPTIONS method. It may include the following headers:
- Access-Control-Request-Method: The HTTP method that will be used for the actual request.
- Access-Control-Request-Headers: A list of request headers that the app sets on the actual request. As stated earlier, this doesn't include headers that the browser sets, such as User-Agent.
If the preflight request is denied, the app returns a 200 OK response but doesn't set the CORS headers. Therefore, the browser doesn't attempt the cross-origin request

##### Set the preflight expiration time
The Access-Control-Max-Age header specifies how long the response to the preflight request can be cached. To set this header, call SetPreflightMaxAge:
```c#
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddCors(options =>
{
    options.AddPolicy("MySetPreflightExpirationPolicy",
        policy =>
        {
            policy.WithOrigins("http://example.com")
                   .SetPreflightMaxAge(TimeSpan.FromSeconds(2520));
        });
});

builder.Services.AddControllers();

var app = builder.Build();
```

### Enable CORS on an endpoint
#### How CORS works
The CORS specification introduced several new HTTP headers that enable cross-origin requests. If a browser supports CORS, it sets these headers automatically for cross-origin requests. Custom JavaScript code isn't required to enable CORS.

The following is an example of a cross-origin request from the Values test button to https://cors1.azurewebsites.net/api/values. The Origin header:
- Provides the domain of the site that's making the request.
- Is required and must be different from the host.
General headers
```
Request URL: https://cors1.azurewebsites.net/api/values
Request Method: GET
Status Code: 200 OK
```

Response headers
```
Content-Encoding: gzip
Content-Type: text/plain; charset=utf-8
Server: Microsoft-IIS/10.0
Set-Cookie: ARRAffinity=8f...;Path=/;HttpOnly;Domain=cors1.azurewebsites.net
Transfer-Encoding: chunked
Vary: Accept-Encoding
X-Powered-By: ASP.NET
```

Request header
```
Accept: */*
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Host: cors1.azurewebsites.net
Origin: https://cors3.azurewebsites.net
Referer: https://cors3.azurewebsites.net/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
User-Agent: Mozilla/5.0 ...
```

In the preceding Response headers, the server sets the Access-Control-Allow-Origin header in the response. The https://cors1.azurewebsites.net value of this header matches the Origin header from the request.

If AllowAnyOrigin is called, the Access-Control-Allow-Origin: *, the wildcard value, is returned. AllowAnyOrigin allows any origin.

If the response doesn't include the Access-Control-Allow-Origin header, the cross-origin request fails. Specifically, the browser disallows the request. Even if the server returns a successful response, the browser doesn't make the response available to the client app.