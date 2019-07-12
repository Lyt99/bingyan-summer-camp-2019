# HTTP notes

## Cookie

When the server is ready to send cookie, it will inform the client some info.

Set-Cookie field properties:

- **NAME=VALUE**
  - cookie's name and its value

- **expires=DATA**

  - default: only valid during **session**, normally before the browser is shut down.
  - once the cookie has been sent from the server to the client,  there is no way to delete it. However, it can be overwritten.

- **path=PATH**

  - limit the range that the cookie can be sent.

- **domain=DOMAIN NAME**

  - other domain can send cookies to the client. Eg: example.com is set, then www.example.com or www2.example.com can all send cookies to this client.
  - less secure if it is set.

- **Secure**

  - cookies can be sent only in **https** connection.

    ```http
    Set-Cookie: name=value; secure
    ```

- **HttpOnly**

  - mainly to prevent **Cross-site scripting, XSS** from stealing cookies.

    ```http
    Set-Cookie: name=value; HttpOnly
    ```

- **Cookie**

  - this field tells the server that when the client wants HTTP status support, it will contain a cookie in its request. This also support multiple cookies.