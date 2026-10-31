# Change a person's role

**Status:** Not built

## Who this is for

An administrator needs to decide whether a person is a member, a librarian, or an administrator.

## Requirement

An administrator can set another person's role to member, librarian, or administrator. The new role is what that person is allowed to do the next time their permissions are checked.

## Information

The account remembers a single role: member, librarian, or administrator.

## Who may do this

- Administrator: yes
- Member and librarian: no
- Anyone who is not signed in: no

## Rules

- The account must already exist.
- The only roles that can be assigned are member, librarian, and administrator.
- Any other role is refused, and the account keeps its current role.
- An administrator cannot take the administrator role away from the last remaining administrator. That attempt is refused.
- An administrator may change their own role only when at least one other administrator would remain.
- The password and the email stay as they were.

## Done when

- An administrator can change a member into a librarian, and that person can then do librarian work.
- An administrator can change a librarian or an administrator into a member, and that person can no longer do staff work.
- Assigning a role outside member, librarian, and administrator is refused.
- Removing the administrator role from the last administrator is refused, and that person remains an administrator.
- A member, a librarian, and anyone who is not signed in are refused.
