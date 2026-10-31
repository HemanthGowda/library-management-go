# Update an author

**Status:** Not built

## Who this is for

A librarian needs to correct an author's name or details after the record was first saved.

## Requirement

A librarian or an administrator can replace the saved details of one author. The updated author can still be found later.

## Information

The same information as adding an author: name (required), biography, year of birth, and nationality.

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- The author must already exist. An unknown author is refused.
- The name stays required, 1 to 160 characters after trimming, and unique regardless of letter case and extra spaces.
- The author may keep their own current name.
- Biography, year of birth, and nationality follow the same limits as when an author is added.
- A refused attempt leaves that author unchanged.

## Done when

- A librarian can change an author's name and optional details, and a later view shows the new values.
- An administrator can do the same.
- Renaming an author to a name another author already has is refused, and both records stay as they were.
- A member, and anyone who is not signed in, is refused.
- An unknown author is refused.
