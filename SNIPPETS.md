# Code Snippets

For your convenience we've added code snippets for importing media in various programming languages including `C`, `C#`, `Go`, `Java`, `jQuery`, `Node.js`, `Objective-C`, `OCaml`, `PHP`, `Python`, `Ruby` and `Swift`.

### cURL

```shell
curl -X POST -H "Gifs-API-Key: gifs56d63999f0f34" -H "Content-Type: application/json" -d '{
 "source": "https://vine.co/v/ibAU6OH2I0K",
 "title": "2015 Craziness",
 "tags": ["crazy", "hand drawn", "2015", "art"],
 "attribution": {
   "site": "vine",
   "user": "AliciaHerber"
 }
}' "https://api.gifs.com/media/import"
```

### C (LibCurl)

```c
CURL *hnd = curl_easy_init();

curl_easy_setopt(hnd, CURLOPT_CUSTOMREQUEST, "POST");
curl_easy_setopt(hnd, CURLOPT_URL, "https://api.gifs.com/media/import");

struct curl_slist *headers = NULL;
headers = curl_slist_append(headers, "Gifs-API-Key: gifs56d63999f0f34");
headers = curl_slist_append(headers, "Content-Type: application/json");
curl_easy_setopt(hnd, CURLOPT_HTTPHEADER, headers);

curl_easy_setopt(hnd, CURLOPT_POSTFIELDS, "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}");

CURLcode ret = curl_easy_perform(hnd);
```

### C# (RestSharp)

```csharp
var client = new RestClient("https://api.gifs.com/media/import");
var request = new RestRequest(Method.POST);
request.AddHeader("Content-Type", "Gifs-API-Key: gifs56d63999f0f34");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

### Go

```go
url := "https://api.gifs.com/media/import"

payload := strings.NewReader(`{
 "source": "https://vine.co/v/ibAU6OH2I0K",
 "title": "2015 Craziness",
 "tags": ["crazy", "hand drawn", "2015", "art"],
 "attribution": {
   "site": "vine.co",
   "user": "AliciaHerber"
 }
}`)

req, _ := http.NewRequest("POST", url, payload)
req.Header.Add("Gifs-API-Key", "gifs56d63999f0f34")
req.Header.Add("Content-Type", "application/json")

res, _ := http.DefaultClient.Do(req)

defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)

fmt.Println(res)
fmt.Println(string(body))
```

### Java (Unirest)

```java
HttpResponse<String> response = Unirest.post("https://api.gifs.com/media/import")
  .header("Gifs-API-Key", "gifs56d63999f0f34")
  .header("Content-Type", "application/json")
  .body("{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}")
  .asString();
```

### JavaScript (jQuery)

```javascript
var settings = {
  "async": true,
  "crossDomain": true,
  "url": "https://api.gifs.com/media/import",
  "method": "POST",
  "headers": {
    "Gifs-API-Key": "gifs56d63999f0f34",
    "Content-Type": "application/json"
  },
  "processData": false,
  "data": "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}"
}

$.ajax(settings).done(function (response) {
  console.log(response);
});
```

### Javascript (Nodejs)

```js
var http = require("https");

var options = {
  "method": "POST",
  "hostname": "api.gifs.com",
  "port": null,
  "path": "/media/import",
  "headers": {
    "Gifs-API-Key": "gifs56d63999f0f34",
    "Content-Type": "application/json"
  }
};

