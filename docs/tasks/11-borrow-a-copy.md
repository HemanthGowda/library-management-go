# Borrow a copy

**Status:** Not built

## Who this is for

A member needs to take a copy of a book home and know when it is due back.

## Requirement

A member can borrow one available copy of a book for themselves. The loan is due 14 days after it starts, and the number of available copies drops by one.

## Information

A loan remembers:

- Which book was borrowed
- Which member borrowed it
- When it was borrowed
- When it is due (14 days after it was borrowed)
- That it is still out

## Who may do this

- Member: yes, for their own account
- Librarian and administrator: no
- Anyone who is not signed in: no

## Rules

- The book must already exist.
- At least one copy must be available. When none is free, the loan is refused and the available count stays the same.
- A member who already has that same book out cannot borrow a second copy of it until the first loan is returned.
- One borrow takes one copy. Available copies decreases by one. Total copies stays the same.
- The due date is 14 days after the day the loan starts.
- A refused attempt creates no loan and does not change copy counts.

## Done when

- A member can borrow an available book and see the book, the due date 14 days out, and that the loan is still out.
- The book's available copies is one less than before the loan.
- Borrowing is refused when no copy is free, and the available count is unchanged.
- Borrowing is refused when that member already has the same book out.
- A librarian, an administrator, and anyone who is not signed in are refused.
