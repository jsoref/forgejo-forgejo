# Development workflow

Forgejo is a soft fork, i.e. a set of commits applied to the Gitea development branch and the stable branches. On a regular basis those commits are rebased and modified if necessary to keep working. All Forgejo commits are merged into a branch from which binary releases and packages are created and distributed. The development workflow is a set of conventions Forgejo developers are expected to follow to work together.

Discussions on how the workflow should evolve happen [in the isssue tracker](https://codeberg.org/forgejo/forgejo/issues?type=all&state=open&labels=&milestone=0&assignee=0&q=%5BWORKFLOW%5D).

## Naming conventions

### Development

* Gitea: main
* Forgejo: forgejo
* Feature branches: forgejo-feature-name

### Stable

* Gitea: release/vX.Y
* Forgejo: vX.Y/forgejo
* Feature branches: vX.Y/forgejo-feature-name

### Soft fork history

Before rebasing on top of Gitea, all branches are copied to `soft-fork/YYYY-MM-DD/<branch>` for safekeeping. Older `soft-fork/*/<branch>` branches are converted into references under the same name. Similar to how pull requests store their head, they do not clutter the list of branches but can be retrieved if needed with `git fetch +refs/soft-fork/*:refs/soft-fork/*`. Tooling to automate this archival process [is available](https://codeberg.org/forgejo-contrib/soft-fork-tools/src/branch/master/README.md#archive-branches).

### Tags

Because the branches are rebased on top of Gitea, only the latest tag will be found in a given branch. For instance `v1.18.0-1` won't be found in the `v1.18/forgejo` branch after it is rebased.

## Rebasing

### *Feature branch*

The *Gitea* branches are mirrored with the Gitea development and stable branches.

On a regular basis, each *Feature branch* is rebased against the base *Gitea* branch.

### forgejo branch

The latest *Gitea* branch resets the *forgejo* branch and all *Feature branches* are merged into it.

If tests pass after pushing *forgejo* to the https://codeberg.org/forgejo-integration/forgejo repository, it can be pushed to the https://codeberg.org/forgejo/forgejo repository.

If tests do not pass, an issue is filed to the *Feature branch* that fails the test. Once the issue is resolved, another round of rebasing starts.

### Cherry picking and rebasing

Because Forgejo is a soft fork of Gitea, the commits in feature branches need to be cherry-picked on top of their base branch. They cannot be rebased using `git rebase`, because their base branch has been rebased.

Here is how the commits in the `forgejo-f3` branch can be cherry-picked on top of the latest `forgejo-development` branch:

```
$ git fetch --all
$ git remote get-url forgejo
git@codeberg.org:forgejo/forgejo.git
$ git checkout -b forgejo/forgejo-f3
$ git reset --hard forgejo/forgejo-development
$ git cherry-pick $(git rev-list --reverse forgejo/soft-fork/2022-12-10/forgejo-development..forgejo/soft-fork/2022-12-10/forgejo-f3)
$ git push --force forgejo-f3 forgejo/forgejo-f3
```

## Feature branches

All *Feature branches* are based on the {vX.Y/,}forgejo-development branch which provides development tools and documentation.

The `forgejo-development` branch is based on the {vX.Y/,}forgejo-ci branch which provides the Woodpecker CI configuration.

The purpose of each *Feature branch* is documented below:

### General purpose

* [forgejo-ci](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-ci) based on [main](https://codeberg.org/forgejo/forgejo/src/branch/main)
  Woodpecker CI configuration, including the release process.
  * Backports: [v1.18/forgejo-ci](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-ci)

* [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development) based on [forgejo-ci](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-ci)
  Forgejo development tools and documentation.
  * Backports: [v1.18/forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-development)

### Dependency

* [forgejo-dependency](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-dependency) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  Each commit is prefixed with the name of dependency in uppercase, for instance **[GOTH]** or **[GITEA]**. They are standalone and implement either a bug fix or a feature that is in the process of being contributed to the dependency. It is better to contribute directly to the dependency instead of adding a commit to this branch but it is sometimes not possible, for instance when someone does not have a GitHub account. The author of the commit is responsible for rebasing and resolve conflicts. The ultimate goal of this branch is to be empty and it is expected that a continuous effort is made to reduce its content so that the technical debt it represents does not burden Forgejo long term.
  * Backports: [v1.18/forgejo-dependency](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-dependency)

### [Privacy](https://codeberg.org/forgejo/forgejo/issues?labels=83271)

* [forgejo-privacy](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-privacy) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  Customize Forgejo to have more privacy.
  * Backports: [v1.18/forgejo-privacy](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-privacy)

### Branding
* [forgejo-branding](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-branding) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  Replacing upstream branding with Forgejo branding

### [Internationalization](https://codeberg.org/forgejo/forgejo/issues?labels=82637)
* [forgejo-i18n](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-i18n) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  Internationalization support for Forgejo with a workflow based on Weblate.

### [Accessibility](https://codeberg.org/forgejo/forgejo/issues?labels=81214)
* Backports only: [v1.18/forgejo-a11y](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-a11y) based on [v1.18/forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-development)
  Backport future upstream a11y improvements to the current release of Forgejo

### [Federation](https://codeberg.org/forgejo/forgejo/issues?labels=79349)

* [forgejo-federation](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-federation) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  Federation support for Forgejo

* [forgejo-f3](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-f3) based on [forgejo-development](https://codeberg.org/forgejo/forgejo/src/branch/forgejo-development)
  [F3](https://lab.forgefriends.org/friendlyforgeformat/gof3) support for Forgejo

## Pull requests and feature branches

Most people who are used to contributing will be familiar with the workflow of sending a pull request against the default branch. When that happens the reviewer should change the base branch to the appropriate *Feature branch* instead. If the pull request does not fit in any *Feature branch*, the reviewer needs to make decision to either:

* Decline the pull request because it is best contributed to Gitea
* Create a new *Feature branch*

Returning contributors can figure out which *Feature branch* to base their pull request on using the list of *Feature branches*.

## Granularity

*Feature branches* can contain a number of commits grouped together, for instance for branding the documentation, the landing page and the footer. It makes it convenient for people working on that topic to get the big picture without browsing multiple branches. Creating a new *Feature branch* for each individual commit, while possible, is likely to be difficult to work with.

Observing the granularity of the existing *Feature branches* is the best way to figure out what works and what does not. It requires adjustments from time to time depending on the number of contributors and the complexity of the Forgejo codebase that sits on top of Gitea.
