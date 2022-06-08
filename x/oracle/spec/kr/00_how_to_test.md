## How To Test
본 항목에서는 어떻게 모듈을 테스트 해볼 수 있는지 간단하게 명령어를 기술한다.

아래 명령어를 이용해서 오라클 상태를 업데이트 할 수 있다.
```shell
$ ./script/run_single_node.sh
$ export KEY=$(./build/novad keys show validator -a --keyring-backend=test --home=~/.novad/validator)
$ ./build/novad tx oracle update_state $KEY atom 5000 6 10 \
      --home=~/.novad/validator \
      --keyring-backend=test \
      --chain-id=testing
```