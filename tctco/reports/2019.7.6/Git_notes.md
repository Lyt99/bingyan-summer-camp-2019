# Git Notes

~~Review~~ 

Learn it again...

### Other Version Control Systems (SVN & CVS)

**Centralized version control System** and **Distributed Version Control System**

The problem with the Centralized version control system is that it requires internet connection.

### Git Initialization

For Windows: download the required exe.

```bash
git config --global user.name "name"
git config --global user.eamil "email"
```

For Linux:

```bash
sudo apt install git
```



## Repository

A directory that can be traced by Git.

To create a repository: 

```bash
git init
```

This will create a dir named `.git`. To see `.git`:

```bash
ls -ah
```

To track any files:

```bash
git add filename.txt
```

```bash
git commit -m "message"
```

Since commit command will submit a lot of files, the designer hope users can add several times.



### Time machine

To see status:  `git status`

To see difference: `git diff`



#### Version Regression

To see log: `git log` This command displays the commit logs from the most recent to the oldest. `git log --pretty=online` will make the info prettier and more simple. The sequence is called **commit id**.

`HEAD` is the current version while `HEAD^` is the last but one version. Similarly, we have `HEAD^^` etc. When the number becomes too large, use `HEAD~100`.

```bash
git reset --hard HEAD^
```

Git log will discard any logging info after the pointed version. To go back again

1. Find the sequence 

   ```bash
   git reset --hard sequence_num(only the first few digits)
   ```

2. Use `git reflog`, this reflog contains everything you write. The reflog only displays **reset** info.



#### Working Directory and Stage (Index)

![An image about Git](./pics/Stage.jpg)



#### Manage changes

To see the difference between working directory and the files stored in the 

repository, use: 

```bash
git diff HEAD -- a.txt b.txt
```



#### Reverse

```bash
git checkout -- try.txt
```

This command will reverse all the changes in the working directory:

1. `try.txt` is not in the **Stage**. Go back to the status of the **repository**.
2. `try.txt` is in the **Stage** and modified again. Go back to the **Stage**.

In summary, `git checkout -- path` will reverse to the last `git commit` or `git add`.

*Remark: `--` is very important! `-` means "switch to another branch"!*



```bash
git reset HEAD <file>
```

This will **unstage** the file (put it back to the working directory).

*Review: `git reset` can be used both to move between versions and move files in **Stage** back to the working directory.*

Summary: 

- `git checkout -- path` reverse working directory.
- `git reset HEAD <file>` move the file out of the **Stage**.

PS: Already commit? Go back to the **Version Regression**.



#### Deleting

`git status` will tell which file has been deleted.

1. Delete file from repository? 

   ```bash
   git rm LICENSE
   ```

   ```bash
   git commit -m "deleting LICENSE"
   ```

2. Error? Reverse the file from the repo.

   ```bash
   git checkout -- LICENSE
   ```

   **This action shows that `checkout` command simply replace the file in the working directory with the one in the repo.** Thus `checkout` can be used to reverse changes or reverse files.



## Remote Repository

The connection between local git repo and the Github repo is encrypted by SSH.

1. Create SSH Key (.ssh in User dir?)

   ```bash
   ssh-keygen -t rsa -C "email
   ```

2. Add SSH Key (.pub) to Github.

   

#### Adding remote repo

1. Create a new repo on Github.

2. Run this at local repo to link 2 repos. **Origin** is the default name for the remote repo.

   ```bash
   git remote add origin git@github@.com:<...>
   ```

3. Run `push` command:

   ```bash
   git push -u origin master
   ```

   Since the remote repo is empty, adding `-u` can push *local* `master` to *remote* `master` and also link them. Thus, fewer commands are needed for future `push` or `pull`.
   
4. For future updates:

   ```bash
   git push origin master
   ```

   

#### Clone

"Download" repos from github.

```bash
git clone git@github.com:<...>
```

Notes: ssh is faster than https protocol. Https also requires password each time.



## Branches

![branches_!](./pics/branches_1.png)

![branches_!](./pics/branches_2.png)

![branches_!](./pics/branches_3.png)

```bash
git checkout -b dev
```

or

```bash
git branch dev
git checkout dev
```

*Review: `checkout` can be used to reverse changes that isn't added to **Stage***

| look up      | create              | switch                | create and switch        | merge to current branch | delete                 |
| ------------ | ------------------- | --------------------- | ------------------------ | ----------------------- | ---------------------- |
| `git branch` | `git branch <name>` | `git checkout <name>` | `git checkout -b <name>` | `git merge <name>`      | `git branch -d <name>` |

#### Collisions

`git status` helps locate the file.

After fixing the collisions, use `git add` and `git commit`.

`git log` shows the merging process.

#### Managing branches

In **Fast forward** mode, after deleting the branch, user will lose info in that branch. If **Fast forward** is forced stop, Git will generate a new commit when merging. Thus the branch info will be saved in the branch history.

To force stop **Fast forward**, use 

```bash
git merge --no-ff -m "message" dev
```

This is more like the case of **collision**.

![noff](./pics/noff.png)

The common strategy for developing new products will be:

![developing_strategy](./pics/developing_strategy.png)