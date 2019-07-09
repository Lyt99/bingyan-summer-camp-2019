

> git的一些操作：
>
> > **creat a new respository on the command line:**
> >
> > echo "# first" >> README.md
> >
> > git init
> >
> > git add RAEADME.md
> >
> > git commit -m "first commit"
> >
> > git remote add origin https://github.com/yjwang346/first.git
> >
> > git push -u origin master
>
> ![picture插入](C:\Users\王源杰\Pictures\190706.png)
>
> > **创建版本库：**
> >
> > $mkdir learngit
> >
> > $cd learngit
> >
> > $pwd
>
> > **git init**命令可以把目录变成Git可以管理的仓库
>
> > $git add [...]   还有   $git commit -m "xxxx"这些操作都是基本的
>
> 还有些操作，比如git status 查看状态，git log等等



关于***fork***和***pull request***：

*例子：*有一个仓库，叫Repo A。你如果要往里贡献代码，首先要Fork这个Repo，于是在你的Github账号下有		了一个Repo A2,。然后你在这个A2下工作，Commit，push等。然后你希望原始仓库Repo A合并你的工		作，你可以在Github上发起一个Pull Request，意思是请求Repo A的所有者从你的A2合并分支。如果被		审核通过并正式合并，这样你就为项目A做贡献了 

![usage_ofgit](C:\Users\王源杰\Pictures\git.jpg)