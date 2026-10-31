# Tie books to authors

**Status:** Not built

## Who this is for

Readers and librarians need a book to name a real author from the catalog, so a rename or a correction shows up everywhere that author is used.

## Requirement

A book names an author the catalog already knows. When that author's name changes, every book that names them shows the new name. Books already in the catalog remain readable.

## Information

Each book remembers which catalog author it names, and shows that author's current name.

## Who may do this

This changes the rules for adding a book and for updating a book.

- Librarian and administrator: may add or update a book so that it names a known author
- Member: no
- Anyone who is not signed in: no

## Rules

- Adding a book, or changing a book, requires an author that can already be found in the catalog.
- An unknown author is refused, and the book is left unchanged. A new book is not created.
- When an author's name is corrected, every book that names that author shows the corrected name.
- Books already on file before this capability exists stay available. Each of them still shows an author name a reader can recognize.
- Removing an author is still refused while any book names that author.

## Done when

- A librarian can save a book only when it names an author already on file.
- Saving a book with an unknown author is refused.
- After an author's name is corrected, a reader sees the new name on every book that names that author.
- Books that were already in the catalog can still be opened, and each still shows an author name.
