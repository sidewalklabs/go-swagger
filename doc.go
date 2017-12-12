// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*Package swagger (2.0) provides a powerful interface to your API

Contains an implementation of Swagger 2.0.
It knows how to serialize, deserialize and validate swagger specifications.

Swagger is a simple yet powerful representation of your RESTful API.
With the largest ecosystem of API tooling on the planet, thousands of developers are supporting Swagger
in almost every modern programming language and deployment environment.

With a Swagger-enabled API, you get interactive documentation, client SDK generation and discoverability.
We created Swagger to help fulfill the promise of APIs.

Swagger helps companies like Apigee, Getty Images, Intuit, LivingSocial, McKesson, Microsoft, Morningstar, and PayPal
build the best possible services with RESTful APIs.Now in version 2.0, Swagger is more enabling than ever.
And it's 100% open source software.

Install:

		go get -u github.com/sidewalklabs/go-swagger/cmd/swagger

The implementation also provides a number of command line tools to help working with swagger.

Currently there is a spec validator tool:

		swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json

To generate a server for a swagger spec document:

		swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]

To generate a client for a swagger spec document:

		swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]

To generate a swagger spec document for a go application:

		swagger generate spec -o ./swagger.json

There are several other sub commands available for the generate command

		Sub command | Description
		------------|----------------------------------------------------------------------------------
		operation   | generates one or more operations specified in the swagger definition
		model       | generates model files for one or more models specified in the swagger definition
		support     | generates the api builder and the main method
		server      | generates an entire server application
		client      | generates a client for a swagger specification
		spec        | generates a swagger spec document based on go code


Generating code

You're free to add files to the directories the generated code lands in, but the files generated by the generator itself
will be regenerated on following generation runs so any changes to those files will be lost.
However extra files you create won't be lost so they are safe to use for customizing the application to your needs.

To generate a server for a swagger spec document:

	swagger generate server -f ./swagger.json -A [application-name] [--principal [principal-name]]

*/
package swagger
