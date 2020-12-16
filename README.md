# git

创建本地分支：
$ git branch [name] ----注意新分支创建后不会自动切换为当前分支

切换分支：$ git checkout [name]

创建新分支并立即切换到新分支：
$ git checkout -b [name]

git add -A                        把所有文件添加到暂存区

git commit -m "xxxx"              提交文件

git reset --hard HEAD^            回退到上一个版本

git status                        查看仓库状态

git log                           查看历史记录

git pull                            从远程获取最新版本并merge到本地

git branch wjh origin/master  建立远程子分支

git branch -a  查看远程分支

删除远程分支：git push origin :***   / git push --delete origin ***

goland对应版本的汉化包 https://github.com/pingfangx/jetbrains-in-chinese/tree/master/GoLand

提交本地a分支代码到远程b分支：git push origin <本地分支名>:<远程分支名>


git stash pop
这个就是恢复的操作啦。会将堆栈中的代码删除，并且本地代码恢复到之前存储的代码。当然可以恢复指定的存储代码: git stash pop stash@{1}

git stash apply
这个也是恢复操作 和上面的区别是 他不会删除堆栈中的代码 如果需要恢复指定的 在后面加入对于的key 值就可以了 stash@{x}

git stash show
显示做了哪些改动，默认show第一个存储,如果要显示其他存贮，后面加上对应的key值stash@{@num}

git stash drop stash@{0}
丢弃stash@{$num}存储，从列表中删除这个存储

1. 找到上次git commit的 id

     git log

     找到你想撤销的commit_id

2.  git reset --hard commit_id

      完成撤销,同时将代码恢复到前一commit_id 对应的版本。

3. git reset commit_id 

     完成Commit命令的撤销，但是不对代码修改进行撤销，可以直接通过git commit 重新提交对本地代码的修改

HEAD^的意思是上一个版本，也可以写成HEAD~1
如果你进行了2次commit，想都撤回，可以使用HEAD~2

--mixed 
意思是：不删除工作空间改动代码，撤销commit，并且撤销git add . 操作
这个为默认参数,git reset --mixed HEAD^ 和 git reset HEAD^ 效果是一样的。
 
--soft  
不删除工作空间改动代码，撤销commit，不撤销git add . 
 
--hard
删除工作空间改动代码，撤销commit，撤销git add . 
注意完成这个操作后，就恢复到了上一次的commit状态。
 
 
顺便说一下，如果commit注释写错了，只是想改一下注释，只需要：
git commit --amend
此时会进入默认vim编辑器，修改注释完毕后保存就好了。
