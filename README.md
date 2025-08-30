# comi-track

## やること

- ディレクトリ構成を TS を参考にする
  - 集約でディレクトリ分割
- ロガー
- エラー設計
  - cause 不要
- ミドルウェアでエラー処理
- インポート制限
- リンター
- トランザクション
- リクエスト ID


## 設計

- artists
  - id
  - x_id
  - name
  - favorite_score
- booths
  - id
  - edition_number
  - day
  - hall_name
  - block_name
  - block_code
  - space_number
  - space_half
  - artist_id