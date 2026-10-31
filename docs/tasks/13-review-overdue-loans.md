# Review overdue loans

**Status:** Not built

## Who this is for

Desk staff need a list of copies that are still out after their due date so they can follow up.

## Requirement

A librarian or an administrator can see every loan that is still out and past its due date.

## Information

Each entry shows the book, the member, when it was borrowed, when it was due, and that it is still out.

## Who may do this

- Librarian and administrator: yes
- Member: no
- Anyone who is not signed in: no

## Rules

- A loan appears when it is still out and the current moment is after its due date.
- A loan that has been returned does not appear, even if it came back late.
- A loan that is still out and not yet due does not appear.
- The list can be paged, at most 20 loans at a time, with the most overdue first.
- When nothing is overdue, the result is an empty page.
- Reading the list does not change any loan or any copy count.

## Done when

- A librarian can see a loan that is still out past its due date, including the book, the member, and the due date.
- An administrator can see the same list.
- A returned loan, and a loan that is not yet due, are absent from the list.
- A member, and anyone who is not signed in, is refused.
