# Confirm an email address

**Status:** Not built

## Who this is for

The library needs a new member to show they control the email address they registered before that account can sign in.

## Requirement

A new account cannot sign in until the person confirms control of the email. Accounts that already exist when this capability is introduced can still sign in.

## Information

A new account remembers whether its email has been confirmed. A confirmation remembers which account it belongs to, when it expires, and whether it has already been used.

## Who may do this

- A person who just registered: must confirm before the first sign-in
- Anyone holding a valid confirmation for that account: may complete it
- Accounts already on file before this capability: may sign in without a new confirmation

## Rules

- Registration of a new account still creates the account, and it still starts as a member. Sign-in is refused until the email is confirmed.
- The person is given a way to confirm. It works once and expires 24 hours after it was issued.
- Confirming with a valid, unused confirmation marks the email confirmed. The person can then sign in.
- An expired or already-used confirmation is refused, and the account stays unconfirmed.
- The person can ask for a fresh confirmation while the account is still unconfirmed. The new one expires 24 hours after it is issued, and the previous one stops working.
- Asking again for an address that is already confirmed does not change the account. The response does not say whether the email is registered.
- Accounts that could already sign in before this capability keep that ability.

## Done when

- A newly registered person is refused at sign-in until they confirm the email, including when the password is correct.
- After a valid confirmation, that person can sign in.
- The same confirmation cannot be used twice.
- An expired confirmation is refused, and a freshly requested confirmation then succeeds.
- A person whose account already existed before this capability can still sign in.
