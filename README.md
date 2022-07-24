# VSCode-Go에서 디버그시 키 입력 안먹는 거

## 머?
* https://github.com/Microsoft/vscode-go/issues/219
* https://github.com/golang/vscode-go/issues/124

## 그래서 머???
* 위에 두번째 깃헙 이슈를 보면 해결하려고 작업을 시작한 것 같긴 한데..
* `.vscode` > `launch.json` & `tasks.json`은 내가 `절대네버` 쓰지 않는 기능이라서 앞으로도 쓸 일은 없을 거 같다. 걍 fmt.Scanf 안쓰고 말지 ㅡ,.ㅡㅋ

## 일단 이렇게 진행
<details>
<summary>펼치기</summary>
* 검색해보니 임시방편이 여럿 있는데 그건 `.vscode` 폴더를 보면 되고, 디버그는 아래와 같이 실행
    * F1 > Run Task: `start dlv-dap`
    * dlv창이 뜨면 소스창으로 돌아가서 > F5
    * 입력할 거 입력
</details>

## Prelaunch
* ~~해봤는데 안된다. `vscode-go-nightly` 확장으로도 시도해봤는데 안된다. 걍 task 따로 debug 따로 실행해야될 듯~~
* `.vscode` > `launch.json` & `tasks.json`, `integratedTerminal` 설정으로 된다. 근데 나는 통합터미널을 아래와 같이 설정해놔서 여전히 안된다.
* `Unexpected token '\bin' in expression or statement.` 이런 에러가 뜨고 내가 파워쉘 프롬프트 띄울 때 추가한 `$env:GOBIN=$pwd\\bin` 이거 때문에 안되는 모양인데, 여전히 걍 task 따로 debug 따로 실행해야될 듯
```json
    "terminal.integrated.profiles.windows": {
        "PowerShell -NoProfile": {
            "source": "PowerShell",
            "args": [
                "-NoProfile",
                "-ExecutionPolicy",
                "Bypass",
                "-NoLogo",
                "-NoProfile",
                "-NoExit",
                "-Command",
                "set-alias ll dir; remove-item alias:curl; remove-item alias:wget; remove-item alias:cp; remove-item alias:mv; remove-item alias:rm; remove-item alias:rmdir; function prompt { $host.ui.rawui.WindowTitle = \"PS $pwd\";  $p = $(($pwd -split '\\\\')[-1] -join '\\') + '$ '; if (-not (Test-Path env:GOENVADDED)) { Remove-Variable -name GOPATH; $env:GOBIN = \"$pwd\\bin\"; $env:GOENVADDED = \"exist\"; } return $p }"
            ]
        }
    },
```

끗.
