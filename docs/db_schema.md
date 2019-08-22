# DB スキーマ

## main_categories

| カラム名    | 型        | 説明               |
| ----------- | --------- | ------------------ |
| id          | UUID      | ID                 |
| name        | string    | メイン階層の名前   |
| description | string    | 説明　使わないかも |
| created_at  | time.Time |                    |
| updated_at  | time.Time |                    |
| deleted_at  | time.Time |                    |

## sub_categories

| カラム名    | 型        | 説明                        |
| ----------- | --------- | --------------------------- |
| id          | UUID      | ID                          |
| main_id     | UUID      | 所属しているメイン階層の ID |
| name        | string    | メイン階層の名前            |
| description | string    | 説明　使わないかも          |
| created_at  | time.Time |                             |
| updated_at  | time.Time |                             |
| deleted_at  | time.Time |                             |

## contents

| カラム名    | 型        | 説明                                       |
| ----------- | --------- | ------------------------------------------ |
| id          | UUID      | ID                                         |
| category    | UUID      | 所属している階層の ID main_category でも可 |
| title       | string    | コンテンツのタイトル                       |
| image       | string    | メイン画像の URL                           |
| description | string    | コンテンツの説明　 markdown                |
| date        | time.Time | 作品の作成日時　手動で設定                 |
| created_at  | time.Time |                                            |
| updated_at  | time.Time |                                            |
| deleted_at  | time.Time |                                            |

## images

| カラム名   | 型        | 説明         |
| ---------- | --------- | ------------ |
| id         | UUID      | ID           |
| url        | string    | 画像の保存先 |
| created_at | time.Time |              |
| updated_at | time.Time |              |
| deleted_at | time.Time |              |

## tags

| カラム名    | 型        | 説明                  |
| ----------- | --------- | --------------------- |
| id          | UUID      | ID                    |
| name        | string    | タグの表示名          |
| description | string    | タグの説明　 markdown |
| created_at  | time.Time |                       |
| updated_at  | time.Time |                       |
| deleted_at  | time.Time |                       |

## tagged_contents

| カラム名   | 型        | 説明              |
| ---------- | --------- | ----------------- |
| id         | UUID      | ID                |
| tag_id     | UUID      | タグの id         |
| content_id | UUID      | コンテンツの UUID |
| created_at | time.Time |                   |
| updated_at | time.Time |                   |
| deleted_at | time.Time |                   |
