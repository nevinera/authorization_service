# Authorization Service

This is a microservice that tracks users, groups, and memberships, and exposes endpoints intended
to help answer questions like "is User(X) a member of any of Groups(Y, Z, Q)?" and "What Users are
members of all of Groups(Y, Z)?". It must be able to answer these questions *quickly*, as various
resource services will be asking them inside every client request.

It represents Users and Groups as uuids (which are supplied to it), and exposes api operations to
create and delete them, as well as to add and remove memberships. Memberships are not complex, and
do not represent roles at all - rather, we can represent having the 'admin' role on the
'Account(55)' Group by creating an 'Account(55).admin' group.

It also includes scripts for conveniently interacting with the API from the command-line - a full
UI is unlikely to be useful, as this service doesn't know anything about the users or groups
involved aside from their UUIDS.
