# Untrustroot

Sparked by the WoSign fiasco the idea was born to make a small utility that
would allow me to inspect and analyze the macOS trusted certificate
authorities.

In its current form it can:

  * Perform rudimentary analysis of the keychain

Stuff I'd like to add:

  * Display information about each CA, potentially group them together
    if an entity has multiple of them.
  * Mark a (number of) certificat(s) as untrusted.
  * Whitelist to avoid untrusting certs that are required for the proper
    functioning of macOS itself.

Running this on my laptop with macOS Sierra (10.12) I get:

```
There are 166 CA's in your keychain:
  • 27 have CRL distribution points specified
  • 2 have OCSP responsders specified
Using the following signing algorithms:
  • SHA1-RSA: 95
  • SHA256-RSA: 52
  • ECDSA-SHA384: 12
  • SHA384-RSA: 5
  • SHA512-RSA: 1
  • ECDSA-SHA256: 1
Issued by entities in the following countries:
  • US: 62
  • CH: 12
  • BM: 6
  • IL: 6
  • ES: 5
  • DE: 5
  • JP: 5
  • TW: 4
  • TR: 4
  • EU: 4
  • FR: 4
  • PL: 4
  • NL: 3
  • GB: 3
  • NO: 3
  • FI: 3
  • SK: 2
  • HU: 2
  • SE: 2
  • EE: 2
  • BE: 2
  • CN: 2
  • GR: 1
  • CZ: 1
  • CA: 1
  • DK: 1
  • IT: 1
  • HK: 1
  • RO: 1
  • VE: 1
  • IE: 1
  • KR: 1
  ```

  I'm not quite sure how the EU ended up being a country but oh well.
