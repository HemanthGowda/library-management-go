# Review activity

**Status:** Not built

## Who this is for

An administrator needs to see who changed the catalog, the loans, and the accounts, and when they did it.

## Requirement

The library records accountable activity, and an administrator can read that history. Passwords are never part of the history.

## Information

Each entry records:

- When it happened
- Who did it (name and email), or that the person was not signed in when the action allows that
- What happened
- Which author, book, loan, or account was affected

The history covers:

- Adding, updating, and removing an author
- Adding, updating, and removing a book
- Borrowing and returning a copy
- Changing a role
- Suspending and restoring an account

## Who may do this

- Administrator: yes
- Member and librarian: no
- Anyone who is not signed in: no

## Rules

- An entry is recorded when the action succeeds. A refused attempt is not recorded.
- Passwords, renewals, and anything else that could be used to sign in are never stored in the history.
- The list can be paged, at most 20 entries at a time, newest first.
- The reader can narrow the list to one kind of record (authors, books, loans, or accounts).
- Reading the history does not change it.

## Done when

- After a librarian adds a book, an administrator can see that librarian, the book, and the time.
- After a member borrows or returns a copy, an administrator can see that action in the history.
- After an administrator changes a role or suspends an account, that action appears in the history.
- No history entry contains a password.
- A member, a librarian, and anyone who is not signed in are refused.
