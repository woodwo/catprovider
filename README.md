# catprovider
This application integrates a gRPC call and can print a cat emoji to your console. It serves as an educational tool or 'lab' to demonstrate key concepts:

- gRPC Client Interaction: Demonstrates how to utilize client code provided by a remote party for effective communication in a distributed system.

- Usage of Mocks: Illustrates the role of mocks in testing. It shows how mocks can simplify the testing process by isolating the application logic from complex external dependencies.

This setup offers valuable learning opportunities by covering various aspects of Go and gRPC development within a controlled environment. It simplifies the understanding of each component, making it easier to grasp how they function together in a real-world application.

# howto

## prerequisites

Clone and run the server from the repository at https://github.com/woodwo/unknown/tree/main. This server will provide a Random function to work with.

## build and run

Run ```make run``` and enjoy all the cats.

# just run

Use compiled binary from `.build` folder : `.build/client`

## test

```make test```
