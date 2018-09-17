# EasySSO
EasySSO is a simple, but nonetheless efficient go package to integrate a Single Sign-On in 
your application. EasySSO is compose of the following projects:
* [easy-sso-common](https://bitbucket.org/twuillemin/easy-sso-common): the common definition and structures that 
your currently browsing. This project Holds the common definition of the various projects.
* [easy-sso](https://bitbucket.org/twuillemin/easy-sso): the SSO server component that your currently browsing.
 Along with the server this project also include components for services (validating the query) and client (
 authenticating and connecting to the services). These components only rely on the Go default http.
* [easy-sso-mux](https://bitbucket.org/twuillemin/easy-sso-mux): a middleware for the [gorilla/mux](https://github.com/gorilla/mux)
router, validating client authentication.

# EasySSO Common
This package is a very simple just targeted at holding together the various common definitions for the sub-projects. The
main definitions are:
* errors: The various errors that could be returned by the components of EasySSO, such as `ErrUnauthorized`
* structures: The definitions of common structures or DTOs, such as:
  *  TokenRequestBody: The body of the request for getting a token
  *  TokenRefreshBody: The body of the request for refreshing a token
  *  AuthenticationResponse: The response of the server with the token information
  *  CustomClaims: The specific claims in the JWT token
  
This package is targeted at only retrieving the minimal dependencies. So for example the use of the [easy-sso-mux](https://bitbucket.org/twuillemin/easy-sso-mux)
package won't impose having all the dependencies from the server. So you will probably never have to include this
package directly by itself.
  
Boring... For real code, please check the [EasySSO](https://bitbucket.org/twuillemin/easy-sso) main project page

#Usage
Very hard: 

    import (
        "bitbucket.org/twuillemin/easy-sso-common/pkg/common"
    )
    
Woah!

# License

Copyright 2018 Thomas Wuillemin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.