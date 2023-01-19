# Changelog

## v0.4.1 (2023-01-19)

- add func `EscapeMarkdownUnescaped` for escape only unescaped symbols

## v0.4.0 (2023-01-13)

- [BREAKING] change some types from `int` to `int64` #11

## v0.3.4 (2023-01-10)

- support bot api 6.4

## v0.3.3 (2022-12-28)

- add `RegisterHandlerMatchFunc` 

## v0.3.2 (2022-11-28)

- fix field type SendDocumentParams.Document from string to InputFile, add example

## v0.3.1 (2022-11-05)

- support bot api 6.3

## v0.3.0 (2022-10-14)

- **[API CHANGE]** method New now call `getMe` after init and returns error as second result on fail

## v0.2.2 (2022-08-23)

- support bot api v6.2

## v0.2.1 (2022-07-08)

- support bot api v6.1

## v0.2.0 (2022-05-06)

- change API. Use bot funcs instead of`methods` package.

## v0.1.1 (2022-05-02)

- add webhook support
- API refactoring
  - use bot.GetUpdates() instead of bot.Start()

## v0.1.0 (2022-05-02)

- initial release