var req = http.request(options, function (res) {
  var chunks = [];

  res.on("data", function (chunk) {
    chunks.push(chunk);
  });

  res.on("end", function () {
    var body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});

req.write(JSON.stringify({ source: 'https://vine.co/v/ibAU6OH2I0K',
  title: '2015 Craziness',
  tags: [ 'cat', 'hand drawn', 'cat flip', 'art' ],
  attribution: { site: 'vine.co', user: 'AliciaHerber' } }));
req.end();
```

### Objective-C

```obj-c
#import <Foundation/Foundation.h>

NSDictionary *headers = @{ @"Gifs-API-Key": @"gifs56d63999f0f34",
                        @"Content-Type": @"application/json" };

NSDictionary *parameters = @{ @"source": @"https://vine.co/v/ibAU6OH2I0K",
                              @"title": @"2015 Craziness",
                              @"tags": @[ @"crazy", @"hand drawn", @"2015", @"art" ],
                              @"attribution": @{ @"site": @"vine.co", @"user": @"AliciaHerber" } };

NSData *postData = [NSJSONSerialization dataWithJSONObject:parameters options:0 error:nil];

NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:[NSURL URLWithString:@"https://api.gifs.com/media/import"]
                                                       cachePolicy:NSURLRequestUseProtocolCachePolicy
                                                   timeoutInterval:10.0];
[request setHTTPMethod:@"POST"];
[request setAllHTTPHeaderFields:headers];
[request setHTTPBody:postData];

NSURLSession *session = [NSURLSession sharedSession];
NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request
                                            completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                                                if (error) {
                                                    NSLog(@"%@", error);
                                                } else {
                                                    NSHTTPURLResponse *httpResponse = (NSHTTPURLResponse *) response;
                                                    NSLog(@"%@", httpResponse);
                                                }
                                            }];
[dataTask resume];
```

### OCaml

```ocaml
open Cohttp_lwt_unix
open Cohttp
open Lwt

let uri = Uri.of_string "https://api.gifs.com/media/import" in
let headers = Header.init ()
  |> fun h -> Header.add h "Gifs-API-Key" "gifs56d63999f0f34"
  |> fun h -> Header.add h "Content-Type" "application/json"
in
let body = Cohttp_lwt_body.of_string "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}" in

Client.call ~headers ~body `POST uri
>>= fun (res, body_stream) ->
  (* Do stuff with the result *)
```

### PHP

```php
$request = new HttpRequest();
$request->setUrl('https://api.gifs.com/media/import');
$request->setMethod(HTTP_METH_POST);

$request->setHeaders(array(
  'Gifs-API-Key' => 'gifs56d63999f0f34'
  'Content-Type' => 'application/json'
));

$request->setBody('{
 "source": "https://vine.co/v/ibAU6OH2I0K",
 "title": "2015 Craziness",
 "tags": ["crazy", "hand drawn", "2015", "art"],
 "attribution": {
   "site": "vine.co",
   "user": "AliciaHerber"
 }
}');

try {
  $response = $request->send();

  echo $response->getBody();
} catch (HttpException $ex) {
  echo $ex;
}
```

### Python

```python
import http.client

conn = http.client.HTTPSConnection("api.gifs.com")

payload = "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}"

headers = { 'Gifs-API-Key': "gifs56d63999f0f34", 'Content-Type': "application/json" }

conn.request("POST", "/media/import", payload, headers)

res = conn.getresponse()
data = res.read()

print(data.decode("utf-8"))
```

### Ruby

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.gifs.com/media/import")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

request = Net::HTTP::Post.new(url)
request["Gifs-API-Key"] = 'gifs56d63999f0f34'
request["Content-Type"] = 'application/json'
request.body = "{\n \"source\": \"https://vine.co/v/ibAU6OH2I0K\",\n \"title\": \"2015 Craziness\",\n \"tags\": [\"crazy\", \"hand drawn\", \"2015\", \"art\"],\n \"attribution\": {\n   \"site\": \"vine.co\",\n   \"user\": \"AliciaHerber\"\n }\n}"

response = http.request(request)
puts response.read_body
```

### Swift

```swift
import Foundation

let parameters = [
  "source": "https://vine.co/v/ibAU6OH2I0K",
  "title": "2015 Craziness",
  "tags": ["crazy", "hand drawn", "2015", "art"],
  "attribution": [
    "site": "vine.co",
    "user": "AliciaHerber"
  ]
]

let postData = NSJSONSerialization.dataWithJSONObject(parameters, options: nil, error: nil)

var request = NSMutableURLRequest(URL: NSURL(string: "https://api.gifs.com/media/import")!,
                                        cachePolicy: .UseProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.HTTPMethod = "POST"
request.allHTTPHeaderFields = [ "Gifs-API-Key": "gifs56d63999f0f34", "Content-Type": "application/json" ]
request.HTTPBody = postData

let session = NSURLSession.sharedSession()
let dataTask = session.dataTaskWithRequest(request, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    println(error)
  } else {
    let httpResponse = response as? NSHTTPURLResponse
    println(httpResponse)
  }
})

dataTask.resume()
```
