# ローカルファイルサーバー

|  メソッド |  URL  |
| ---- | ---- |
|  　POST  | /upload   |
|  　POST  |  /upload/file  |


### /upload 
base64にエンコードされた画像をデコードして画像ファイルを生成し、/uploadimagesディレクトリにファイルを格納する。

### /upload/file
受け取ったファイルを/uploadimagesディレクトリにファイルを格納する。
