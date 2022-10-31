package api

import (
	"net/http"
)

// $ docker run -it --rm --name git -v "~/.ssh/id_rsa":/root/.ssh/id_rsa -v "~/.ssh/id_rsa.pub":/root/.ssh/id_rsa.pub wuliangxue/git:0.1 git clone git@gitlab.com:yisar/vite-project.git

func Prepare(w http.ResponseWriter, r *http.Request) {
	GlobalChan.Messages <- "准备开始"
	Exec("docker", "run", "--rm", "--name", "git", "-it", "-v", "/Users/yisar/repo:/git/repo/", "-v", "/Users/yisar/id_rsa:/root/.ssh/id_rsa", "-v", "/Users/yisar/id_rsa.pub:/root/.ssh/id_rsa.pub", "wuliangxue/git:0.1", "git", "clone", "https://gitlab.com/yisar/vite-project.git")
	GlobalChan.Messages <- "clone完毕"
	Exec("ls", "/Users/yisar/repo/vite-project")
}
