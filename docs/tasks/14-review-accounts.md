# Review accounts

**Status:** Not built

## Who this is for

An administrator needs to see who can use the library and what each person is allowed to do.

## Requirement

An administrator can page through accounts. The list never includes a password or any other secret.

## Information

Each account shows:

- Full name
- Email
- Roles held (member, librarian, administrator)
- Whether the account can currently sign in

## Who may do this

- Administrator: yes
- Member and librarian: no
- Anyone who is not signed in: no

## Rules

- Results come back one page at a time, at most 20 accounts, in name order.
- The reader can narrow the list by name or email. The match ignores letter case.
- When nothing matches, the result is an empty page.
- Passwords, and anything that could be used to sign in as that person, are never included.
- Reading the list does not change any account.

## Done when

- An administrator can page through accounts and see each person's name, email, roles, and whether they can sign in.
- A name or email filter returns only the matching accounts.
- The result contains no password.
- A member, a librarian, and anyone who is not signed in are refused.
