# Slow repeated sign-in failures

**Status:** Not built

## Who this is for

The library needs to slow someone who keeps guessing a password, without telling them whether an email address is on file.

## Requirement

After repeated wrong passwords for the same email, further sign-in attempts for that email are refused for a cooling-off period. The response is the same whether or not that email is registered.

## Information

The library remembers, for a short time, how many recent failures belong to an email and when the cooling-off period ends.

## Who may do this

This applies to every sign-in attempt, including people who are not yet signed in.

## Rules

- Five wrong passwords for the same email, counted within 15 minutes, start a cooling-off period of 15 minutes for that email.
- During the cooling-off period, sign-in for that email is refused even when the password is correct.
- A successful sign-in clears the failure count for that email.
- Failures older than 15 minutes no longer count.
- The message shown for a wrong password, for an unknown email, and for a cooling-off period is the same. It does not say whether the email is registered, and it does not say that a cooling-off period is the reason.
- A correct password for a different email is still accepted during another email's cooling-off period.

## Done when

- The fifth wrong password for one email, within 15 minutes, causes later attempts for that email to be refused for 15 minutes, including an attempt that uses the right password.
- After those 15 minutes, the right password is accepted again.
- A successful sign-in clears the count, so five new failures are required before the next cooling-off period.
- An attempt for an unknown email receives the same message as a wrong password for a known email.
- Signing in with a different email still succeeds while the first email is in its cooling-off period.
