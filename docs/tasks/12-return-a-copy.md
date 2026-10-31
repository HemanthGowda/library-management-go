# Return a copy

**Status:** Not built

## Who this is for

A reader, or the desk staff, needs to close a loan when the copy comes back so someone else can borrow it.

## Requirement

The member who borrowed the copy, a librarian, or an administrator can return it. The loan is complete, and one copy of that book becomes available again.

## Information

The completed loan remembers when it was returned, and that it is no longer out.

## Who may do this

- The member who owns that loan: yes
- Librarian and administrator: yes, for any open loan
- A different member: no
- Anyone who is not signed in: no

## Rules

- The loan must already exist and still be out.
- Returning an unknown loan is refused.
- Returning a loan that is already complete is refused, and the available count does not change again.
- A successful return marks the loan complete, records the return time, and increases that book's available copies by one. Available copies never rises above total copies.
- The member who borrowed the book can borrow it again after the return.

## Done when

- The borrower can return their own open loan. The loan shows as complete, and the book's available copies is one higher than before the return.
- A librarian and an administrator can return someone else's open loan with the same result.
- A different member is refused, and the loan stays open.
- Returning a loan that is already complete is refused, and the available count stays the same.
- Anyone who is not signed in is refused.
