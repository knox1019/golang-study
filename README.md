# golang-study
为了成为初级golang工程师的自我学习star from 2025/10/27
时光机——回退
a. 故意把 main.go 改烂，保存。
b. git log --oneline 找到上一条好版本的短 ID（如 a3f1e2d）。
c. git reset --hard a3f1e2d → 文件瞬间恢复。
d. 体验“后悔药”。
已完成√

实验田——新建分支
a. git switch -c feature-test （创建并切换）
b. 随便改代码、commit。
c. git switch main → 发现 main 纹丝不动。
d. 明白“分支就是平行宇宙”。
已完成√

合并——把实验田合并回主线
a. 先回到 main：git switch main
b. git merge feature-test
c. 如无冲突，合并成功；有冲突就按 VS Code 弹出的“Accept Incoming”等按钮解决。
d. 合并后 git branch -d feature-test 删除已完工分支。