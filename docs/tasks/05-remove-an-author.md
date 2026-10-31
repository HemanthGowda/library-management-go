# Remove an author

**Status:** Not built

## Who this is for

A librarian needs to take an author out of the catalog when that writer should no longer be offered for new books.

## Requirement

A librarian or an administrator can remove an author who has no books left in the catalog. After removal, that author cannot be found.

## Information

The library needs to know which author is being removed.

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- The author must already exist. An unknown author is refused.
- Removal is refused while any book still names that author. Those books stay in the catalog.
- After a successful removal, a search and a direct view no longer find that author.
- A refused attempt leaves the author and their books unchanged.

## Done when

- A librarian can remove an author who has no books, and that author can no longer be found.
- An administrator can do the same.
- Removal is refused while a book still names that author, and the author and the book both remain.
- A member, and anyone who is not signed in, is refused.
- An unknown author is refused.
