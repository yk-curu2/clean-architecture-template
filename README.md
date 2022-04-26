# clean-architecture-template

## クリーンアーキテクチャの特徴

>  1. Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.  
>  1. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
>  1. Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
>  1. Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
>  1. Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

1. フレームワークに依存しない
1. UI・DB・Web サーバーなどに依存せずにビジネスロジックのテスト可能
1. UI に依存しない. たとえ Web UI を CUI に変更してもビジネスロジックはどうするように
1. DB に依存しない. MySQL / Mongo などに DB を容易に変更できるように
1. 外側のレイヤに依存しない

## 命名規則

| | | | 
| --- | --- | --- |
| パッケージ名 | 単数形 | Go では一般的に複数形の単語をパッケージ名に使わない

## 構成

| layer | folder name | description |
| --- | --- | --- |
| Entities | entity |  ビジネスモデル（オブジェクト）<br>特定の「図」ではなく図が伝えようとしている「考え方」 | 
| Use Cases | usecase | <ul><li>アプリケーション固有のビジネスロジック</li><li>Entity -> Entity へのデータフローの制御</li><li>DB や UI、その他フレームワークが変わっても影響を受けない</li></ul> |
| Interface Adapters | controller, service | <ul><li>interface の method を定義する</li><li>DB の操作 HTTP 入出力などを定義</li></ul> |
| Web | app | web アプリケーションを起動する | 
