# Add an author

**Status:** Not built

## Who this is for

A librarian needs one catalog record per writer so books can name that writer consistently.

## Requirement

The library keeps an author and can find that same author later.

## Information

- Name (required)
- Biography (optional)
- Year of birth (optional)
- Nationality (optional)

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- The name is required. After trimming extra spaces, it is between 1 and 160 characters.
- Two authors cannot share the same name. Comparison ignores letter case and extra spaces at the ends.
- Biography, when present, is at most 4,000 characters.
- Year of birth, when present, is from 1 through the current calendar year.
- Nationality, when present, is at most 80 characters.
- A refused attempt leaves the catalog unchanged.

## Done when

- A librarian can add an author and find that author again with the name, biography, year of birth, and nationality that were saved.
- An administrator can do the same.
- A member, and anyone who is not signed in, is refused.
- A second author with the same name is refused, and the first record is unchanged.
- A missing name, a name that is too long, a biography that is too long, or a year of birth outside the allowed range is refused, and the reason names the problem.
