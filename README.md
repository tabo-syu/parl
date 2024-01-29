# Parl

Parl は  Palworld のゲームサーバーである PalServer-Linux を操作するコマンドを Discord サーバーに追加します。
`/parl` は、ゲームサーバーの起動、停止、状態確認を行えます。

- `/parl status`: Palworld ゲームサーバーの状態を確認します。
- `/parl start`: Palworld ゲームサーバーを起動します。
- `/parl stop`: Palworld ゲームサーバーを停止します。

不意に停止したゲームサーバーの起動を Discord サーバーに参加しているメンバーに任せることができるようになります。

## 導入手順

### 環境

- 必須
  - [Go](https://go.dev/)
- 推奨
  - [Git](https://git-scm.com/)
  - [Task](https://taskfile.dev/ja-JP/)
  - screen コマンド

### Bot の作成

1. [Discord Developer Portal](https://discord.com/developers/applications) を開きます。
1. 「New Application」 から Bot を作成します。
1. Bot タブを開き、「Reset Token」から Token を発行し、控えます。
1. 「PUBLIC BOT」のチェックを外します。
1. OAuth2 > URL Generator から 招待 URL を作成します。
    - SCOPES: bot, application.commands
    - BOT PERMISSIONS: Send Messages, Embed Links
1. 生成した URL へ遷移し、サーバーへ Bot を追加します。

### Bot の起動

1. `git clone https://github.com/tabo-syu/parl.git && cd ./parl` を実行します。
1. `cp .env.example .env` を実行します。
1. 好きなエディタで `.env` ファイルを開き、各項目を埋めます。
    - `DISCORD_ICON_URL`: 送信されるメッセージに付与される画像の URL
    - `DISCORD_TOKEN`: 「Bot の作成」で控えた Token
    - `RCON_HOST`: ゲームサーバーの IP アドレス
    - `RCON_PORT`: `PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini` で設定したゲームサーバーのポート番号
    - `RCON_PASSWORD`: `PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini` で設定したパスワード 
    - `SERVER_PATH`: `PalServer/PalServer.sh` のパス
1. `screen` を実行後、`task run` で Bot が起動します。
1. `Ctrl` + `a` > `d` でセッションから抜けます。

# EN

Translated by DeepL.

Parl adds commands to the Discord server to control PalServer-Linux, Palworld's game server. /parl can start, stop, and check the status of the game server.

- `/parl status`: Checks the status of the Palworld game server.
- `/parl start`: Starts the Palworld game server.
- `/parl stop`: Stop the Palworld game server.

This allows a member of the Discord server to start up a game server that has been stopped unexpectedly.

## Installation

### Requirements

- required
  - [Go](https://go.dev/)
- recommend
  - [Git](https://git-scm.com/)
  - [Task](https://taskfile.dev/ja-JP/)
  - `screen` command

### Create a Bot

1. Open the [Discord Developer Portal](https://discord.com/developers/applications).
2. Create a Bot from "New Application".
3. Open the Bot tab, issue a Token from "Reset Token", and copy the Token.
4. Uncheck "PUBLIC BOT".
5. Create an invitation URL from OAuth2 > URL Generator.
    - SCOPES: bot, application.commands
    - BOT PERMISSIONS: Send Messages, Embed Links
6. Go to the generated URL and add the Bot to the server.

### Start Bot

1. Execute `git clone https://github.com/tabo-syu/parl.git && cd . /parl.`
1. Execute `cp .env.example .env`.
1. Open the .env file in your favorite editor and fill in each item.
    - `DISCORD_ICON_URL`: URL of the image to be attached to the message to be sent.
    - `DISCORD_TOKEN`: The Token that you have taken down in the "Creating a Bot" section.
    - `RCON_HOST`: IP address of the game server
    - `RCON_PORT`: Port number of the game server set in `PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini`
    - `RCON_PASSWORD`: Password set in `PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini`
    - `SERVER_PATH`: Path to `PalServer/PalServer.sh`
1. After screen is executed, `task run` will start the Bot.
1. `Ctrl + a` > `d` to exit from the session.
