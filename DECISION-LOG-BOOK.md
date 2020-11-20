# Decision Log Book

Taken from [Decision Management in Software Engineering](https://medium.com/swlh/decision-management-in-software-engineering-ca60f9d40e02)

---

## 2020-11-20: Why should we build this?

### Decision Makers

* Christopher Lamm
* Thomas Jean
* Bob Yexley

### Context

For our web/ui projects that we want to share, we need to publish them. In order to be able to publish them, we need to be able to handle incrementing the versions and maintaining CHANGELOGs for them. We need to be able to do this in a consistent, repeatable manner. We also need to be able to do this for multiple packages within the context of a monorepo project (i.e. @aoeu/ui-common), in which case we need to be able to do these things in multiple folders at once. Existing tooling doesn't currently match the approach that we wish to take to handling executing these tasks.

### Solution

We've decided to roll our own CLI tool, which will ultimately just execute child commands in a terminal for executing the actions that we want to take.

#### Why This Solution

Because existing tools either don't support the workflow that we want, or cost money
#### Limitation

We will have to build and maintain this ourselves.

### Rejected Solutions

* [Lerna](https://lerna.js.org/)
    * Too opinionated and intrusive. Doesn't support making changes without committing and auto-publishing, etc.
