# Remove a book

**Status:** Not built

## Who this is for

A librarian needs to take a title out of the catalog when the library no longer offers it.

## Requirement

A librarian or an administrator can remove a book that nobody currently has on loan. After removal, that book cannot be found.

## Information

The library needs to know which book is being removed.

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- The book must already exist. An unknown book is refused.
- Removal is refused while any copy of that book is still on loan. The book and those loans stay as they are.
- After a successful removal, a search and a direct view no longer find that book.
- A refused attempt leaves the book unchanged.

## Done when

- A librarian can remove a book that has no active loan, and that book can no longer be found.
- An administrator can do the same.
- Removal is refused while a copy is still on loan, and the book remains.
- A member, and anyone who is not signed in, is refused.
- An unknown book is refused.
