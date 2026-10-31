# Reset a forgotten password

**Status:** Not built

## Who this is for

A person who has forgotten their password needs to choose a new one without already being signed in.

## Requirement

A person can ask for a password reset for an email address and then set a new password that meets the current password rules. The reset can be used once and then expires. The response to the request is the same whether or not that email is registered.

## Information

A reset remembers which account it belongs to, when it expires, and whether it has already been used. The new password replaces the old one. The old password stops working.

## Who may do this

- Anyone who can reach the library, including someone who is not signed in: may request a reset and may complete one they were given
- A reset only changes the account for the email it was issued to

## Rules

- Asking for a reset always ends with the same confirmation. When the email is registered, the person is given a way to finish the reset. When the email is not registered, nothing is sent and no account changes.
- The way to finish the reset works once and expires 30 minutes after it was issued.
- Using an expired or already-used reset is refused, and the password stays as it was.
- The new password is required, is between 8 and 72 characters, and contains at least one letter and one number.
- A new password that breaks those rules is refused, and the reset can still be used until it expires or succeeds.
- After a successful reset, the person can sign in with the new password, and the previous password is refused.
- Any earlier renewal for that account stops working when the password changes.

## Done when

- A person with a registered email can set a new valid password and then sign in with it.
- The previous password is refused after the reset.
- The same reset cannot be used twice.
- An expired reset is refused, and the password is unchanged.
- Asking for a reset of an unknown email returns the same confirmation as a known email, and no account is created.
- A password that is too short, or that has no letter or no number, is refused.
