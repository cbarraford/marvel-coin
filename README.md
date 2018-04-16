# Marvel Coin

Marvel Coin is a Bitcoin conceptual clone but with added security both in
securing your coins as well as making it much harder to send coins to the
wrong address/person.

## Features

Marvel Coin implements two-factor auth on the protocol layer to ensure that
even if someone steals your private key(s), they cannot steal your funds. A
brief [white paper](https://github.com/CBarraford/2FA-Coin) can be read
explaining this implementation details.

*Human friendly* addresses Marvel Coin also enables the ability to send coin(s) to an email address
rather than a public key. This makes it easier for humans to verify they are
sending to the right person. This feature is optional, for private
transactions, new public keys can be generated on each transaction.

*Confirm Codes* adds the ability for the receiver to get a one-time-use pin
which is required for the sender to enter to complete the transaction. This
ensures that you can confirm that the person you want is the one actually
getting the coin(s).

### In Development
Marvel is in development and not considered to be "production ready".
