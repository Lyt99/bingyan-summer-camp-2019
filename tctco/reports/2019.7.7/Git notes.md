# Git notes

### Bug branches

Main idea: store current working **Stage**. Clear and create a new work space.

`git stash`

To fix a bug on `master`:

```bash
git stash
git checkout master
git checkout -b issue-101
```

some debugging... and then

```bash
git add <filename>
git commit -m "message"
git checkout master
git merge --no-ff issue-101
```

Then move back to the former working stage

```bash
git checkout dev
git stash list
```

2 methods to restore:

- `git stash apply (stash@{0})` and then `git stash drop`
- `git stash pop`

### Delete branch

For branches that hasn't been merged `git branch -D <name>`

## Cooperation

To checkout the remote repo: `git remote`, use `git remote -v` for details.

To push local commits to the remote repo: `git push origin master`. If you want to push to other branches: `git push origin dev`

To create sync local store: `git checkout -b dev origin/dev`. If there is any collisions, `git pull` and then `merge` at local. If there is no connection between local and remote repos, do as instructions suggest.

### Rebase

An explicit way to show the path: `git log --graph --pretty=oneline --abbrev-commit`

**`git rebase`** can make git log clearer.



# Tag

```bash
git branch
git checkout master
git tag v1.0
git tag
```

Forget to tag?

```bash
git log --pretty=oneline --abbrev-commit
git tag v0.9 f52c633
git tag
```

To check tag info: `git show <tagname>`

`-a`: tag name, `-m`: description

To delete: 

```bash
git tag -d v0,1
```

To push tag name to the remote repo:

```bash
git push origin v1.0
```

or:

```bash
git push origin --tags
```

To delete from remote repo:

```bash
git tag -d v0.9
git push origin :refs/tags/v0.9
```

