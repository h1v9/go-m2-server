# Go M2 Server

Drop-in reimplementation of the Metin2 Server in Go.

This was made as a proof of concept. No real gameplay code is included.

Currently it only allows to join the game and walk around. No syncronization is implemented.

## Drop In

This project has been built to not require (almost) any modification to the game client.

Protocol has been adapted to TMP4 version. Packet sequence has to be disabled

### Encryption
Game protocol has been reimplemented thoroughly, original Diffie-Hellman is supported.
Normally the game chooses a random algorithm to encrypt the game packets. Currently only AES is supported.

## Building

1. Install go
2. Run `go mod install`
3. Run `build.bat build-all` (Windows) / Run `go build` on other OSes

## Running

Just run the executable once compiled. Auth will be listening on port `23070`, Game on port `23000`

## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](LICENSE) file for details.
