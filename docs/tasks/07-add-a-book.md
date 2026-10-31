# Add a book

**Status:** Not built

## Who this is for

A librarian needs to put a new title into the catalog so members can find it and borrow a copy.

## Requirement

A librarian or an administrator can add a book. The new book can be found later with the details that were saved.

## Information

- Title (required)
- Author (required). Once authors are part of the catalog, the book names an author the catalog already knows.
- ISBN (required)
- Description (optional)
- Publication year (required)
- Total copies (required)
- Available copies (required)

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- Title is required and is between 1 and 200 characters.
- Author is required. The displayed author name is between 1 and 160 characters.
- ISBN is required. It is an ISBN-10 or an ISBN-13, written without hyphens. A trailing check character may be the digit `X`.
- Two books cannot share the same ISBN. Hyphens and letter case are ignored when comparing.
- Description, when present, is at most 2,000 characters.
- Publication year is from 1450 through 2100.
- Total copies is a whole number from 1 through 10,000.
- Available copies is a whole number from 0 through 10,000, and is never greater than total copies.
- A refused attempt does not add a book.

## Done when

- A librarian can add a book and open it later with the same title, author, ISBN, description, year, and copy counts.
- An administrator can do the same.
- A member, and anyone who is not signed in, is refused.
- A second book with the same ISBN is refused, and the first book is unchanged.
- A title, author, ISBN, year, or copy count that breaks the rules above is refused, and the reason names the problem.
