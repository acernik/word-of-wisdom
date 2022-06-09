## Word Of Wisdom
This repository contains implementation of a simple Word Of Wisdom TCP server. The server is protected from 
DDOS attacks by a Proof Of Work Hashcash algorithm.

## Hashcash
The reason why Hashcash has been chosen is that it is a strong algorithm which is used by the projects like `Bitcoin`.

### Advantages of Hashcash
It is quite simple to implement in an application and no central server is needed. It provides good protection 
against DDOS attacks and email spam.

### Disadvantages of Hashcash
It could potentially require quite a lot of computational resources. And another thing is that as the time goes by and 
computers are getting more and more powerful more and more complex calculations need to be requested to be done by the 
clients in order to prove the work. However, not everyone has access to the lates and the most powerful computers, which 
could mean that for some users Hashcash will not work very well.

For the purpose of this task the Hashcash algorithm was not implemented from scratch. The existing `Go` implementation 
was used --> `https://github.com/umahmood/hashcash`.

## Unit tests
To run unit tests execute `make tests` command from the root directory. This will run all unit tests with coverage.

## Usage
To run the app first run the `make server` command. This will start the server which is now accepting TCP requests.
To get a Word Of Wisdom quote from the server run `make client` command. This will start up a client app and will 
make a request to the server which will ask to perform the PoW and once the PoW is verified the server will return 
a quote from Word Of Wisdom.

## Docker
Both `server` and `client` applications have `Dockerfile` each. `client` is `Dockerfile.client` and 
`server` is `Dockerfile.server`.

To run those files you need to run the following commands:
- `make launch-srv` to run `server` 
- `make launch-cli` to run `client`