# Search the catalog

**Status:** Not built

## Who this is for

A signed-in person needs to find books without already knowing which record to open.

## Requirement

Anyone who is signed in can page through books and narrow the list by title, author, or ISBN.

## Information

Each entry shows the same details as opening that book: title, author, ISBN, description, publication year, total copies, and available copies.

## Who may do this

- Member, librarian, and administrator: yes
- Anyone who is not signed in: no

## Rules

- Results come back one page at a time. The reader can ask for a later page.
- A page holds at most 20 books.
- A title filter matches books whose title contains the given text, ignoring letter case.
- An author filter matches books whose author name contains the given text, ignoring letter case.
- An ISBN filter matches books whose ISBN contains the given text. Hyphens and letter case are ignored.
- Filters that are used together all have to match.
- With no filters, the page lists books in title order.
- When nothing matches, the result is an empty page.
- Searching does not change any book.

## Done when

- A signed-in person can move through pages and see books in title order.
- Title, author, and ISBN filters return only the books that match, including when more than one filter is used together.
- An empty match is an empty page.
- Anyone who is not signed in is refused.
