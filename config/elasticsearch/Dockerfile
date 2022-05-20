FROM docker.elastic.co/elasticsearch/elasticsearch:8.0.0

# analyzsis-kuromoji プラグインのインストール
# see: https://www.elastic.co/guide/en/elasticsearch/plugins/current/analysis-kuromoji-analyzer.html
# kuromoji はOSSの日本語形態素解析エンジン
# このプラグインをインストールすることで、プリセット済みにkuromoji analyzerや各種filter, tokenizerが使えるようになる
RUN elasticsearch-plugin install analysis-kuromoji

# ICU Analysisプラグインは、Lucene ICUモジュールをElasticsearchに統合し、
# ICU ライブラリを使用して拡張Unicodeサポートを追加します。
# これには、アジア言語のより優れた分析、Un​​icode正規化、Unicode対応の大文字小文字の区別、照合サポート、および文字変換が含まれます。
RUN elasticsearch-plugin install analysis-icu

RUN echo "search.allow_expensive_queries: false" >> /usr/share/elasticsearch/config/elasticsearch.yml
RUN echo "xpack.security.enabled: false" >> /usr/share/elasticsearch/config/elasticsearch.yml
