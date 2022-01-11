# VSCode-Go에서 디버그시 키 입력 안먹는 거

## 머?
* https://github.com/Microsoft/vscode-go/issues/219
* https://github.com/golang/vscode-go/issues/124

## 그래서 머???
* 위에 두번째 깃헙 이슈를 보면 해결하려고 작업을 시작한 것 같긴 한데..
* `.vscode` > `launch.json` & `tasks.json`은 내가 `절대네버` 쓰지 않는 기능이라서 앞으로도 쓸 일은 없을 거 같다. 걍 fmt.Scanf 안쓰고 말지 ㅡ,.ㅡㅋ

## 일단 이렇게 진행
* 검색해보니 임시방편이 여럿 있는데 그건 `.vscode` 폴더를 보면 되고, 디버그는 아래와 같이 실행
    * F1 > Run Task: `start dlv-dap`
    * dlv창이 뜨면 소스창으로 돌아가서 > F5
    * 입력할 거 입력

끗.
