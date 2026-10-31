# Search authors

**Status:** Not built

## Who this is for

A signed-in person needs to find authors without already knowing which record to open.

## Requirement

Anyone who is signed in can page through authors and narrow the list by name.

## Information

Each entry in the list shows the same details as opening that author: name, biography, year of birth, and nationality.

## Who may do this

- Member, librarian, and administrator: yes
- Anyone who is not signed in: no

## Rules

- Results come back one page at a time. The reader can ask for a later page.
- A page holds at most 20 authors.
- A name filter matches authors whose name contains the given text, ignoring letter case.
- With no name filter, the page lists authors in name order.
- When nothing matches, the result is an empty page.
- Searching does not change any author record.

## Done when

- A signed-in person can move through pages and see authors in name order.
- A name filter returns only authors whose name contains that text, regardless of letter case.
- An empty match is an empty page.
- Anyone who is not signed in is refused.
