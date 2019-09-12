# Learn words cl

learn-words-cl is a command line tool to learn words.

## 追加したい機能
- アルクから取得した単語の意味をパース、見やすく表示
- 一問ごとに`clear`コマンドをかませて、見やすくしたい
- 現状 List file での管理であるため、SQLiteに移行を検討したい

## Version

### Version 2.0.0
課題
- Mac `say` コマンドでは発音が違う単語がある
- 辞書アプリでは意味を網羅できない単語が多い

機能
- [x] アルクを採用
- [ ] 単語の音声はアルクから入手
- [x] 単語の意味のアルクから入手
- [x] json形式にする意味がないので、ただのリスト形式にする

### Version 1.0.0
機能
- Mac 専用
- Mac `say` コマンドで音声を鳴らす
- 辞書アプリから単語の意味を取得、表示
- json形式で単語を指定
- Gopher道場卒業式で発表

