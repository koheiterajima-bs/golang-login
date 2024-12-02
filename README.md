## 学んだこと
- volumesについて
  - [ホスト側の相対パス]:コンテナの絶対パスで記述で記述すると、service内に記述
  - contextにホスト側のパスを記述し、[volume名]:[コンテナの絶対パス]で記述すると、service外に記述
- docker-compose run --rm app /bin/ashコマンド
  - appサービスの実行
  - /bin/ashはAlpine Linuxのシェルで、コンテナ内のコマンドを直接実行
- シバンについて
  - #!/bin/sh
  - このスクリプトをどのプログラムで実行するかの指定
- 実行権限について
  - % chmod a+x ./init/*.sh
    - Change Mode(権限変更)の略
    - a + x : a(all)にx(実行権限)を与える記述
    - 全ユーザーに実行権限が付与される
- airについて
  - golangをDockerコンテナで扱う時、ローカルでのコード変更をコンテナ側に反映するためにリロードをする必要がある、airではこのリロード作業を自動化してくれる
  - .air.tomlファイルは一旦公式のexampleをそのまま使用
- GORMについて
  - GORMはGoのORM(Object-Relational Mapping)ライブラリで、データベース操作を簡素化し、効率化するツール
  - データベース接続、クエリ作成、リレーション管理、トランザクション管理、マイグレーションなど煩雑な作業を省略し、Goでの開発が効率的に行える


## わからないこと
- .air.tomlファイルが見つからないとエラー -> ルートディレクトリに配置し、解決
```
2024-12-02 16:17:36 2024/12/02 07:17:36 open .air.toml: no such file or directory
```

