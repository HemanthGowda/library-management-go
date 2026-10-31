# Renew a session

**Status:** Not built

## Who this is for

A person who is already signed in needs to keep working after a short period without typing their password again.

## Requirement

Signing in gives a short period of access and a separate way to renew that access. The person can renew without typing the password again. Renewal stops when the account is suspended.

## Information

A successful sign-in gives:

- A short-lived proof of access that expires
- A longer-lived renewal that can be exchanged once for a new proof of access and a new renewal

The renewal remembers which account it belongs to and when it expires.

## Who may do this

- Any person with an account that can sign in: yes
- A suspended account: no
- Anyone who is not holding a valid renewal: no

## Rules

- The proof of access from sign-in expires after 15 minutes.
- The renewal expires after 14 days.
- Presenting a valid renewal, before it expires, gives a new 15-minute proof of access and a new 14-day renewal. The previous renewal cannot be used again.
- An expired renewal is refused. The person must sign in with their password.
- A renewal for a suspended account is refused, even if the renewal has not yet reached its expiry.
- A renewal never includes the password.

## Done when

- After signing in, a person can renew and receive a new period of access without typing the password.
- The same renewal cannot be used a second time.
- An expired renewal is refused.
- Renewal is refused after the account is suspended.
- A person whose proof of access has expired, and who still holds a valid renewal, can obtain a new proof of access.
