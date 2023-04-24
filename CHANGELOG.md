# Changelog

## v0.7.7 (2023-04-24)

- change `models.ReplyMarkup` from `any` to interface

## v0.7.6 (2023-04-24)

- support bot api 6.7

## v0.7.5 (2023-04-13)

- add option `WithDebugHandler`

## v0.7.4 (2023-04-10)

- [BUGFIX] change field `SendPollParams.IsAnonymous` to *bool (#26)
- add helpers `bot.True()` and `bot.False()` for define *bool values

## v0.7.3 (2023-04-05)

- make `bot.ProcessUpdate` public 

## v0.7.2 (2023-04-03)

- fix copyMessage method (#24)

## v0.7.1 (2023-03-20)

- fix field name in MenuButtonWebApp #21

## v0.7.0 (2023-03-14)

- support bot api 6.6

## v0.6.0 (2023-03-01)

- [BREAKING] change UserID field type in methods params and models from `int` to `int64` #19

## v0.5.1 (2023-02-10)

- [BUGFIX] send http.NoBody if all params fields are empty. For example, method `getMyCommands` has not required fields

## v0.5.0 (2023-02-06)

- support bot api 6.5 

## v0.4.3 (2023-02-05)

- [BUGFIX] SendChatAction now calls `sendChatAction` instead `sendDice`

## v0.4.2 (2023-01-30)

- [BUGFIX] change field Thumb to InputFile type

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
