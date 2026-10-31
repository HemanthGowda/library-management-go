# Suspend or restore an account

**Status:** Not built

## Who this is for

An administrator needs to stop a person from signing in, and to allow them back later, without deleting the account.

## Requirement

An administrator can suspend an account and later restore it. A suspended account cannot sign in. A session that was already issued keeps working until it expires. Restoring the account allows sign-in again.

## Information

The account remembers whether it is allowed to sign in.

## Who may do this

- Administrator: yes
- Member and librarian: no
- Anyone who is not signed in: no

## Rules

- The account must already exist.
- Suspending an account that is already suspended leaves it suspended.
- Restoring an account that can already sign in leaves it able to sign in.
- While an account is suspended, a sign-in with the correct password is refused.
- A session issued before the suspension keeps working until that session expires. It cannot be renewed after the suspension.
- The last remaining administrator cannot be suspended. That attempt is refused, and that account can still sign in.
- Suspension does not delete the person, their loans, or their role.

## Done when

- An administrator can suspend an account, and a later sign-in with the correct password is refused.
- A session that was already issued before the suspension still works until it expires.
- An administrator can restore that account, and the person can sign in again.
- Suspending the last administrator is refused.
- A member, a librarian, and anyone who is not signed in are refused.
